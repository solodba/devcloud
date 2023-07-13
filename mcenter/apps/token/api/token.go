package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/devcloud/tree/main/mcenter/common/response"
)

// 颁发令牌
func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {
	req := token.NewIssueTokenRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	tk, err := h.svc.IssueToken(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, tk))
}
