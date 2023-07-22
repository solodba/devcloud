package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/solodba/devcloud/mpaas/apps/node"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// 查询Node集合
func (i *impl) QueryNode(ctx context.Context, in *node.QueryNodeRequest) (*node.NodeSet, error) {
	k8sNodeList, err := i.clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get node list error, err: %s", err.Error())
	}
	nodeResList := node.NewNodeSet()
	for _, item := range k8sNodeList.Items {
		if strings.Contains(item.Name, in.Keyword) {
			// 数据结构体转换
			nodeRes := i.GetNodeResItem(item)
			nodeResList.AddItems(nodeRes)
		}
	}
	nodeResList.Total = int64(len(nodeResList.Items))
	return nodeResList, nil
}

// 查询Node详情
func (i *impl) DescribeNode(ctx context.Context, in *node.DescribeNodeRequest) (*node.NodeSetItem, error) {
	k8sNode, err := i.clientSet.CoreV1().Nodes().Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get node detail error, err: %s", err.Error())
	}
	return i.GetNodeDetail(k8sNode), nil
}

// 更新NodeLabel
func (i *impl) UpdateNodeLabel(ctx context.Context, in *node.UpdatedLabelRequest) (*node.UpdatedLabelResponse, error) {
	labelMap := make(map[string]string, 0)
	for _, label := range in.Labels {
		labelMap[label.Key] = label.Value
	}
	labelMap["$patch"] = "replace"
	patchData := map[string]any{
		"metadata": map[string]any{
			"labels": labelMap,
		},
	}
	patchDataBytes, _ := json.Marshal(&patchData)
	_, err := i.clientSet.CoreV1().Nodes().Patch(
		context.TODO(),
		in.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return nil, err
}

// 更新NodeTaint
func (i *impl) UpdateNodeTaint(ctx context.Context, in *node.UpdatedTaintRequest) (*node.UpdatedTaintResponse, error) {
	return nil, nil
}
