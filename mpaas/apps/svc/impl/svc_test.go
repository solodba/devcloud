package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/svc"
	"github.com/solodba/devcloud/mpaas/test/tools"
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
