package policy

import (
	"github.com/solodba/devcloud/mcenter/apps/role"
	"github.com/solodba/mcube/pb/meta"
)

// Policy构造函数
func NewPolicy(req *CreatePolicyRequest) *Policy {
	return &Policy{
		Meta: meta.NewMeta(),
		Spec: req,
		Role: role.NewDefaultRole(),
	}
}

// PolicySet构造函数
func NewPolicySet() *PolicySet {
	return &PolicySet{
		Items: make([]*Policy, 0),
	}
}

// PolicySet添加犯方法
func (p *PolicySet) AddItems(items ...*Policy) {
	p.Items = append(p.Items, items...)
}

// policy默认构造函数
func NewDefaultPolicy() *Policy {
	return &Policy{
		Meta: meta.NewMeta(),
		Spec: NewCreatePolicyRequest(),
		Role: role.NewDefaultRole(),
	}
}

// 通过PolicySet返回role id
func (p *PolicySet) RoleIds() []string {
	roleIds := make([]string, 0)
	for _, item := range p.Items {
		roleIds = append(roleIds, item.Spec.RoleId)
	}
	return roleIds
}

// PolicySet添加方法
func (p *PolicySet) Len() int64 {
	return int64(len(p.Items))
}

// PolicySet添加方法
func (p *PolicySet) SetRole(role *role.Role) {
	for _, item := range p.Items {
		if item.Spec.RoleId == role.Meta.Id {
			item.Role = role
		}
	}
}
