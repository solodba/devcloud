package protocol

import (
	"fmt"
	"net"

	"github.com/infraboard/mcube/grpc/middleware/recovery"
	"github.com/solodba/devcloud/mcenter/conf"
	"github.com/solodba/devcloud/mcenter/protocol/auth"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	"google.golang.org/grpc"
)

// GRPC服务结构体
type GrpcService struct {
	svr *grpc.Server
	c   *conf.Config
}

// GRPC服务结构体初始化
func NewGrpcService() *GrpcService {
	// 添加grpc认证中间
	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(rc.UnaryServerInterceptor()),
		grpc.ChainUnaryInterceptor(auth.NewGrpcAuther().AuthFunc))
	return &GrpcService{
		svr: srv,
		c:   conf.C(),
	}
}

// GRPC服务启动方法
func (s *GrpcService) Start() error {
	// 初始化GRPC服务配置
	if err := apps.InitGrpcApps(s.svr); err != nil {
		logger.L().Info().Msgf("initial grpc object config error, err: %s", err.Error())
		return err
	}
	listener, err := net.Listen("tcp", s.c.App.Grpc.Addr())
	if err != nil {
		logger.L().Info().Msgf("tcp listener error, err: %s", err.Error())
		return err
	}
	err = s.svr.Serve(listener)
	if err != nil {
		if err == grpc.ErrServerStopped {
			logger.L().Info().Msgf("grpc service is stopped")
		}
		return fmt.Errorf("start grpc service error, err: %s", err.Error())
	}
	return nil
}

// GRPC服务停止方法
func (s *GrpcService) Stop() error {
	s.svr.GracefulStop()
	return nil
}
