package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/dashboard"
	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/mcube/apps"
)

// hander结构体指针变量
var (
	h = &handler{}
)

// handler结构体
type handler struct {
	svc metrics.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return dashboard.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取metrics实例
	h.svc = apps.GetInternalApp(metrics.AppName).(metrics.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"收集Dashboard数据"}
	// webservice定义路由信息
	ws.Route(ws.GET("/usage").To(h.QueryClusterUsage).
		Doc("查询K8S集群使用情况").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(metrics.QueryClusterUsageRequest{}).
		Writes(metrics.MetricSet{}).
		Returns(200, "OK", metrics.MetricSet{}))

	ws.Route(ws.GET("/cluster").To(h.QueryClusterInfo).
		Doc("查询K8S集群信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(metrics.QueryClusterInfoRequest{}).
		Writes(metrics.MetricSet{}).
		Returns(200, "OK", metrics.MetricSet{}))

	ws.Route(ws.GET("/resource").To(h.QueryResource).
		Doc("查询K8S集群资源信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(metrics.QueryResourceRequest{}).
		Writes(metrics.MetricSet{}).
		Returns(200, "OK", metrics.MetricSet{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(h)
}
