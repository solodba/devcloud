package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	req.Username = "admin"
	req.Password = "admin"
	tk, err := svc.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cjdfh9oeaqnibn3428qg")
	tk, err := svc.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}
