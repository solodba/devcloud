package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mkube/apps/prometheus"
	"github.com/solodba/devcloud/mkube/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc prometheus.Service
	ctx = context.Background()
)

// 初始化函数
func init() {
	// 加载项目环境变量和IOC配置初始化
	tools.DevelopmentSet()
	// 从IOC获取user测试实例
	svc = apps.GetInternalApp(prometheus.AppName).(prometheus.Service)
}
