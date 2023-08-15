package permission

import "github.com/solodba/devcloud/mcenter/apps/role"

// CheckPermissionResponse构造函数
func NewCheckPermissionResponse() *CheckPermissionResponse {
	return &CheckPermissionResponse{
		Role: role.NewDefaultRole(),
	}
}
