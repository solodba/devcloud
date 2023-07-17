package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	"github.com/solodba/mcube/response"
)

// 创建Pod
func (h *handler) CreatePod(r *restful.Request, w *restful.Response) {
	req := pod.NewCreatePodRequest()
	err := r.ReadEntity(req.Pod)
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

// 删除Pod
func (h *handler) DeletePod(r *restful.Request, w *restful.Response) {
	req := pod.NewDeletePodRequestFromRestful(r)
	pod, err := h.svc.DeletePod(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pod))
}

// 更新Pod
func (h *handler) UpdatePod(r *restful.Request, w *restful.Response) {
	podReq := pod.NewDefaultPod()
	err := r.ReadEntity(podReq)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	req := pod.NewUpdatePodRequest(podReq)
	pod, err := h.svc.UpdatePod(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, pod))
}

// 查询Pod
func (h *handler) QueryOrDescribePod(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") != "" {
		req := pod.NewDescribePodRequestFromRestful(r)
		pod, err := h.svc.DescribePod(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, pod))
	} else {
		req := pod.NewQueryPodRequestFromRestful(r)
		podSet, err := h.svc.QueryPod(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, podSet))
	}
}
