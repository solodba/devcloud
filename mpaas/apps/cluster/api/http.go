package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/cluster"
	"github.com/solodba/mcube/apps"
)

// 实例类
type handler struct {
	svc cluster.Service
}

// 实现Ioc的Name方法
func (h *handler) Name() string {
	return cluster.AppName
}

// 实现Ioc的Conf方法
func (h *handler) Conf() error {
	h.svc = apps.GetInternalApp(cluster.AppName).(cluster.Service)
	return nil
}

// 实现Ioc的Version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现Ioc的RegistryHandler方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"基于K8S的CI/CD平台"}
	ws.Route(ws.GET("/").
		To(h.QueryCluster).
		Doc("查询集群信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

// 注册函数
func init() {
	apps.RegistryRestfulApp(&handler{})
}
