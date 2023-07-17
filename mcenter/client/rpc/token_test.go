package rpc_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestIssueToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cio16drjrq647o5boo1g")
	tk, err := c.NewTokenRPCClient().ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}
