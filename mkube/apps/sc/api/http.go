package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/sc"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc sc.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return sc.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(sc.AppName).(sc.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"StorageClass管理"}
	// webservice定义路由信息
	// 创建StorageClass
	ws.Route(ws.POST("/").To(h.CreateSC).
		Doc("创建StorageClass").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(sc.CreateSCRequest{}).
		Writes(sc.SC{}).
		Returns(200, "OK", sc.SC{}))

	// 删除StorageClass
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteSC).
		Doc("删除StorageClass").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(sc.DeleteSCRequest{}).
		Writes(sc.SC{}).
		Returns(200, "OK", sc.SC{}))

	// 查询StorageClass
	ws.Route(ws.GET("/{namespace}").To(h.QuerySC).
		Doc("查询StorageClass").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(sc.QuerySCRequest{}).
		Writes(sc.SCSet{}).
		Returns(200, "OK", sc.SCSet{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
