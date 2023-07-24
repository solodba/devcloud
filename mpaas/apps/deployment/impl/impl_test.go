package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/deployment"
	"github.com/solodba/devcloud/mpaas/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svc deployment.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(deployment.AppName).(deployment.Service)
}
