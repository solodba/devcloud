package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/ingroute"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc ingroute.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return ingroute.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(ingroute.AppName).(ingroute.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"IngressRoute管理"}
	// webservice定义路由信息
	// 创建IngressRoute
	ws.Route(ws.POST("/").To(h.CreateIngressRoute).
		Doc("创建IngressRoute").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingroute.CreateIngressRouteRequest{}).
		Writes(ingroute.IngressRoute{}).
		Returns(200, "OK", ingroute.IngressRoute{}))

	// 删除IngressRoute
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteIngressRoute).
		Doc("删除IngressRoute").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingroute.DeleteIngressRouteRequest{}).
		Writes(ingroute.CreateIngressRouteRequest{}).
		Returns(200, "OK", ingroute.CreateIngressRouteRequest{}))

	// 更新IngressRoute
	ws.Route(ws.PUT("/").To(h.UpdateIngressRoute).
		Doc("更新IngressRoute").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingroute.UpdateIngressRouteRequest{}).
		Writes(ingroute.IngressRoute{}).
		Returns(200, "OK", ingroute.IngressRoute{}))

	// 查询IngressRoute
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeIngressRoute).
		Doc("查询IngressRoute").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingroute.QueryIngRouteMwRequest{}).
		Writes(ingroute.IngressRouteSet{}).
		Returns(200, "OK", ingroute.IngressRouteSet{}).
		Reads(ingroute.DescribeIngressRouteRequest{}).
		Writes(ingroute.CreateIngressRouteRequest{}).
		Returns(200, "OK", ingroute.CreateIngressRouteRequest{}))

	// 查询IngressRoute Middlware List
	ws.Route(ws.GET("/ingroute/{namespace}/middleware").To(h.QueryIngRouteMiddlewareList).
		Doc("查询IngressRoute Middleware List").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(ingroute.QueryIngRouteMwRequest{}).
		Writes(ingroute.MiddlewareList{}).
		Returns(200, "OK", ingroute.MiddlewareList{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
