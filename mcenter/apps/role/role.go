package role

import "github.com/solodba/mcube/pb/meta"

// Role构造函数
func NewRole(req *CreateRoleRequest) *Role {
	return &Role{
		Meta: meta.NewMeta(),
		Spec: req,
	}
}
