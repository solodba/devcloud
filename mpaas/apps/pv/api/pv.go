package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/pv"
	"github.com/solodba/mcube/response"
)

// 创建PersistentVolume
func (h *handler) CreatePV(r *restful.Request, w *restful.Response) {
	req := pv.NewCreatePVRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	pv, err := h.svc.CreatePV(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pv))
}

// 删除PersistentVolume
func (h *handler) DeletePV(r *restful.Request, w *restful.Response) {
	req := pv.NewDeletePVRequestFromRestful(r)
	pv, err := h.svc.DeletePV(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pv))
}

// 查询PersistentVolume集合
func (h *handler) QueryPV(r *restful.Request, w *restful.Response) {

}
