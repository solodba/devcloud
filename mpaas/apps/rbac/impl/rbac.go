package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/rbac"
)

// 创建或者更新ServcieAccount
func (i *impl) CreateOrUpdateServiceAccount(ctx context.Context, in *rbac.ServiceAccount) (*rbac.Response, error) {
	return nil, nil
}

// 删除ServiceAccount
func (i *impl) DeleteServiceAccount(ctx context.Context, in *rbac.DeleteServiceAccountRequest) (*rbac.Response, error) {
	return nil, nil
}

// 查询ServiceAccount集合
func (i *impl) GetServiceAccountList(ctx context.Context, in *rbac.GetServiceAccountListRequest) (*rbac.ServiceAccountSet, error) {
	return nil, nil
}

// 创建或者更新ClusterRole
func (i *impl) CreateOrUpdateClusterRole(ctx context.Context, in *rbac.Role) (*rbac.Response, error) {
	return nil, nil
}

// 删除ClusterRole
func (i *impl) DeleteClusterRole(ctx context.Context, in *rbac.DeleteClusterRoleRequest) (*rbac.Response, error) {
	return nil, nil
}

// 查询ClusterRole集合
func (i *impl) GetClusterRoleList(ctx context.Context, in *rbac.GetClusterRoleListRequest) (*rbac.ClusterRoleSet, error) {
	return nil, nil
}

// 查询ClusterRole详情
func (i *impl) GetClusterRoleDetail(ctx context.Context, in *rbac.GetClusterRoleDetailRequest) (*rbac.Role, error) {
	return nil, nil
}

// 创建Role
func (i *impl) CreateOrUpdateRole(ctx context.Context, in *rbac.Role) (*rbac.Response, error) {
	return nil, nil
}

// 删除Role
func (i *impl) DeleteRole(ctx context.Context, in *rbac.DeleteRoleRequest) (*rbac.Response, error) {
	return nil, nil
}

// 查询Role集合
func (i *impl) GetRoleList(ctx context.Context, in *rbac.GetRoleListRequest) (*rbac.RoleSet, error) {
	return nil, nil
}

// 查询Role详情
func (i *impl) GetRoleDetail(ctx context.Context, in *rbac.GetRoleDetailRequest) (*rbac.Role, error) {
	return nil, nil
}

// 创建RoleBinding
func (i *impl) CreateOrUpdateRb(ctx context.Context, in *rbac.RoleBinding) (*rbac.Response, error) {
	return nil, nil
}

// 删除RoleBinding
func (i *impl) DeleteRb(ctx context.Context, in *rbac.DeleteRbRequest) (*rbac.Response, error) {
	return nil, nil
}

// 查看RoleBinding列表
func (i *impl) GetRbList(ctx context.Context, in *rbac.GetRbListRequest) (*rbac.RoleBindingSet, error) {
	return nil, nil
}

// 查询RoleBinding详情
func (i *impl) GetRbDetail(ctx context.Context, in *rbac.GetRbDetailRequest) (*rbac.RoleBinding, error) {
	return nil, nil
}
