package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/maudit/apps/audit"
	"github.com/solodba/mcube/response"
)

// 查询审计日志
func (h *handler) QueryAuditLog(r *restful.Request, w *restful.Response) {
	req := audit.NewQueryAuditLogRequestFromRestful(r)
	auditLogSet, err := h.svc.QueryAuditLog(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, auditLogSet))
}
