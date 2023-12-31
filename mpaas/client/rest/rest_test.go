package rest_test

import (
	"context"

	"github.com/solodba/devcloud/mpaas/client/rest"
)

// 全局参数
var (
	c   *rest.Client
	ctx = context.Background()
)

// 初始化函数
func init() {
	conf := rest.NewConfig()
	c = rest.NewClient(conf)
}
