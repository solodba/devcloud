package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/policy"
	"github.com/solodba/devcloud/mcenter/apps/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	filter := bson.M{}
	if in.Namespace != "" {
		filter["namespace"] = in.Namespace
	}
	if in.UserId != "" {
		filter["user_id"] = in.UserId
	}
	opts := &options.FindOptions{}
	opts.SetLimit(in.Page.PageSize)
	in.Page.ComputeOffSet()
	opts.SetSkip(in.Page.Offset)
	policyInsSet := policy.NewPolicySet()
	cursor, err := i.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		policyIns := policy.NewDefaultPolicy()
		err = cursor.Decode(policyIns)
		if err != nil {
			return nil, err
		}
		policyInsSet.AddItems(policyIns)
	}
	if in.WithRole {
		req := role.NewQueryRoleRequest()
		req.AddItems(policyInsSet.RoleIds()...)
		req.Page.PageSize = policyInsSet.Len()
		roleInsSet, err := i.svc.QueryRole(ctx, req)
		if err != nil {
			return nil, err
		}
		for _, item := range roleInsSet.Items {
			policyInsSet.SetRole(item)
		}
	}
	policyInsSet.Total, err = i.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return policyInsSet, nil
}
