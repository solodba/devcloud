package collector

import "github.com/prometheus/client_golang/prometheus"

// 模块名称
const (
	AppName = "collector"
)

// Metrics收集器接口
type Service interface {
	// 嵌套prometheus收集器接口
	prometheus.Collector
}
