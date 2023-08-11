package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/service"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestCreateService(t *testing.T) {
	req := service.NewCreateServiceRequest()
	req.Domain = "default"
	req.Namespace = "default"
	req.Name = "test"
	serviceIns, err := svc.CreateService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(serviceIns))
}
