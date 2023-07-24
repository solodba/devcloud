package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/daemonset"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc daemonset.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return daemonset.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(daemonset.AppName).(daemonset.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"DaemonSet管理"}
	// webservice定义路由信息
	// 创建DaemonSet
	ws.Route(ws.POST("/").To(h.CreateDaemonSet).
		Doc("创建DaemonSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(daemonset.CreateDaemonSetRequest{}).
		Writes(daemonset.DaemonSet{}).
		Returns(200, "OK", daemonset.DaemonSet{}))

	// 删除DaemonSet
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteDaemonSet).
		Doc("删除DaemonSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(daemonset.DeleteDaemonSetRequest{}).
		Writes(daemonset.CreateDaemonSetRequest{}).
		Returns(200, "OK", daemonset.CreateDaemonSetRequest{}))

	// 更新DaemonSet
	ws.Route(ws.PUT("/").To(h.UpdateDaemonSet).
		Doc("更新DaemonSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(daemonset.UpdateDaemonSetRequest{}).
		Writes(daemonset.DaemonSet{}).
		Returns(200, "OK", daemonset.DaemonSet{}))

	// 查询DaemonSet
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeDaemonSet).
		Doc("查询DaemonSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(daemonset.QueryDaemonSetRequest{}).
		Writes(daemonset.DaemonSetList{}).
		Returns(200, "OK", daemonset.DaemonSetList{}).
		Reads(daemonset.DescribeDaemonSetRequest{}).
		Writes(daemonset.CreateDaemonSetRequest{}).
		Returns(200, "OK", daemonset.CreateDaemonSetRequest{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
