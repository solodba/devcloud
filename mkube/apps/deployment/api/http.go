package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/deployment"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc deployment.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return deployment.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(deployment.AppName).(deployment.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Deployment管理"}
	// webservice定义路由信息
	// 创建Deployment
	ws.Route(ws.POST("/").To(h.CreateDeployment).
		Doc("创建Deployment").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(deployment.CreateDeploymentRequest{}).
		Writes(deployment.Deployment{}).
		Returns(200, "OK", deployment.Deployment{}))

	// 删除Deployment
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteDeployment).
		Doc("删除Deployment").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(deployment.DeleteDeploymentRequest{}).
		Writes(deployment.CreateDeploymentRequest{}).
		Returns(200, "OK", deployment.CreateDeploymentRequest{}))

	// 更新Deployment
	ws.Route(ws.PUT("/").To(h.UpdateDeployment).
		Doc("更新Deployment").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(deployment.UpdateDeploymentRequest{}).
		Writes(deployment.Deployment{}).
		Returns(200, "OK", deployment.Deployment{}))

	// 查询Deployment
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeDeployment).
		Doc("查询Deployment").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(deployment.QueryDeploymentRequest{}).
		Writes(deployment.DeploymentSet{}).
		Returns(200, "OK", deployment.DeploymentSet{}).
		Reads(deployment.DescribeDeploymentRequest{}).
		Writes(deployment.CreateDeploymentRequest{}).
		Returns(200, "OK", deployment.CreateDeploymentRequest{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
