package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/svc"
	"github.com/solodba/devcloud/mpaas/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svcService svc.SvcService
	ctx        = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svcService = apps.GetInternalApp(svc.AppName).(svc.SvcService)
}
