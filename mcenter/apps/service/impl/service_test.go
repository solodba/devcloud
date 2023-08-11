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

func TestDescribeService(t *testing.T) {
	req := service.NewDescribeServiceRequest()
	// req.DescribeType = service.DESCRIBE_BY_SERVICE_ID
	// req.DescribeValue = "cjauuq4fd1fkek1bmfpg"
	req.DescribeType = service.DESCRIBE_BY_SERVICE_CREDENTIAL_ID
	req.DescribeValue = "cjauuq4fd1fkek1bmfq0"
	serviceIns, err := svc.DescribeService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(serviceIns))
}
