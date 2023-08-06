package metrics

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetricItem初始化函数
func NewMetricItem() *MetricItem {
	return &MetricItem{}
}

// MetricSet初始化函数
func NewMetricSet() *MetricSet {
	return &MetricSet{
		Items: make([]*MetricItem, 0),
	}
}

// MetricSet添加方法
func (m *MetricSet) AddItems(items ...*MetricItem) {
	m.Items = append(m.Items, items...)
}

// NodeMetrics结构体
type NodeMetric struct {
	Metadata  *metav1.ObjectMeta   `json:"metadata"`
	Timestamp *time.Time           `json:"timestamp"`
	Window    string               `json:"window"`
	Usage     *corev1.ResourceList `json:"usage"`
}

// NodeMetricSet结构体
type NodeMetricSet struct {
	Kind       string             `json:"kind"`
	ApiVersion string             `json:"apiVersion"`
	Metadata   *metav1.ObjectMeta `json:"metadata"`
	Items      []*NodeMetric      `json:"items"`
}

// NodeMetrics结构体初始化函数
func NewNodeMetric() *NodeMetric {
	return &NodeMetric{
		Metadata:  &metav1.ObjectMeta{},
		Timestamp: &time.Time{},
		Usage:     &corev1.ResourceList{},
	}
}

// NodeMetricSet结构体初始化函数
func NewNodeMetricSet() *NodeMetricSet {
	return &NodeMetricSet{
		Metadata: &metav1.ObjectMeta{},
		Items:    make([]*NodeMetric, 0),
	}
}

// NodeMetricSet结构体添加方法
func (n *NodeMetricSet) AddItems(items ...*NodeMetric) {
	n.Items = append(n.Items, items...)
}
