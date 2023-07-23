package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/pvc"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreatePVC(t *testing.T) {
	req := pvc.NewCreatePVCRequest()
	dj, err := ioutil.ReadFile("pvc.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	pvc, err := svc.CreatePVC(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pvc))
}
