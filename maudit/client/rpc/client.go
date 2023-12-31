package rpc

import (
	"github.com/solodba/mcube/logger"
	"google.golang.org/grpc"
)

// GRPC客户端结构体
type Client struct {
	c *grpc.ClientConn
}

// GRPC客户端初始化函数
func NewClient(conf *Config) *Client {
	clientConn, err := grpc.Dial(conf.address, grpc.WithInsecure())
	if err != nil {
		logger.L().Panic().Msgf("grpc client dial error, err: %s", err.Error())
	}
	return &Client{
		c: clientConn,
	}
}

// GRPC初始化客户端
