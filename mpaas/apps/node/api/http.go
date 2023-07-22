package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/node"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc node.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return node.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(node.AppName).(node.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Node管理"}
	// webservice定义路由信息
	// 查询Node
	ws.Route(ws.GET("/").To(h.QueryOrDescribeNode).
		Doc("查询Node").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(node.QueryNodeRequest{}).
		Writes(node.NodeSet{}).
		Returns(200, "OK", node.NodeSet{}).
		Reads(node.DescribeNodeRequest{}).
		Writes(node.NodeSetItem{}).
		Returns(200, "OK", node.NodeSetItem{}))

	// 更新节点标签
	ws.Route(ws.PUT("/label").To(h.UpdateNodeLabel).
		Doc("更新节点标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(node.UpdatedLabelRequest{}).
		Writes(node.UpdatedLabelResponse{}).
		Returns(200, "OK", node.UpdatedLabelResponse{}))

	// 更新节点标签
	ws.Route(ws.PUT("/taint").To(h.UpdateNodeLabel).
		Doc("更新节点污点").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(node.UpdatedTaintRequest{}).
		Writes(node.UpdatedTaintResponse{}).
		Returns(200, "OK", node.UpdatedTaintResponse{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
