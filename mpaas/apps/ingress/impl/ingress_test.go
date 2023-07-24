package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/ingress"
	"github.com/solodba/devcloud/mpaas/test/tools"
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