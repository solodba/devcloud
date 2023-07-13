package impl

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/user"
)

// 创建用户
func (i *impl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	return nil, nil
}

// 删除用户
func (i *impl) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.User, error) {
	return nil, nil
}

// 更新用户
func (i *impl) UpdateUser(ctx context.Context, in *user.UpdateUserRequest) (*user.User, error) {
	return nil, nil
}

// 查询用户
func (i *impl) Queryuser(ctx context.Context, in *user.QueryUserRequest) (*user.UserSet, error) {
	return nil, nil
}

// 查询用户详情
func (i *impl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}
