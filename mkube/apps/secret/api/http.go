package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/secret"
	"github.com/solodba/mcube/apps"
)

// handler结构体
type handler struct {
	svc secret.Service
}

// 实现注册restful实例Name方法
func (h *handler) Name() string {
	return secret.AppName
}

// 实现注册restful实例Conf方法
func (h *handler) Conf() error {
	// 从IOC获取Service实例
	h.svc = apps.GetInternalApp(secret.AppName).(secret.Service)
	return nil
}

// 实现注册restful实例version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现注册restful实例Registry方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"Secret管理"}
	// webservice定义路由信息
	// 创建Secret
	ws.Route(ws.POST("/").To(h.CreateSecret).
		Doc("创建Secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(secret.CreateSecretRequest{}).
		Writes(secret.Secret{}).
		Returns(200, "OK", secret.Secret{}))

	// 删除Secret
	ws.Route(ws.DELETE("/{namespace}/{name}").To(h.DeleteSecret).
		Doc("删除Secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(secret.DeleteSecretRequest{}).
		Writes(secret.SecretSetItem{}).
		Returns(200, "OK", secret.SecretSetItem{}))

	// 更新Secret
	ws.Route(ws.PUT("/").To(h.UpdateSecret).
		Doc("更新Secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(secret.UpdateSecretRequest{}).
		Writes(secret.Secret{}).
		Returns(200, "OK", secret.Secret{}))

	// 查询Secret
	ws.Route(ws.GET("/{namespace}").To(h.QueryOrDescribeSecret).
		Doc("查询Secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(secret.QuerySecretRequest{}).
		Writes(secret.SecretSet{}).
		Returns(200, "OK", secret.SecretSet{}).
		Reads(secret.DescribeSecretRequest{}).
		Writes(secret.SecretSetItem{}).
		Returns(200, "OK", secret.SecretSetItem{}))
}

// 初始化函数注册restful实例
func init() {
	apps.RegistryRestfulApp(&handler{})
}
