package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/role"
)

// 创建role
func (i *impl) CreateRole(ctx context.Context, in *role.CreateRoleRequest) (*role.Role, error) {
	roleIns := role.NewRole(in)
	_, err := i.col.InsertOne(ctx, roleIns)
	if err != nil {
		return nil, err
	}
	return roleIns, nil
}

// 查询role
func (i *impl) QueryRole(ctx context.Context, in *role.QueryRoleRequest) (*role.RoleSet, error) {
	return nil, nil
}
