package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/daemonset"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestCreateDaemonSet(t *testing.T) {
	req := daemonset.NewCreateDaemonSetRequest()
	dj, err := ioutil.ReadFile("daemonset.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	daemonset, err := svc.CreateDaemonSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(daemonset))
}

func TestUpdateDaemonSet(t *testing.T) {
	req := daemonset.NewUpdateDaemonSetRequest()
	dj, err := ioutil.ReadFile("daemonset.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.DaemonSet)
	if err != nil {
		t.Fatal(err)
	}
	daemonset, err := svc.UpdateDaemonSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(daemonset))
}

func TestQueryDaemonSet(t *testing.T) {
	req := daemonset.NewQueryDaemonSetRequest()
	req.Namespace = "test"
	req.Keyword = "nginx"
	daemonsetList, err := svc.QueryDaemonSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(daemonsetList))
}

func TestDescribeDaemonSet(t *testing.T) {
	req := daemonset.NewDescribeDaemonSetRequest()
	req.Namespace = "test"
	req.Name = "nginx-daemonset"
	daemonset, err := svc.DescribeDaemonSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(daemonset))
}

func TestDeleteDaemonSet(t *testing.T) {
	req := daemonset.NewDeleteDaemonSetRequest()
	req.Namespace = "test"
	req.Name = "nginx-daemonset"
	daemonset, err := svc.DeleteDaemonSet(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(daemonset))
}
