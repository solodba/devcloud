package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/configmap"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc configmap.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return configmap.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(configmap.AppName).(configmap.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"ConfigMap管理"}
	// webservice定义路由信息
	// 创建ConfigMap
	ws.Route(ws.POST("/").To(h.CreateConfigMap).
		Doc("创建ConfigMap").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(configmap.CreateConfigMapRequest{}).
		Writes(configmap.ConfigMap{}).
		Returns(200, "OK", configmap.ConfigMap{}))

	// 删除ConfigMap
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteConfigMap).
		Doc("删除ConfigMap").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(configmap.DeleteConfigMapRequest{}).
		Writes(configmap.CreateConfigMapRequest{}).
		Returns(200, "OK", configmap.CreateConfigMapRequest{}))

	// 更新ConfigMap
	ws.Route(ws.PUT("/").To(h.UpdateConfigMap).
		Doc("更新ConfigMap").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(configmap.UpdateConfigMapRequest{}).
		Writes(configmap.ConfigMap{}).
		Returns(200, "OK", configmap.ConfigMap{}))

	// 查询ConfigMap
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeConfigMap).
		Doc("查询ConfigMap").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(configmap.QueryConfigMapRequest{}).
		Writes(configmap.ConfigMapSet{}).
		Returns(200, "OK", configmap.ConfigMapSet{}).
		Reads(configmap.DescribeConfigMapRequest{}).
		Writes(configmap.ConfigMapSetItem{}).
		Returns(200, "OK", configmap.ConfigMapSetItem{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
