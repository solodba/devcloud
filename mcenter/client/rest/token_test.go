package rest_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	req.Username = "test2"
	req.Password = "123456"
	tk, err := c.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}
