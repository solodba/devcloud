package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mkube/apps/namespace"
	"github.com/solodba/devcloud/mkube/test/tools"
)

func TestQueryNamespace(t *testing.T) {
	req := namespace.NewQueryNamespaceRequest()
	namespaceSet, err := svc.QueryNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(namespaceSet))
}
