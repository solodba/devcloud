package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/pv"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc pv.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return pv.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(pv.AppName).(pv.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"PersistentVolume管理"}
	// webservice定义路由信息
	// 创建PersistentVolume
	ws.Route(ws.POST("/").To(h.CreatePV).
		Doc("创建PersistentVolume").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(pv.CreatePVRequest{}).
		Writes(pv.PV{}).
		Returns(200, "OK", pv.PV{}))

	// 删除PersistentVolume
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeletePV).
		Doc("删除PersistentVolume").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(pv.DeletePVRequest{}).
		Writes(pv.PV{}).
		Returns(200, "OK", pv.PV{}))

	// 查询PersistentVolume
	ws.Route(ws.GET("/{namespace}").To(h.QueryPV).
		Doc("查询PersistentVolume").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(pv.QueryPVRequest{}).
		Writes(pv.PVSet{}).
		Returns(200, "OK", pv.PVSet{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
