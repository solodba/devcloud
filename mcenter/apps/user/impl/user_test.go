package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/user"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "test"
	req.Password = "test"
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

func TestDescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequest()
	req.DescribeType = user.DESCRIBE_BY_USER_ID
	req.DescribeValue = "cinr4c3jrq66iv4319rg"
	userIns, err := svc.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(userIns))

	// 用户密码校验
	err = userIns.CheckPassword("test2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUser(t *testing.T) {
	req := user.NewDeleteUserRequest()
	req.Username = "test1"
	userIns, err := svc.DeleteUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(userIns))
}

func TestUpdateUser(t *testing.T) {
	req := user.NewUpdateUserRequest()
	req.Username = "test2"
	req.Password = "123456"
	userIns, err := svc.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(userIns))

	// 密码校验
	err = userIns.CheckPassword("123456")
	if err != nil {
		t.Fatal(err)
	}
}
