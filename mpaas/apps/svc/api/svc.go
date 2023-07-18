package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/svc"
	"github.com/solodba/mcube/response"
)

// 创建Service
func (h *handler) CreateService(r *restful.Request, w *restful.Response) {
	req := svc.NewCreateServiceRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	service, err := h.svc.CreateService(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, service))
}

// 创建Service
func (h *handler) DeleteService(r *restful.Request, w *restful.Response) {
	req := svc.NewDeleteServiceRequestFromRestful(r)
	service, err := h.svc.DeleteService(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, service))
}

// 创建Service
func (h *handler) UpdateService(r *restful.Request, w *restful.Response) {

}

// 创建Service
func (h *handler) QueryOrDescribeService(r *restful.Request, w *restful.Response) {

}
