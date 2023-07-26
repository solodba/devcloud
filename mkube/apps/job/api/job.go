package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/job"
	"github.com/solodba/mcube/response"
)

// 创建Job
func (h *handler) CreateJob(r *restful.Request, w *restful.Response) {
	req := job.NewCreateJobRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	job, err := h.svc.CreateJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, job))
}

// 删除Job
func (h *handler) DeleteJob(r *restful.Request, w *restful.Response) {
	req := job.NewDeleteJobRequestFromRestful(r)
	job, err := h.svc.DeleteJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, job))
}

// 修改Job
func (h *handler) UpdateJob(r *restful.Request, w *restful.Response) {
	req := job.NewUpdateJobRequest()
	err := r.ReadEntity(req.Job)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	job, err := h.svc.UpdateJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, job))
}

// 查询Job
func (h *handler) QueryOrDescribeJob(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := job.NewQueryJobRequestFromRestful(r)
		jobSet, err := h.svc.QueryJob(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, jobSet))
	} else {
		req := job.NewDescribeJobRequestFromRestful(r)
		job, err := h.svc.DescribeJob(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, job))
	}
}
