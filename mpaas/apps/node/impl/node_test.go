package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/node"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestQueryNode(t *testing.T) {
	req := node.NewQueryNodeRequest()
	nodeSet, err := svc.QueryNode(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(nodeSet))
}
