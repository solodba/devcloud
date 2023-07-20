package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/secret"
	"github.com/solodba/mcube/response"
)

// 创建Secret
func (h *handler) CreateSecret(r *restful.Request, w *restful.Response) {
	req := secret.NewCreateSecretRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	secret, err := h.svc.CreateSecret(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, secret))
}

// 删除Secret
func (h *handler) DeleteSecret(r *restful.Request, w *restful.Response) {

}

// 更新Secret
func (h *handler) UpdateSecret(r *restful.Request, w *restful.Response) {
	req := secret.NewUpdateSecretRequest(secret.NewCreateSecretRequest())
	err := r.ReadEntity(req.Secret)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	secret, err := h.svc.UpdateSecret(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, secret))
}

// 查询Secret
func (h *handler) QueryOrDescribeSecret(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := secret.NewQuerySecretRequestFromRestful(r)
		secretSet, err := h.svc.QuerySecret(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, secretSet))
	} else {
		req := secret.NewDescribeSecretRequestFromRestful(r)
		secret, err := h.svc.DescribeSecret(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, secret))
	}
}
