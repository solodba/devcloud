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

// QueryResourceRequest初始化函数
func NewQueryResourceRequest() *QueryResourceRequest {
	return &QueryResourceRequest{}
}

// QueryClusterInfoRequest初始化函数
func NewQueryClusterInfoRequest() *QueryClusterInfoRequest {
	return &QueryClusterInfoRequest{}
}
