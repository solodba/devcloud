package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/maudit/apps/audit"
	"github.com/solodba/mcube/apps"
)

// 实例类
type handler struct {
	svc audit.Service
}

// 实现Ioc的Name方法
func (h *handler) Name() string {
	return audit.AppName
}

// 实现Ioc的Conf方法
func (h *handler) Conf() error {
	h.svc = apps.GetInternalApp(audit.AppName).(audit.Service)
	return nil
}

// 实现Ioc的Version方法
func (h *handler) Version() string {
	return "v1"
}

// 实现Ioc的RegistryHandler方法
func (h *handler) RegistryHandler(ws *restful.WebService) {
	tags := []string{"用户审计日志管理"}
	ws.Route(ws.GET("/").
		To(h.QueryAuditLog).
		Doc("查询审计日志").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(audit.QueryAuditLogRequest{}).
		Writes(audit.AuditLogSet{}).
		Returns(200, "OK", audit.AuditLogSet{}))
}

// 注册函数
func init() {
	apps.RegistryRestfulApp(&handler{})
}
