package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/configmap"
	"github.com/solodba/devcloud/mpaas/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svc configmap.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(configmap.AppName).(configmap.Service)
}