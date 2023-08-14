package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/policy"
)

// 创建Policy
func (i *impl) CreatePolicy(ctx context.Context, in *policy.CreatePolicyRequest) (*policy.Policy, error) {
	policyIns := policy.NewPolicy(in)
	_, err := i.col.InsertOne(ctx, policyIns)
	if err != nil {
		return nil, err
	}
	return policyIns, nil
}

// 查询Policy
func (i *impl) QueryPolicy(ctx context.Context, in *policy.QueryPolicyRequest) (*policy.PolicySet, error) {
	return nil, nil
}
