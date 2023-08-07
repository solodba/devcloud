package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/ingroute"
	"github.com/solodba/mcube/response"
)

// 创建IngressRoute
func (h *handler) CreateIngressRoute(r *restful.Request, w *restful.Response) {
	req := ingroute.NewCreateIngressRouteRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	ingressRoute, err := h.svc.CreateIngressRoute(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, ingressRoute))
}

// 更新IngressRoute
func (h *handler) UpdateIngressRoute(r *restful.Request, w *restful.Response) {
	req := ingroute.NewUpdateIngressRouteRequest()
	err := r.ReadEntity(req.IngressRoute)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	ingressRoute, err := h.svc.UpdateIngressRoute(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, ingressRoute))
}

// 删除IngressRoute
func (h *handler) DeleteIngressRoute(r *restful.Request, w *restful.Response) {
	req := ingroute.NewDeleteIngressRouteRequestFromRestful(r)
	ingressRoute, err := h.svc.DeleteIngressRoute(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, ingressRoute))
}

// 查询IngressRoute
func (h *handler) QueryOrDescribeIngressRoute(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := ingroute.NewQueryIngressRouteRequestFromRestful(r)
		ingressRouteSet, err := h.svc.QueryIngressRoute(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, ingressRouteSet))
	} else {
		req := ingroute.NewDescribeIngressRouteRequestFromRestful(r)
		ingressRoute, err := h.svc.DescribeIngressRoute(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, ingressRoute))
	}
}

// 查询IngressRoute中间件列表
func (h *handler) QueryIngRouteMiddlewareList(r *restful.Request, w *restful.Response) {
	req := ingroute.NewQueryIngRouteMwRequestFromRestful(r)
	mwList, err := h.svc.QueryIngRouteMiddlewareList(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, mwList))
}
