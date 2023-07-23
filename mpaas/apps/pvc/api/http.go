package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/pvc"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc pvc.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return pvc.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(pvc.AppName).(pvc.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"PersistentVolumeClaim管理"}
	// webservice定义路由信息
	// 创建PersistentVolumeClaim
	ws.Route(ws.POST("/").To(h.CreatePVC).
		Doc("创建PersistentVolumeClaim").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(pvc.CreatePVCRequest{}).
		Writes(pvc.PVC{}).
		Returns(200, "OK", pvc.PVC{}))

	// 删除PersistentVolumeClaim
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeletePVC).
		Doc("删除PersistentVolumeClaim").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(pvc.DeletePVCRequest{}).
		Writes(pvc.PVC{}).
		Returns(200, "OK", pvc.PVC{}))

	// 查询PersistentVolumeClaim
	ws.Route(ws.GET("/{namespace}").To(h.QueryPVC).
		Doc("查询PersistentVolumeClaim").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(pvc.QueryPVCRequest{}).
		Writes(pvc.PVCSet{}).
		Returns(200, "OK", pvc.PVCSet{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
