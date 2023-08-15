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
	req.RoleId = "cjdiiboeaqnjfb39enj0"
	req.UserId = "cjdfgs0eaqnmjv6mnb7g"
	policyIns, err := svc.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(policyIns))
}

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.UserId = "cjdfgs0eaqnmjv6mnb7g"
	req.WithRole = true
	policySet, err := svc.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(policySet))
}
