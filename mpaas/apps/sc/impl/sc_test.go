package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/sc"
	"github.com/solodba/devcloud/mpaas/test/tools"
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
