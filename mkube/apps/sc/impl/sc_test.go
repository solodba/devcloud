package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/sc"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestCreateSC(t *testing.T) {
	req := sc.NewCreateSCRequest()
	dj, err := ioutil.ReadFile("sc.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	sc, err := svc.CreateSC(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(sc))
}

func TestDeleteSC(t *testing.T) {
	req := sc.NewDeleteSCRequest("testsc")
	sc, err := svc.DeleteSC(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(sc))
}

func TestQuerySC(t *testing.T) {
	req := sc.NewQuerySCRequest("tes")
	scSet, err := svc.QuerySC(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(scSet))
}
