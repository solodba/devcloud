package rpc_test

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/client/rpc"
)

// 全局参数
var (
	c   *rpc.Client
	ctx = context.Background()
)

// 初始化函数
func init() {
	c = rpc.NewClient("127.0.0.1:8889")
}
