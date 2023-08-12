package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mcenter/apps/endpoint"
	"github.com/solodba/devcloud/mcenter/test/tools"
)

func TestRegistryEndpoint(t *testing.T) {
	registryReq := endpoint.NewRegistryEndpointRequest()
	userReq := endpoint.NewCreateEndpointRequest()
	userReq.ServiceId = "cjauuq4fd1fkek1bmfpg"
	userReq.Auth = true
	userReq.Method = "POST"
	userReq.Operation = "Test"
	userReq.Path = "api/v1/token"
	userReq.Perm = true
	tokenReq := endpoint.NewCreateEndpointRequest()
	tokenReq.ServiceId = "aaaaaaaaaaaaaaaaaaaa"
	tokenReq.Auth = true
	tokenReq.Method = "GET"
	tokenReq.Operation = "GetToken"
	tokenReq.Path = "api/v1/token"
	tokenReq.Perm = true
	registryReq.AddItems(userReq, tokenReq)
	es, err := svc.RegistryEndpoint(ctx, registryReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(es))
}
