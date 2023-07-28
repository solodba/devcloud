package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/harbor"
	"github.com/solodba/mcube/response"
)

// 查询Harbor Projects
func (h *handler) QueryProjects(r *restful.Request, w *restful.Response) {
	req, err := harbor.NewQueryProjectsRequestFromRestful(r)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	projects, err := h.svc.QueryProjects(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, projects))
}
