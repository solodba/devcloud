package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/user"
	"github.com/solodba/devcloud/tree/main/mcenter/test/tools"
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "test6"
	req.Password = "test6"
	userIns, err := svc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(userIns))
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	userSet, err := svc.Queryuser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(userSet))
}
