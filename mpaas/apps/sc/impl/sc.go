package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/sc"
)

// 创建StorageClass
func (i *impl) CreateSC(ctx context.Context, in *sc.CreateSCRequest) (*sc.SC, error) {
	return nil, nil
}

// 删除StorageClass
func (i *impl) DeleteSC(ctx context.Context, in *sc.DeleteSCRequest) (*sc.SC, error) {
	return nil, nil
}

// 查询StorageClass集合
func (i *impl) QuerySC(ctx context.Context, in *sc.QuerySCRequest) (*sc.SCSet, error) {
	return nil, nil
}
