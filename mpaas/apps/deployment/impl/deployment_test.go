package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/deployment"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreateDeployment(t *testing.T) {
	req := deployment.NewCreateDeploymentRequest()
	dj, err := ioutil.ReadFile("deployment.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	deployment, err := svc.CreateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(deployment))
}
