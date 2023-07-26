package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/sc"
	"github.com/solodba/mcube/response"
)

// 创建StorageClass
func (h *handler) CreateSC(r *restful.Request, w *restful.Response) {
	req := sc.NewCreateSCRequest()
	if err := r.ReadEntity(req); err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	sc, err := h.svc.CreateSC(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, sc))
}

// 删除StorageClass
func (h *handler) DeleteSC(r *restful.Request, w *restful.Response) {
	req := sc.NewDeleteSCRequestFromRestful(r)
	sc, err := h.svc.DeleteSC(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, sc))
}

// 查询StorageClass集合
func (h *handler) QuerySC(r *restful.Request, w *restful.Response) {
	req := sc.NewQuerySCRequestFromRestful(r)
	scSet, err := h.svc.QuerySC(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, scSet))
}
