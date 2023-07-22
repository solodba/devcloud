package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/pv"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建PersistentVolume
func (i *impl) CreatePV(ctx context.Context, in *pv.CreatePVRequest) (*pv.PV, error) {
	k8sPV := i.PVReq2K8s(in)
	pvApi := i.clientSet.CoreV1().PersistentVolumes()
	if _, err := pvApi.Create(ctx, k8sPV, metav1.CreateOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] PersistentVolume create fail", k8sPV.Namespace, k8sPV.Name)
	}
	pv := pv.NewPV(in)
	return pv, nil
}

// 删除PersistentVolume
func (i *impl) DeletePV(ctx context.Context, in *pv.DeletePVRequest) (*pv.PV, error) {
	return nil, nil
}

// 查询PersistentVolume集合
func (i *impl) QueryPV(ctx context.Context, in *pv.QueryPVRequest) (*pv.PVSet, error) {
	return nil, nil
}
