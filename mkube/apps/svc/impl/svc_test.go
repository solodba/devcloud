package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/svc"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestCreateService(t *testing.T) {
	req := svc.NewCreateServiceRequest()
	dj, err := ioutil.ReadFile("svc.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	service, err := svcService.CreateService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(service))
}

func TestUpdateService(t *testing.T) {
	req := svc.NewUpdateServiceRequest(svc.NewCreateServiceRequest())
	dj, err := ioutil.ReadFile("svc.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.Service)
	if err != nil {
		t.Fatal(err)
	}
	service, err := svcService.UpdateService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(service))
}

func TestQueryService(t *testing.T) {
	req := svc.NewQueryServiceRequest()
	req.Namespace = "test"
	req.Keyword = "haha"
	serviceSet, err := svcService.QueryService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(serviceSet))
}

func TestDescribeService(t *testing.T) {
	req := svc.NewDescribeServiceRequest()
	req.Namespace = "test"
	req.Name = "test-svc"
	service, err := svcService.DescribeService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(service))
}

func TestDeleteService(t *testing.T) {
	req := svc.NewDeleteServiceRequest()
	req.Namespace = "test"
	req.Name = "test-svc"
	service, err := svcService.DeleteService(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(service))
}
