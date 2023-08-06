package metrics

// 模块名称
const (
	AppName = "metrics"
)

// metrics服务接口
type Service interface {
	RPCServer
}

// QueryClusterUsageRequest初始化函数
func NewQueryClusterUsageRequest() *QueryClusterUsageRequest {
	return &QueryClusterUsageRequest{}
}
