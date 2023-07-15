package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/tree/main/mpaas/apps/pod"
	"github.com/solodba/mcube/response"
)

// 创建Pod
func (h *handler) CreatePod(r *restful.Request, w *restful.Response) {
	req := pod.NewCreatePodRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	pod, err := h.svc.CreatePod(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pod))
}
