package rpc

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/apps/user"
	"github.com/solodba/mcube/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPC客户端结构体
type Client struct {
	c *grpc.ClientConn
}

// GRPC客户端初始化函数
func NewClient(conf *Config) *Client {
	clientConn, err := grpc.DialContext(
		context.Background(),
		conf.address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 客户端拨号传递认证信息到grpc客户端
		grpc.WithPerRPCCredentials(conf),
	)
	if err != nil {
		logger.L().Panic().Msgf("grpc client dial error, err: %s", err.Error())
	}
	return &Client{
		c: clientConn,
	}
}

// GRPC初始化客户端
func (c *Client) NewUserRPCClient() user.RPCClient {
	return user.NewRPCClient(c.c)
}

func (c *Client) NewTokenRPCClient() token.RPCClient {
	return token.NewRPCClient(c.c)
}
