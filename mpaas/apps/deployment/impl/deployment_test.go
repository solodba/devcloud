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

func TestUpdateDeployment(t *testing.T) {
	req := deployment.NewUpdateDeploymentRequest()
	dj, err := ioutil.ReadFile("deployment.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.Deployment)
	if err != nil {
		t.Fatal(err)
	}
	deployment, err := svc.UpdateDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(deployment))
}

func TestQueryDeployment(t *testing.T) {
	req := deployment.NewQueryDeploymentRequest()
	req.Namespace = "test"
	req.Keyword = "nginx"
	deploymentSet, err := svc.QueryDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(deploymentSet))
}

func TestDescribeDeployment(t *testing.T) {
	req := deployment.NewDescribeDeploymentRequest()
	req.Namespace = "test"
	req.Name = "nginx-deployment"
	deployment, err := svc.DescribeDeployment(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(deployment))
}
