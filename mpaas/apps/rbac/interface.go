package rbac

// 模块名称
const (
	AppName = "rbac"
)

// RBAC相关功能接口
type Service interface {
	// 嵌套RBAC GRPC接口
	RPCServer
}
