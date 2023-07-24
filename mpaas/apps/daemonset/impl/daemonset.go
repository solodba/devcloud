package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/daemonset"
)

// 创建DaemonSet
func (i *impl) CreateDaemonSet(ctx context.Context, in *daemonset.CreateDaemonSetRequest) (*daemonset.DaemonSet, error) {
	return nil, nil
}

// 删除DaemonSet
func (i *impl) DeleteDaemonSet(ctx context.Context, in *daemonset.DeleteDaemonSetRequest) (*daemonset.CreateDaemonSetRequest, error) {
	return nil, nil
}

// 更新DaemonSet
func (i *impl) UpdateDaemonSet(ctx context.Context, in *daemonset.UpdateDaemonSetRequest) (*daemonset.DaemonSet, error) {
	return nil, nil
}

// 查询DaemonSet
func (i *impl) QueryDaemonSet(ctx context.Context, in *daemonset.QueryDaemonSetRequest) (*daemonset.DaemonSetList, error) {
	return nil, nil
}

// 查询DaemonSet详情
func (i *impl) DescribeDaemonSet(ctx context.Context, in *daemonset.DescribeDaemonSetRequest) (*daemonset.CreateDaemonSetRequest, error) {
	return nil, nil
}
