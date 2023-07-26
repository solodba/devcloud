package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/namespace"
	"github.com/solodba/mcube/response"
)

// 获取Namespace列表
func (h *handler) QueryNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewQueryNamespaceRequest()
	namespaceSet, err := h.svc.QueryNamespace(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, namespaceSet))
}
