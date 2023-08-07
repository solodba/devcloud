package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/ingroute"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestDescribeIngressRoute(t *testing.T) {
	req := ingroute.NewDescribeIngressRouteRequest()
	req.Namespace = "test"
	req.Name = "web-ingressroute"
	ingressRoute, err := svc.DescribeIngressRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingressRoute))
}

func TestQueryIngressRoute(t *testing.T) {
	req := ingroute.NewQueryIngressRouteRequest()
	req.Namespace = "test"
	req.Keyword = "web"
	ingressRouteSet, err := svc.QueryIngressRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingressRouteSet))
}

func TestCreateIngressRoute(t *testing.T) {
	req := ingroute.NewCreateIngressRouteRequest()
	dj, err := ioutil.ReadFile("ingroute.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	ingressRoute, err := svc.CreateIngressRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingressRoute))
}

func TestUpdateIngressRoute(t *testing.T) {
	req := ingroute.NewUpdateIngressRouteRequest()
	dj, err := ioutil.ReadFile("ingroute.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.IngressRoute)
	if err != nil {
		t.Fatal(err)
	}
	ingressRoute, err := svc.UpdateIngressRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingressRoute))
}

func TestDeleteIngressRoute(t *testing.T) {
	req := ingroute.NewDeleteIngressRouteRequest()
	req.Namespace = "test"
	req.Name = "test"
	ingressRoute, err := svc.DeleteIngressRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingressRoute))
}
