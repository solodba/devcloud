package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/devcloud/tree/main/mcenter/test/tools"
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	tk, err := svc.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(tk))
}
