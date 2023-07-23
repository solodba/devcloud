package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/pvc"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建PersistentVolumeClaim
func (i *impl) CreatePVC(ctx context.Context, in *pvc.CreatePVCRequest) (*pvc.PVC, error) {
	k8sPVC := i.PVCReq2K8s(in)
	pvcApi := i.clientSet.CoreV1().PersistentVolumeClaims(k8sPVC.Namespace)
	_, err := pvcApi.Create(ctx, k8sPVC, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] PersistentVolumeClaim create fail", k8sPVC.Namespace, k8sPVC.Name)
	}
	pvc := pvc.NewPVC(in)
	_, err = i.col.InsertOne(ctx, pvc)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] PersistentVolumeClaim insert mongodb fail, err: %s", k8sPVC.Namespace, k8sPVC.Name, err.Error())
	}
	return pvc, nil
}

// 删除PersistentVolumeClaim
func (i *impl) DeletePVC(ctx context.Context, in *pvc.DeletePVCRequest) (*pvc.PVC, error) {
	pvcApi := i.clientSet.CoreV1().PersistentVolumeClaims(in.Namespace)
	_, err := pvcApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] PersistentVolumeClaim not found", in.Namespace, in.Name)
	}
	err = pvcApi.Delete(ctx, in.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] PersistentVolume delete fail", in.Namespace, in.Name)
	}
	filter := bson.M{"name": in.Name}
	pvcReq := pvc.NewDefaultPVC()
	err = i.col.FindOne(ctx, filter).Decode(pvcReq)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] is not found in mongodb", in.Namespace, in.Name)
	}
	_, err = i.col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] delete from mongodb fail", in.Namespace, in.Name)
	}
	return pvcReq, nil
}

// 查询PersistentVolumeClaim集合
func (i *impl) QueryPVC(ctx context.Context, in *pvc.QueryPVCRequest) (*pvc.PVCSet, error) {
	return nil, nil
}
