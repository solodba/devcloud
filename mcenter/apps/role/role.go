package role

import "github.com/solodba/mcube/pb/meta"

// Role构造函数
func NewRole(req *CreateRoleRequest) *Role {
	return &Role{
		Meta: meta.NewMeta(),
		Spec: req,
	}
}

// RoleSet构造函数
func NewRoleSet() *RoleSet {
	return &RoleSet{
		Items: make([]*Role, 0),
	}
}

// RoleSet添加方法
func (r *RoleSet) AddItems(items ...*Role) {
	r.Items = append(r.Items, items...)
}

// Role默认构造函数
func NewDefaultRole() *Role {
	return &Role{
		Meta: meta.NewMeta(),
		Spec: NewCreateRoleRequest(),
	}
}
