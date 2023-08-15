package role

import (
	"github.com/solodba/mcube/pb/meta"
)

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

// Role添加鉴权方法
func (r *Role) HasPermission(ServiceId, HttpMethod, HttpPath string) bool {
	for _, item := range r.Spec.Feature {
		if item.IsEqual(ServiceId, HttpMethod, HttpPath) {
			return true
		}
	}
	return false
}

// Feature添加判断方法
func (f *Feature) IsEqual(ServiceId, HttpMethod, HttpPath string) bool {
	return f.ServiceId == ServiceId && f.HttpMethod == HttpMethod && f.HttpPath == HttpPath
}
