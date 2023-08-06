package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/devcloud/mkube/apps/prometheus"
	"github.com/solodba/mcube/response"
)

// 获取集群时间段变化数据
func (h *handler) QueryClusterUsageRange(r *restful.Request, w *restful.Response) {
	req := prometheus.NewQueryClusterUsageRangeRequest()
	metricSet, err := h.svc.QueryClusterUsageRange(r.Request.Context(), req)
	if err != nil {
		w.WriteEntity(response.NewFail(500, err.Error()))
		return
	}
	resultMap := make(map[string]*metrics.MetricSet)
	resultMap["usageRange"] = metricSet
	w.WriteEntity(response.NewSuccess(200, resultMap))
}
