package start

import (
	"github.com/solodba/devcloud/cmdb/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "cmdb start service",
	Long:    "cmdb start service",
	Example: "cmdb-api start -f etc/config.toml",
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
	return nil
}
