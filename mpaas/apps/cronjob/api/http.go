package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/cronjob"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc cronjob.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return cronjob.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(cronjob.AppName).(cronjob.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"CronJob管理"}
	// webservice定义路由信息
	// 创建CronJob
	ws.Route(ws.POST("/").To(h.CreateCronJob).
		Doc("创建CronJob").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(cronjob.CreateCronJobRequest{}).
		Writes(cronjob.CronJob{}).
		Returns(200, "OK", cronjob.CronJob{}))

	// 删除CronJob
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteCronJob).
		Doc("删除CronJob").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(cronjob.DeleteCronJobRequest{}).
		Writes(cronjob.CreateCronJobRequest{}).
		Returns(200, "OK", cronjob.CreateCronJobRequest{}))

	// 更新CronJob
	ws.Route(ws.PUT("/").To(h.UpdateCronJob).
		Doc("更新CronJob").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(cronjob.UpdateCronJobRequest{}).
		Writes(cronjob.CronJob{}).
		Returns(200, "OK", cronjob.CronJob{}))

	// 查询CronJob
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeCronJob).
		Doc("查询CronJob").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(cronjob.QueryCronJobRequest{}).
		Writes(cronjob.CronJobSet{}).
		Returns(200, "OK", cronjob.CronJobSet{}).
		Reads(cronjob.DescribeCronJobRequest{}).
		Writes(cronjob.CreateCronJobRequest{}).
		Returns(200, "OK", cronjob.CreateCronJobRequest{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
