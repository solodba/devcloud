package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/statefulset"
	"github.com/solodba/mcube/response"
)

// 创建StatefulSet
func (h *handler) CreateStatefulSet(r *restful.Request, w *restful.Response) {
	req := statefulset.NewCreateStatefulSetRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	statefulset, err := h.svc.CreateStatefulSet(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, statefulset))
}

// 删除StatefulSet
func (h *handler) DeleteStatefulSet(r *restful.Request, w *restful.Response) {
	req := statefulset.NewDeleteStatefulSetRequestFromRestful(r)
	statefulset, err := h.svc.DeleteStatefulSet(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, statefulset))
}

// 修改StatefulSet
func (h *handler) UpdateStatefulSet(r *restful.Request, w *restful.Response) {
	req := statefulset.NewUpdateStatefulSetRequest()
	err := r.ReadEntity(req.StatefulSet)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	statefulset, err := h.svc.UpdateStatefulSet(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, statefulset))
}

// 查询StatefulSet
func (h *handler) QueryOrDescribeStatefulSet(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := statefulset.NewQueryStatefulSetRequestFromRestful(r)
		statefulsetSet, err := h.svc.QueryStatefulSet(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, statefulsetSet))
	} else {
		req := statefulset.NewDescribeStatefulSetRequestFromRestful(r)
		statefulset, err := h.svc.DescribeStatefulSet(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, statefulset))
	}
}
