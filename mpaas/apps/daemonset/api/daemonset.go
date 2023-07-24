package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/daemonset"
	"github.com/solodba/mcube/response"
)

// 创建DaemonSet
func (h *handler) CreateDaemonSet(r *restful.Request, w *restful.Response) {
	req := daemonset.NewCreateDaemonSetRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	configmap, err := h.svc.CreateDaemonSet(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, configmap))
}

// 删除DaemonSet
func (h *handler) DeleteDaemonSet(r *restful.Request, w *restful.Response) {
	req := daemonset.NewDeleteDaemonSetRequestFromRestful(r)
	daemonset, err := h.svc.DeleteDaemonSet(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, daemonset))
}

// 修改DaemonSet
func (h *handler) UpdateDaemonSet(r *restful.Request, w *restful.Response) {
	req := daemonset.NewUpdateDaemonSetRequest()
	err := r.ReadEntity(req.DaemonSet)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	daemonset, err := h.svc.UpdateDaemonSet(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, daemonset))
}

// 查询Deployment
func (h *handler) QueryOrDescribeDaemonSet(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := daemonset.NewQueryDaemonSetRequestFromRestful(r)
		daemonsetList, err := h.svc.QueryDaemonSet(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, daemonsetList))
	} else {
		req := daemonset.NewDescribeDaemonSetRequestFromRestful(r)
		daemonset, err := h.svc.DescribeDaemonSet(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, daemonset))
	}
}
