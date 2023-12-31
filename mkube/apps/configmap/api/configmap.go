package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/configmap"
	"github.com/solodba/mcube/response"
)

// 创建ConfigMap
func (h *handler) CreateConfigMap(r *restful.Request, w *restful.Response) {
	req := configmap.NewCreateConfigMapRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	configmap, err := h.svc.CreateConfigMap(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, configmap))
}

// 删除ConfigMap
func (h *handler) DeleteConfigMap(r *restful.Request, w *restful.Response) {
	req := configmap.NewDeleteConfigMapRequestFromRestful(r)
	configmap, err := h.svc.DeleteConfigMap(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, configmap))
}

// 修改ConfigMap
func (h *handler) UpdateConfigMap(r *restful.Request, w *restful.Response) {
	req := configmap.NewUpdateConfigMapRequest(configmap.NewCreateConfigMapRequest())
	err := r.ReadEntity(req.ConfigMap)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	configmap, err := h.svc.UpdateConfigMap(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, configmap))
}

// 查询ConfigMap
func (h *handler) QueryOrDescribeConfigMap(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := configmap.NewQueryConfigMapRequestFromRestful(r)
		configmapSet, err := h.svc.QueryConfigMap(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, configmapSet))
	} else {
		req := configmap.NewDescribeConfigMapRequestFromRestful(r)
		configmap, err := h.svc.DescribeConfigMap(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, configmap))
	}
}
