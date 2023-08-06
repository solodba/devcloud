package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/solodba/devcloud/mkube/apps/metrics"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 查询集群资源使用情况
func (i *impl) QueryClusterUsage(ctx context.Context, in *metrics.QueryClusterUsageRequest) (*metrics.MetricSet, error) {
	metricSet := metrics.NewMetricSet()
	raw, err := i.clientSet.RESTClient().Get().AbsPath(QUERY_CLUSTER_USAGE).DoRaw(ctx)
	if err != nil {
		return metricSet, err
	}
	nodeMetricSet := metrics.NewNodeMetricSet()
	err = json.Unmarshal(raw, nodeMetricSet)
	if err != nil {
		return metricSet, err
	}
	nodeList, err := i.clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return metricSet, err
	}
	if len(nodeList.Items) != len(nodeMetricSet.Items) {
		return metricSet, err
	}
	var cpuUsage, cpuTotal int64
	var memUsage, memTotal int64
	podList, err := i.clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return metricSet, err
	}
	var podUsage, podTotal int64 = int64(len(podList.Items)), 0
	for i, item := range nodeList.Items {
		cpuUsage += nodeMetricSet.Items[i].Usage.Cpu().MilliValue()
		cpuTotal += item.Status.Capacity.Cpu().MilliValue()
		memUsage += nodeMetricSet.Items[i].Usage.Memory().Value()
		memTotal += item.Status.Capacity.Memory().Value()
		podTotal += item.Status.Capacity.Pods().Value()
	}
	podUsageFormat := fmt.Sprintf("%.2f", (float64(podUsage)/float64(podTotal))*100)
	podMetricItem := metrics.NewMetricItem()
	podMetricItem.Value = podUsageFormat
	podMetricItem.Title = "Pod使用占比"
	cpuUsageFormat := fmt.Sprintf("%.2f", (float64(cpuUsage)/float64(cpuTotal))*100)
	cpuMetricItem := metrics.NewMetricItem()
	cpuMetricItem.Value = cpuUsageFormat
	cpuMetricItem.Title = "CPU使用占比"
	cpuMetricItem.Label = "cluster_cpu"
	memUsageFormat := fmt.Sprintf("%.2f", (float64(memUsage)/float64(memTotal))*100)
	memMetricItem := metrics.NewMetricItem()
	memMetricItem.Value = memUsageFormat
	memMetricItem.Title = "Memory使用占比"
	memMetricItem.Label = "cluster_mem"
	metricSet.AddItems(podMetricItem, cpuMetricItem, memMetricItem)
	metricSet.Total = int64(len(metricSet.Items))
	// 写入数据库
	_, err = i.col.InsertOne(ctx, metricSet)
	if err != nil {
		return nil, err
	}
	return metricSet, nil
}
