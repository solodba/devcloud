package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/harbor"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc harbor.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return harbor.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(harbor.AppName).(harbor.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Harbor管理"}
	// webservice定义路由信息
	// 查询Harbor Projects
	ws.Route(ws.GET("/projects").To(h.QueryProjects).
		Doc("查询Harbor Projects").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(harbor.QueryProjectsRequest{}).
		Writes(harbor.Projects{}).
		Returns(200, "OK", harbor.Projects{}))

	// 查询Harbor Repository
	ws.Route(ws.GET("/projects/{projectName}").To(h.QueryRepositories).
		Doc("查询Harbor Repository").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(harbor.QueryRepositoriesRequest{}).
		Writes(harbor.Repositories{}).
		Returns(200, "OK", harbor.Repositories{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
