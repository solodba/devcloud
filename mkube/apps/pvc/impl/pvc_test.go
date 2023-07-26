package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mkube/apps/pvc"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestCreatePVC(t *testing.T) {
	req := pvc.NewCreatePVCRequest()
	dj, err := ioutil.ReadFile("pvc_sc.json")
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

func TestDeletePVC(t *testing.T) {
	req := pvc.NewDeletePVCRequest()
	req.Namespace = "test"
	req.Name = "testpvc"
	pvc, err := svc.DeletePVC(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pvc))
}

func TestQueryPVC(t *testing.T) {
	req := pvc.NewQueryPVCRequest()
	req.Namespace = "test"
	req.Keyword = "tes"
	pvcSet, err := svc.QueryPVC(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(pvcSet))
}
