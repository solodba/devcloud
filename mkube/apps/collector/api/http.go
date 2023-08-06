package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/collector"
	"github.com/solodba/mcube/apps"
)

// hander结构体指针变量
var (
	h = &handler{}
)

// handler结构体
type handler struct {
	svc collector.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return collector.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取collector实例
	h.svc = apps.GetInternalApp(collector.AppName).(collector.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"prometheus采集器"}
	// webservice定义路由信息
	ws.Route(ws.GET("/prometheus").To(h.GetMetrics).
		Doc("prometheus采集器").
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(h)
}
