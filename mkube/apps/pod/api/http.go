package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/pod"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc pod.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return pod.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取token实例
	h.svc = apps.GetInternalApp(pod.AppName).(pod.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Pod管理"}
	// webservice定义路由信息
	// 创建Pod
	ws.Route(ws.POST("/").To(h.CreatePod).
		Doc("创建Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(pod.CreatePodRequest{}).
		Writes(pod.Pod{}).
		Returns(200, "OK", pod.Pod{}))

	// 删除Pod
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeletePod).
		Doc("删除Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(pod.DeletePodRequest{}).
		Writes(pod.Pod{}).
		Returns(200, "OK", pod.Pod{}))

	// 更新Pod
	ws.Route(ws.PUT("/").To(h.UpdatePod).
		Doc("更新Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(pod.UpdatePodRequest{}).
		Writes(pod.Pod{}).
		Returns(200, "OK", pod.Pod{}))

	// 查询Pod
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribePod).
		Doc("查询Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 装饰路由, 是否开启权限认证
		Metadata("auth", true).
		// 装饰路由, 是否开启用户访问鉴权
		Metadata("perm", true).
		Reads(pod.QueryPodRequest{}).
		Writes(pod.PodSet{}).
		Returns(200, "OK", pod.PodSet{}).
		Reads(pod.DescribePodRequest{}).
		Writes(pod.Pod{}).
		Returns(200, "OK", pod.Pod{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
