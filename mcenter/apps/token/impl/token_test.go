package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/devcloud/tree/main/mcenter/test/tools"
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	req.Username = "test2"
	req.Password = "123456"
	tk, err := svc.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cinv113jrq66s110li70")
	tk, err := svc.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}
