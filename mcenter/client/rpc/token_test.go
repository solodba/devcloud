package rpc_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cjavod4fd1flcq19qgk0")
	tk, err := c.NewTokenRPCClient().ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}
