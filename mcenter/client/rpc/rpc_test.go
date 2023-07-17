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
	conf := rpc.NewConfig()
	c = rpc.NewClient(conf)
}
