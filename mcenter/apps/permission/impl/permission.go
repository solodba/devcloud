package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/permission"
	"github.com/solodba/devcloud/mcenter/apps/policy"
)

// 用户访问鉴权
func (i *impl) CheckPermission(ctx context.Context, in *permission.CheckPermissionRequest) (*permission.CheckPermissionResponse, error) {
	policyReq := policy.NewQueryPolicyRequest()
	policyReq.UserId = in.UserId
	policyReq.Namespace = in.Namespace
	policyReq.WithRole = true
	policySet, err := i.svc.QueryPolicy(ctx, policyReq)
	if err != nil {
		return nil, err
	}
	permission := permission.NewCheckPermissionResponse()
	roles := policySet.GetRoles()
	for _, item := range roles {
		if item.HasPermission(in.ServiceId, in.HttpMethod, in.HttpPath) {
			permission.HasPermission = true
			permission.Role = item
		}
	}
	return permission, nil
}
