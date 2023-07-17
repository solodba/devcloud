package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/tree/main/mpaas/apps/test"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc test.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return test.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Restful中间件测试"}
	// webservice定义路由信息
	// 定义测试路由
	ws.Route(ws.GET("/").To(h.QueryTest).
		Doc("restful中间件测试").
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
