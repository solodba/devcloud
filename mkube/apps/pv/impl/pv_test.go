package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/pv"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestCreatePV(t *testing.T) {
	req := pv.NewCreatePVRequest()
	dj, err := ioutil.ReadFile("pv.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	pv, err := svc.CreatePV(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pv))
}

func TestDeletePV(t *testing.T) {
	req := pv.NewDeletePVRequest("testpv")
	pv, err := svc.DeletePV(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pv))
}

func TestQueryPV(t *testing.T) {
	req := pv.NewQueryPVRequest("test")
	pvSet, err := svc.QueryPV(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pvSet))
}
