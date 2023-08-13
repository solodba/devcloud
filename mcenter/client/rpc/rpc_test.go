package rpc_test

import (
	"context"

	"github.com/solodba/devcloud/mcenter/client/rpc"
)

// 全局参数
var (
	c   *rpc.Client
	ctx = context.Background()
)

// 初始化函数
func init() {
	conf := rpc.NewDefaultMcenterGrpcClientConfig()
	conf.ClientID = "cjauuq4fd1fkek1bmfq0"
	conf.ClientSecret = "cjauuq4fd1fkek1bmfqg"
	c = rpc.NewClient(conf)
}
