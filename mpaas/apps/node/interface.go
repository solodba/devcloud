package node

// 模块名称
const (
	AppName = "node"
)

// Node相关功能接口
type Service interface {
	// 嵌套Node GRPC接口
	RPCServer
}

// QueryNodeRequest结构体初始化函数
func NewQueryNodeRequest() *QueryNodeRequest {
	return &QueryNodeRequest{}
}
