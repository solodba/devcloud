package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/statefulset"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreateStatefulSet(t *testing.T) {
	req := statefulset.NewCreateStatefulSetRequest()
	dj, err := ioutil.ReadFile("statefulset.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	statefulset, err := svc.CreateStatefulSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(statefulset))
}

func TestUpdateStatefulSet(t *testing.T) {
	req := statefulset.NewUpdateStatefulSetRequest()
	dj, err := ioutil.ReadFile("statefulset.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.StatefulSet)
	if err != nil {
		t.Fatal(err)
	}
	statefulset, err := svc.UpdateStatefulSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(statefulset))
}

func TestQueryStatefulSet(t *testing.T) {
	req := statefulset.NewQueryStatefulSetRequest()
	req.Namespace = "test"
	req.Keyword = "nginx"
	statefulSetSet, err := svc.QueryStatefulSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(statefulSetSet))
}
