package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/secret"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestCreateSecret(t *testing.T) {
	req := secret.NewCreateSecretRequest()
	dj, err := ioutil.ReadFile("secret.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	secret, err := svc.CreateSecret(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(secret))
}

func TestUpdateSecret(t *testing.T) {
	req := secret.NewUpdateSecretRequest(secret.NewCreateSecretRequest())
	dj, err := ioutil.ReadFile("secret.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req.Secret)
	if err != nil {
		t.Fatal(err)
	}
	secret, err := svc.UpdateSecret(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(secret))
}
