package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/user"
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "test1"
	req.Password = "test1"
	userIns, err := svc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userIns)
}
