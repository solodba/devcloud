package user

import context "context"

// 业务模块名称
const (
	AppName = "user"
)

// 定义业务接口
type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	// 更新用户
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// GRPC业务接口
	RPCServer
}
