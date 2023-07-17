package rpc_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/user"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	userSet, err := c.NewUserRPCClient().Queryuser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(userSet))
}

func TestDescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequest()
	// req.DescribeValue = "test2"
	req.DescribeType = user.DESCRIBE_BY_USER_ID
	req.DescribeValue = "cinr4c3jrq66iv4319rg"
	userIns, err := c.NewUserRPCClient().DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(userIns))
}
