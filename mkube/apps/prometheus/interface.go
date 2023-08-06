package prometheus

// 模块名称
const (
	AppName = "prometheus"
)

// 从Prometheus获取数据接口
type Service interface {
	// 嵌套Prometheus GRPC接口
	RPCServer
}

// QueryMetricsFromPromRequest初始化函数
func NewQueryMetricsFromPromRequest(metricName string) *QueryMetricsFromPromRequest {
	return &QueryMetricsFromPromRequest{
		MetricName: metricName,
	}
}

// QueryClusterUsageRangeRequest初始化函数
func NewQueryClusterUsageRangeRequest() *QueryClusterUsageRangeRequest {
	return &QueryClusterUsageRangeRequest{}
}
