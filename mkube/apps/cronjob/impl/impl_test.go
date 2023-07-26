package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mkube/apps/cronjob"
	"github.com/solodba/devcloud/mkube/test/tools"
	"github.com/solodba/mcube/apps"
)

// 全局变量
var (
	svc cronjob.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(cronjob.AppName).(cronjob.Service)
}
