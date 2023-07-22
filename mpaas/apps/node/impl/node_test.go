package impl_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/solodba/devcloud/mpaas/apps/node"
	"github.com/solodba/devcloud/mpaas/test/tools"
)

func TestQueryNode(t *testing.T) {
	req := node.NewQueryNodeRequest("node")
	nodeSet, err := svc.QueryNode(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(nodeSet))
}

func TestDescribeNode(t *testing.T) {
	req := node.NewDescribeNodeRequest("node2")
	node, err := svc.DescribeNode(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(node))
}

func TestUpdateNodeLabel(t *testing.T) {
	req := node.NewUpdatedLabelRequest()
	dj, err := ioutil.ReadFile("label.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(dj, req)
	if err != nil {
		t.Fatal(err)
	}
	_, err = svc.UpdateNodeLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("更新" + req.Name + "节点标签成功")
}
