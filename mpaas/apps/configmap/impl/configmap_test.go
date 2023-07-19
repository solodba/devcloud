package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/configmap"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreateConfigMap(t *testing.T) {
	req := configmap.NewCreateConfigMapRequest()
	dj, err := ioutil.ReadFile("configmap.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	configmap, err := svc.CreateConfigMap(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(configmap))
}

func TestUpdateConfigMap(t *testing.T) {
	req := configmap.NewUpdateConfigMapRequest(configmap.NewCreateConfigMapRequest())
	dj, err := ioutil.ReadFile("configmap.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.ConfigMap)
	if err != nil {
		t.Fatal(err)
	}
	configmap, err := svc.UpdateConfigMap(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(configmap))
}

func TestQueryConfigMap(t *testing.T) {
	req := configmap.NewQueryConfigMapRequest()
	req.Namespace = "test"
	req.Keyword = "te"
	configmapSet, err := svc.QueryConfigMap(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(configmapSet))
}
