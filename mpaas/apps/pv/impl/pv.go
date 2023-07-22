package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/pv"
)

// 创建PersistentVolume
func (i *impl) CreatePV(ctx context.Context, in *pv.CreatePVRequest) (*pv.PV, error) {
	return nil, nil
}

// 删除PersistentVolume
func (i *impl) DeletePV(ctx context.Context, in *pv.DeletePVRequest) (*pv.PV, error) {
	return nil, nil
}

// 查询PersistentVolume集合
func (i *impl) QueryPV(ctx context.Context, in *pv.QueryPVRequest) (*pv.PVSet, error) {
	return nil, nil
}
