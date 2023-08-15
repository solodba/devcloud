package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/mcube/apps"
)

// 实例类
type handler struct {
	svc token.Service
}

// 实现Ioc的Name方法
func (h *handler) Name() string {
	return token.AppName
}

// 实现Ioc的Conf方法
func (h *handler) Conf() error {
	h.svc = apps.GetInternalApp(token.AppName).(token.Service)
	return nil
}

// 实现Ioc的Version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现Ioc的RegistryHandler方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"登录"}
	ws.Route(ws.POST("/").
		To(h.IssueToken).
		Doc("颁发令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// 路由装饰, 是否开启认证
		Metadata("auth", false).
		// 路由装饰, 是否开启鉴权
		Metadata("perm", false).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))
}

// 注册函数
func init() {
	apps.RegistryRestfulApp(&handler{})
}
