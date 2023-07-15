package impl_test

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mpaas/apps/pod"
	"github.com/solodba/devcloud/tree/main/mpaas/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svc pod.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(pod.AppName).(pod.Service)
}
