package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/namespace"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc namespace.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return namespace.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(namespace.AppName).(namespace.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Namespace管理"}
	// webservice定义路由信息
	// 查询ConfigMap
	ws.Route(ws.GET("/").To(h.QueryNamespace).
		Doc("查询Namespace").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(namespace.QueryNamespaceRequest{}).
		Writes(namespace.NamespaceSet{}).
		Returns(200, "OK", namespace.NamespaceSet{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
