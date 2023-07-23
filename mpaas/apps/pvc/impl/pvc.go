package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/pvc"
)

// 创建PersistentVolumeClaim
func (i *impl) CreatePVC(ctx context.Context, in *pvc.CreatePVCRequest) (*pvc.PVC, error) {
	return nil, nil
}

// 删除PersistentVolumeClaim
func (i *impl) DeletePVC(ctx context.Context, in *pvc.DeletePVCRequest) (*pvc.PVC, error) {
	return nil, nil
}

// 查询PersistentVolumeClaim集合
func (i *impl) QueryPVC(ctx context.Context, in *pvc.QueryPVCRequest) (*pvc.PVCSet, error) {
	return nil, nil
}
