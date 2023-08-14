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
