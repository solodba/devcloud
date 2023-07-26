package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/pvc"
	"github.com/solodba/mcube/response"
)

// 创建PersistentVolumeClaim
func (h *handler) CreatePVC(r *restful.Request, w *restful.Response) {
	req := pvc.NewCreatePVCRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	pvc, err := h.svc.CreatePVC(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pvc))
}

// 删除PersistentVolumeClaim
func (h *handler) DeletePVC(r *restful.Request, w *restful.Response) {
	req := pvc.NewDeletePVCRequestFromRestful(r)
	pvc, err := h.svc.DeletePVC(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pvc))
}

// 查询PersistentVolumeClaim集合
func (h *handler) QueryPVC(r *restful.Request, w *restful.Response) {
	req := pvc.NewQueryPVCRequestFromRestful(r)
	pvcSet, err := h.svc.QueryPVC(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pvcSet))
}
