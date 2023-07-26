package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/ingress"
	"github.com/solodba/mcube/response"
)

// 创建Ingress
func (h *handler) CreateIngress(r *restful.Request, w *restful.Response) {
	req := ingress.NewCreateIngressRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	ingress, err := h.svc.CreateIngress(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, ingress))
}

// 删除Ingress
func (h *handler) DeleteIngress(r *restful.Request, w *restful.Response) {
	req := ingress.NewDeleteIngressRequestFromRestful(r)
	ingress, err := h.svc.DeleteIngress(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, ingress))
}

// 修改Ingress
func (h *handler) UpdateIngress(r *restful.Request, w *restful.Response) {
	req := ingress.NewUpdateIngressRequest(ingress.NewCreateIngressRequest())
	err := r.ReadEntity(req.Ingress)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	ingress, err := h.svc.UpdateIngress(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, ingress))
}

// 查询Ingress
func (h *handler) QueryOrDescribeIngress(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := ingress.NewQueryIngressRequestFromRestful(r)
		ingressSet, err := h.svc.QueryIngress(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, ingressSet))
	} else {
		req := ingress.NewDescribeIngressRequestFromRestful(r)
		ingress, err := h.svc.DescribeIngress(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, ingress))
	}
}
