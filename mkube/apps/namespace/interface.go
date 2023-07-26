package namespace

// 模块名称
const (
	AppName = "namespace"
)

// Namespace业务接口
type Service interface {
	// 嵌套Namespace GRPC业务接口
	RPCServer
}

// QueryNamespaceRequest结构体初始化函数
func NewQueryNamespaceRequest() *QueryNamespaceRequest {
	return &QueryNamespaceRequest{}
}
