package impl

import (
	"context"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/solodba/devcloud/mkube/apps/metrics"
)

// 实现接口
func (i *impl) Describe(desc chan<- *prometheus.Desc) {
	i.cpuMemCollector.clusterCpu.Describe(desc)
	i.cpuMemCollector.clusterMem.Describe(desc)
}

func (i *impl) Collect(metricsName chan<- prometheus.Metric) {
	req := metrics.NewQueryClusterUsageRequest()
	ctx := context.TODO()
	metricSet, err := i.svc.QueryClusterUsage(ctx, req)
	if err != nil {
		return
	}
	for _, item := range metricSet.Items {
		switch item.Label {
		case "cluster_cpu":
			newValue, _ := strconv.ParseFloat(item.Value, 64)
			i.cpuMemCollector.clusterCpu.Set(newValue)
			i.cpuMemCollector.clusterCpu.Collect(metricsName)
		case "cluster_mem":
			newValue, _ := strconv.ParseFloat(item.Value, 64)
			i.cpuMemCollector.clusterMem.Set(newValue)
			i.cpuMemCollector.clusterMem.Collect(metricsName)
		}
	}
}
