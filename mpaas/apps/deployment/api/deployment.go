package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/deployment"
	"github.com/solodba/mcube/response"
)

// 创建Deployment
func (h *handler) CreateDeployment(r *restful.Request, w *restful.Response) {
	req := deployment.NewCreateDeploymentRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	deployment, err := h.svc.CreateDeployment(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, deployment))
}

// 删除Deployment
func (h *handler) DeleteDeployment(r *restful.Request, w *restful.Response) {
	req := deployment.NewDeleteDeploymentRequestFromRestful(r)
	deployment, err := h.svc.DeleteDeployment(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, deployment))
}

// 修改Deployment
func (h *handler) UpdateDeployment(r *restful.Request, w *restful.Response) {
	req := deployment.NewUpdateDeploymentRequest()
	err := r.ReadEntity(req.Deployment)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	deployment, err := h.svc.UpdateDeployment(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, deployment))
}

// 查询Deployment
func (h *handler) QueryOrDescribeDeployment(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := deployment.NewQueryDeploymentRequestFromRestful(r)
		deploymentSet, err := h.svc.QueryDeployment(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, deploymentSet))
	} else {
		req := deployment.NewDescribeDeploymentRequestFromRestful(r)
		deployment, err := h.svc.DescribeDeployment(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, deployment))
	}
}
