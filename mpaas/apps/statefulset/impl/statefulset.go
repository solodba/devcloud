package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/statefulset"
)

// 创建StatefulSet
func (i *impl) CreateStatefulSet(ctx context.Context, in *statefulset.CreateStatefulSetRequest) (*statefulset.StatefulSet, error) {
	return nil, nil
}

// 删除StatefulSet
func (i *impl) DeleteStatefulSet(ctx context.Context, in *statefulset.DeleteStatefulSetRequest) (*statefulset.CreateStatefulSetRequest, error) {
	return nil, nil
}

// 更新StatefulSet
func (i *impl) UpdateStatefulSet(ctx context.Context, in *statefulset.UpdateStatefulSetRequest) (*statefulset.StatefulSet, error) {
	return nil, nil
}

// 查询StatefulSet
func (i *impl) QueryStatefulSet(ctx context.Context, in *statefulset.QueryStatefulSetRequest) (*statefulset.StatefulSetSet, error) {
	return nil, nil
}

// 查询StatefulSet详情
func (i *impl) DescribeStatefulSet(ctx context.Context, in *statefulset.DescribeStatefulSetRequest) (*statefulset.CreateStatefulSetRequest, error) {
	return nil, nil
}
