package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/rbac"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc rbac.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return rbac.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(rbac.AppName).(rbac.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"RBAC管理"}
	// webservice定义路由信息
	// 创建ServiceAccount
	ws.Route(ws.POST("/sa").To(h.CreateOrUpdateServiceAccount).
		Doc("创建ServiceAccount").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 删除ServiceAccount
	ws.Route(ws.DELETE("/sa/{namespace}/{name}").To(h.DeleteServiceAccount).
		Doc("删除ServiceAccount").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 查询ServiceAccount
	ws.Route(ws.GET("/sa/{namespace}").To(h.GetServiceAccountList).
		Doc("查询ServiceAccount").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 创建Role
	ws.Route(ws.POST("/role").To(h.CreateOrUpdateRole).
		Doc("创建Role").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 删除Role
	ws.Route(ws.DELETE("/role").To(h.DeleteRole).
		Doc("删除Role").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 查询Role
	ws.Route(ws.GET("/role").To(h.GetRoleDetailOrList).
		Doc("查询Role").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 创建RoleBinding
	ws.Route(ws.POST("/rb").To(h.CreateOrUpdateRb).
		Doc("创建RoleBinding").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 删除RoleBinding
	ws.Route(ws.DELETE("/rb").To(h.DeleteRb).
		Doc("创建RoleBinding").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 查询RoleBinding
	ws.Route(ws.GET("/rb").To(h.GetRbDetailOrList).
		Doc("查询RoleBinding").
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
