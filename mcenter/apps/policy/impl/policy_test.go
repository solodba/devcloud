package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/namespace"
	"github.com/solodba/devcloud/mcenter/apps/policy"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.RoleId = "cjd2b08eaqnhr91ir71g"
	req.UserId = "cjd3dm0eaqnl1k6f5ctg"
	policyIns, err := svc.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(policyIns))
}
