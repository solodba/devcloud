package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/statefulset"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc statefulset.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return statefulset.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(statefulset.AppName).(statefulset.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"StatefulSet管理"}
	// webservice定义路由信息
	// 创建StatefulSet
	ws.Route(ws.POST("/").To(h.CreateStatefulSet).
		Doc("创建StatefulSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(statefulset.CreateStatefulSetRequest{}).
		Writes(statefulset.StatefulSet{}).
		Returns(200, "OK", statefulset.StatefulSet{}))

	// 删除StatefulSet
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteStatefulSet).
		Doc("删除StatefulSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(statefulset.DeleteStatefulSetRequest{}).
		Writes(statefulset.CreateStatefulSetRequest{}).
		Returns(200, "OK", statefulset.CreateStatefulSetRequest{}))

	// 更新StatefulSet
	ws.Route(ws.PUT("/").To(h.UpdateStatefulSet).
		Doc("更新StatefulSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(statefulset.UpdateStatefulSetRequest{}).
		Writes(statefulset.StatefulSet{}).
		Returns(200, "OK", statefulset.StatefulSet{}))

	// 查询StatefulSet
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeStatefulSet).
		Doc("查询StatefulSet").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(statefulset.QueryStatefulSetRequest{}).
		Writes(statefulset.StatefulSetSet{}).
		Returns(200, "OK", statefulset.StatefulSetSet{}).
		Reads(statefulset.DescribeStatefulSetRequest{}).
		Writes(statefulset.CreateStatefulSetRequest{}).
		Returns(200, "OK", statefulset.CreateStatefulSetRequest{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
