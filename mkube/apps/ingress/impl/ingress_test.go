package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/ingress"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestCreateIngress(t *testing.T) {
	req := ingress.NewCreateIngressRequest()
	dj, err := ioutil.ReadFile("ingress.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	ingress, err := svc.CreateIngress(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingress))
}

func TestUpdateIngress(t *testing.T) {
	req := ingress.NewUpdateIngressRequest(ingress.NewCreateIngressRequest())
	dj, err := ioutil.ReadFile("ingress.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.Ingress)
	if err != nil {
		t.Fatal(err)
	}
	ingress, err := svc.UpdateIngress(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingress))
}

func TestQueryIngress(t *testing.T) {
	req := ingress.NewQueryIngressRequest()
	req.Namespace = "test"
	req.Keyword = "te"
	ingressSet, err := svc.QueryIngress(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingressSet))
}

func TestDescribeIngress(t *testing.T) {
	req := ingress.NewDescribeIngressRequest()
	req.Namespace = "test"
	req.Name = "test-ingress"
	ingress, err := svc.DescribeIngress(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingress))
}

func TestDeleteIngress(t *testing.T) {
	req := ingress.NewDeleteIngressRequest()
	req.Namespace = "test"
	req.Name = "test-ingress"
	ingress, err := svc.DeleteIngress(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ingress))
}
