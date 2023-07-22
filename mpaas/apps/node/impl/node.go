package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/node"
)

// 查询Node集合
func (i *impl) QueryNode(ctx context.Context, in *node.QueryNodeRequest) (*node.NodeSet, error) {
	return nil, nil
}

// 查询Node详情
func (i *impl) DescribeNode(ctx context.Context, in *node.DescribeNodeRequest) (*node.NodeSetItem, error) {
	return nil, nil
}

// 更新NodeLabel
func (i *impl) UpdateNodeLabel(ctx context.Context, in *node.UpdatedLabelRequest) (*node.UpdatedLabelResponse, error) {
	return nil, nil
}

// 更新NodeTaint
func (i *impl) UpdateNodeTaint(ctx context.Context, in *node.UpdatedTaintRequest) (*node.UpdatedTaintResponse, error) {
	return nil, nil
}
