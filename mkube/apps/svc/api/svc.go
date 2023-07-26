package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/svc"
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
	req := svc.NewDefaultUpdateServiceRequest()
	err := r.ReadEntity(req.Service)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	service, err := h.svc.UpdateService(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, service))
}

// 创建Service
func (h *handler) QueryOrDescribeService(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := svc.NewQueryServiceRequestFromRestful(r)
		serviceSet, err := h.svc.QueryService(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, serviceSet))
	} else {
		req := svc.NewDescribeServiceRequestFromRestful(r)
		service, err := h.svc.DescribeService(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, service))
	}
}
