package impl

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/solodba/devcloud/mkube/apps/collector"
	"github.com/solodba/devcloud/mkube/apps/metrics"
	"github.com/solodba/mcube/apps"
)

var (
	svc = &impl{}
)

// 定义Collector的service的实例类
type impl struct {
	// 引入metrics收集
	svc metrics.Service
	// CPU和Memory收集器
	cpuMemCollector *CpuMemCollector
}

// 定义CPU和Memory收集
type CpuMemCollector struct {
	// CPU收集指标
	clusterCpu prometheus.Gauge
	// Memory收集指标
	clusterMem prometheus.Gauge
}

// CpuMemCollector初始化函数
func NewCpuMemCollector() *CpuMemCollector {
	return &CpuMemCollector{
		clusterCpu: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "cluster_cpu",
				Help: "collector cluster cpu info",
			},
		),
		clusterMem: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "cluster_mem",
				Help: "collector cluster memory info",
			},
		),
	}
}

// 实现注册内部实例注册Name方法
func (i *impl) Name() string {
	return collector.AppName
}

// 实现内部实例注册Conf方法
func (i *impl) Conf() error {
	// 初始化svc
	i.svc = apps.GetInternalApp(metrics.AppName).(metrics.Service)
	return nil
}

// 注册user模块实例
func init() {
	// 注册内部实例到IOC托管
	apps.RegistryInternalApp(svc)
	// 注册实例到prometheus
	svc.cpuMemCollector = NewCpuMemCollector()
	prometheus.MustRegister(svc)
}
