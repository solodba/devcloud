package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/cronjob"
	"github.com/solodba/mcube/response"
)

// 创建CronJob
func (h *handler) CreateCronJob(r *restful.Request, w *restful.Response) {
	req := cronjob.NewCreateCronJobRequest()
	err := r.ReadEntity(req)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	cronjob, err := h.svc.CreateCronJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, cronjob))
}

// 删除CronJob
func (h *handler) DeleteCronJob(r *restful.Request, w *restful.Response) {
	req := cronjob.NewDeleteCronJobRequestFromRestful(r)
	cronjob, err := h.svc.DeleteCronJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, cronjob))
}

// 修改CronJob
func (h *handler) UpdateCronJob(r *restful.Request, w *restful.Response) {
	req := cronjob.NewUpdateCronJobRequest()
	err := r.ReadEntity(req.CronJob)
	if err != nil {
		w.WriteEntity(response.NewFail(400, err.Error()))
		return
	}
	cronjob, err := h.svc.UpdateCronJob(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	w.WriteEntity(response.NewSuccess(200, cronjob))
}

// 查询CronJob
func (h *handler) QueryOrDescribeCronJob(r *restful.Request, w *restful.Response) {
	if r.QueryParameter("name") == "" {
		req := cronjob.NewQueryCronJobRequestFromRestful(r)
		cronJobSet, err := h.svc.QueryCronJob(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, cronJobSet))
	} else {
		req := cronjob.NewDescribeCronJobRequestFromRestful(r)
		cronjob, err := h.svc.DescribeCronJob(r.Request.Context(), req)
		if err != nil {
			w.WriteEntity(response.NewFail(500, err.Error()))
			return
		}
		w.WriteEntity(response.NewSuccess(200, cronjob))
	}
}
