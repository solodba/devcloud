package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/svc"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc svc.SvcService
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return svc.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(svc.AppName).(svc.SvcService)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Service管理"}
	// webservice定义路由信息
	// 创建Service
	ws.Route(ws.POST("/").To(h.CreateService).
		Doc("创建Service").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(svc.CreateServiceRequest{}).
		Writes(svc.Service{}).
		Returns(200, "OK", svc.Service{}))

	// 删除Service
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteService).
		Doc("删除Service").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(svc.DeleteServiceRequest{}).
		Writes(svc.Service{}).
		Returns(200, "OK", svc.Service{}))

	// 更新Service
	ws.Route(ws.PUT("/").To(h.UpdateService).
		Doc("更新Service").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(svc.UpdateServiceRequest{}).
		Writes(svc.Service{}).
		Returns(200, "OK", svc.Service{}))

	// 查询Service
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeService).
		Doc("查询Service").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(svc.QueryServiceRequest{}).
		Writes(svc.ServiceSet{}).
		Returns(200, "OK", svc.ServiceSet{}).
		Reads(svc.DescribeServiceRequest{}).
		Writes(svc.Service{}).
		Returns(200, "OK", svc.Service{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
