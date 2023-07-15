package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/tree/main/mpaas/apps/pod"
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
	tags := []string{"登录"}
	// webservice定义路由信息
	ws.Route(ws.POST("/").To(h.CreatePod).
		Doc("创建Pod").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(pod.CreatePodRequest{}).
		Writes(pod.Pod{}).
		Returns(200, "OK", pod.Pod{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
