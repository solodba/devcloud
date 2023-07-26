package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/ingress"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc ingress.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return ingress.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(ingress.AppName).(ingress.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Ingress管理"}
	// webservice定义路由信息
	// 创建Ingress
	ws.Route(ws.POST("/").To(h.CreateIngress).
		Doc("创建Ingress").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingress.CreateIngressRequest{}).
		Writes(ingress.Ingress{}).
		Returns(200, "OK", ingress.Ingress{}))

	// 删除Ingress
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteIngress).
		Doc("删除Ingress").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingress.DeleteIngressRequest{}).
		Writes(ingress.CreateIngressRequest{}).
		Returns(200, "OK", ingress.CreateIngressRequest{}))

	// 更新Ingress
	ws.Route(ws.PUT("/").To(h.UpdateIngress).
		Doc("更新Ingress").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingress.UpdateIngressRequest{}).
		Writes(ingress.Ingress{}).
		Returns(200, "OK", ingress.Ingress{}))

	// 查询ConfigMap
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeIngress).
		Doc("查询Ingress").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingress.QueryIngressRequest{}).
		Writes(ingress.IngressSet{}).
		Returns(200, "OK", ingress.IngressSet{}).
		Reads(ingress.DescribeIngressRequest{}).
		Writes(ingress.CreateIngressRequest{}).
		Returns(200, "OK", ingress.CreateIngressRequest{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
