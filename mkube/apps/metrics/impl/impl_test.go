package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/devcloud/mkube/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc metrics.Service
	ctx = context.Background()
)

// 初始化函数
func init() {
	// 加载项目环境变量和IOC配置初始化
	tools.DevelopmentSet()
	// 从IOC获取user测试实例
	svc = apps.GetInternalApp(metrics.AppName).(metrics.Service)
}
