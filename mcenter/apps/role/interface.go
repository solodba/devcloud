package role

import context "context"

// 模块名称
const (
	AppName = "role"
)

// Role相关业务接口
type Service interface {
	// 创建role
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	// 嵌套role grpc接口
	RPCServer
}
