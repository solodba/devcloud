package start

import (
	"context"

	"github.com/solodba/devcloud/maudit/conf"
	"github.com/solodba/devcloud/maudit/controller"
	"github.com/solodba/devcloud/maudit/protocol"
	"github.com/solodba/mcube/logger"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "maudit start service",
	Long:    "maudit start service",
	Example: "maudit-api start -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewServer()
		if err := srv.Start(); err != nil {
			return err
		}
		return nil
	},
}

// 服务结构体
type Server struct {
	GrpcService *protocol.GrpcService
	HttpService *protocol.HttpService
	// audit log controller
	Controller *controller.AuditLogSaveController
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		GrpcService: protocol.NewGrpcService(),
		HttpService: protocol.NewHttpService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	s.Controller = controller.NewAuditLogSaveController(
		conf.C().Kafka.Endpoints,
		conf.C().Kafka.ConsumerGroup,
		conf.C().Kafka.Topic,
	)
	go s.Controller.Run(context.Background())
	logger.L().Info().Msgf("start controller ok, consumer group: %s", conf.C().Kafka.ConsumerGroup)
	go s.GrpcService.Start()
	return s.HttpService.Start()
}

// Server服务停止方法
func (s *Server) Stop() error {
	err := s.GrpcService.Stop()
	if err != nil {
		return err
	}
	err = s.HttpService.Stop()
	if err != nil {
		return err
	}
	err = s.Controller.Stop()
	if err != nil {
		return err
	}
	return nil
}
