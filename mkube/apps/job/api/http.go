package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/job"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc job.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return job.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(job.AppName).(job.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Job管理"}
	// webservice定义路由信息
	// 创建Job
	ws.Route(ws.POST("/").To(h.CreateJob).
		Doc("创建Job").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(job.CreateJobRequest{}).
		Writes(job.Job{}).
		Returns(200, "OK", job.Job{}))

	// 删除Job
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteJob).
		Doc("删除Job").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(job.DeleteJobRequest{}).
		Writes(job.CreateJobRequest{}).
		Returns(200, "OK", job.CreateJobRequest{}))

	// 更新Job
	ws.Route(ws.PUT("/").To(h.UpdateJob).
		Doc("更新Job").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(job.UpdateJobRequest{}).
		Writes(job.Job{}).
		Returns(200, "OK", job.Job{}))

	// 查询Job
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeJob).
		Doc("查询Job").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(job.QueryJobRequest{}).
		Writes(job.JobSet{}).
		Returns(200, "OK", job.JobSet{}).
		Reads(job.DescribeJobRequest{}).
		Writes(job.CreateJobRequest{}).
		Returns(200, "OK", job.CreateJobRequest{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
