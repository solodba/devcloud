package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/statefulset"
	"github.com/solodba/devcloud/mpaas/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svc statefulset.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(statefulset.AppName).(statefulset.Service)
}
