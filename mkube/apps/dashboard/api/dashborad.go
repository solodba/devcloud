package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/mcube/response"
)

// 查询集群资源使用情况
func (h *handler) QueryClusterUsage(r *restful.Request, w *restful.Response) {
	req := metrics.NewQueryClusterUsageRequest()
	metricSet, err := h.svc.QueryClusterUsage(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	resultMap := make(map[string]*metrics.MetricSet)
	resultMap["usage"] = metricSet
	w.WriteEntity(response.NewSuccess(200, resultMap))
}

// 查询K8S资源情况
func (h *handler) QueryResource(r *restful.Request, w *restful.Response) {
	req := metrics.NewQueryResourceRequest()
	metricSet, err := h.svc.QueryResource(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	resultMap := make(map[string]*metrics.MetricSet)
	resultMap["resource"] = metricSet
	w.WriteEntity(response.NewSuccess(200, resultMap))
}

// 获取集群信息
func (h *handler) QueryClusterInfo(r *restful.Request, w *restful.Response) {
	req := metrics.NewQueryClusterInfoRequest()
	metricSet, err := h.svc.QueryClusterInfo(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	resultMap := make(map[string]*metrics.MetricSet)
	resultMap["cluster"] = metricSet
	w.WriteEntity(response.NewSuccess(200, resultMap))
}
