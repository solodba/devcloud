package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/configmap"
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

}

// 查询ConfigMap
func (h *handler) QueryOrDescribeConfigMap(r *restful.Request, w *restful.Response) {

}
