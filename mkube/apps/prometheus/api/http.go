package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/devcloud/mkube/apps/prometheus"
	"github.com/solodba/mcube/apps"
)

// hander结构体指针变量
var (
	h = &handler{}
)

// handler结构体
type handler struct {
	svc prometheus.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return prometheus.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取prometheus实例
	h.svc = apps.GetInternalApp(prometheus.AppName).(prometheus.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"prometheus监控数据"}
	// webservice定义路由信息
	ws.Route(ws.GET("/usageRange").To(h.QueryClusterUsageRange).
		Doc("获取集群时间段变化Prometheus监控数据").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(prometheus.QueryClusterUsageRangeRequest{}).
		Writes(metrics.MetricSet{}).
		Returns(200, "ok", metrics.MetricSet{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(h)
}
