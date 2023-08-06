package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/devcloud/mkube/apps/prometheus"
)

// 从Prometheus获取Metrics数据
func (i *impl) QueryMetricsFromProm(ctx context.Context, in *prometheus.QueryMetricsFromPromRequest) (*prometheus.PromMetrics, error) {
	if !i.prom.Enable {
		return prometheus.NewPromMetrics(""), fmt.Errorf("Prometheus is not enable")
	}
	resultMap := make(map[string][]string)
	now := time.Now()
	start, end := now.Add(-time.Hour*24), now
	r := promv1.Range{
		Start: start,
		End:   end,
		Step:  time.Minute * time.Duration(5),
	}
	queryRange, _, err := i.promapi.QueryRange(ctx, in.MetricName, r)
	if err != nil {
		return prometheus.NewPromMetrics(""), err
	}
	matrix := queryRange.(model.Matrix)
	if len(matrix) == 0 {
		return prometheus.NewPromMetrics(""), fmt.Errorf("Prometheus data is empty")
	}
	x := make([]string, 0)
	y := make([]string, 0)
	for _, value := range matrix[0].Values {
		format := value.Timestamp.Time().Format("15:04")
		x = append(x, format)
		y = append(y, value.Value.String())
	}
	resultMap["x"] = x
	resultMap["y"] = y
	raw, err := json.Marshal(resultMap)
	if err != nil {
		return prometheus.NewPromMetrics(""), fmt.Errorf("resultMap json Marshal error")
	}
	return prometheus.NewPromMetrics(string(raw)), nil
}

// 获取集群时间段变化数据
func (i *impl) QueryClusterUsageRange(ctx context.Context, in *prometheus.QueryClusterUsageRangeRequest) (*metrics.MetricSet, error) {
	metricSet := metrics.NewMetricSet()
	clusterCpuReq := prometheus.NewQueryMetricsFromPromRequest("cluster_cpu")
	promMetrics, err := i.QueryMetricsFromProm(ctx, clusterCpuReq)
	if err != nil {
		return metricSet, err
	}
	clusterCpuMetric := metrics.NewMetricItem()
	clusterCpuMetric.Title = "CPU变化趋势"
	clusterCpuMetric.Value = promMetrics.Metrics
	clusterMemReq := prometheus.NewQueryMetricsFromPromRequest("cluster_mem")
	promMetrics, err = i.QueryMetricsFromProm(ctx, clusterMemReq)
	if err != nil {
		return metricSet, err
	}
	clusterMemMetric := metrics.NewMetricItem()
	clusterMemMetric.Title = "Memory变化趋势"
	clusterMemMetric.Value = promMetrics.Metrics
	metricSet.AddItems(clusterCpuMetric, clusterMemMetric)
	metricSet.Total = int64(len(metricSet.Items))
	// 监控数据插入数据库
	_, err = i.col.InsertOne(ctx, metricSet)
	if err != nil {
		return metricSet, err
	}
	return metricSet, nil
}
