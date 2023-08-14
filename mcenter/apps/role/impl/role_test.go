package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/role"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestCreateRole(t *testing.T) {
	f1 := role.NewFeature()
	f1.ServiceId = "cjd1nu0eaqnlb444u060"
	f1.HttpMethod = "POST"
	f1.HttpPath = "/mpaas/api/v1/cluster"
	f2 := role.NewFeature()
	f2.ServiceId = "cjd1nu0eaqnlb444u060"
	f2.HttpMethod = "GET"
	f2.HttpPath = "/mpaas/api/v1/cluster"
	req := role.NewCreateRoleRequest()
	req.Name = "test"
	req.AddItems(f1, f2)
	roleIns, err := svc.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(roleIns))
}

func TestQueryRole(t *testing.T) {
	req := role.NewQueryRoleRequest()
	// req.Page.PageNumber = 1
	// req.Page.PageSize = 1
	req.AddItems("cjd2b08eaqnhr91ir71g", "cjd2qs0eaqnjhu20sgqg")
	roleSetIns, err := svc.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(roleSetIns))
}
