package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/solodba/devcloud/mkube/apps/pv"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建PersistentVolume
func (i *impl) CreatePV(ctx context.Context, in *pv.CreatePVRequest) (*pv.PV, error) {
	k8sPV := i.PVReq2K8s(in)
	pvApi := i.clientSet.CoreV1().PersistentVolumes()
	if _, err := pvApi.Create(ctx, k8sPV, metav1.CreateOptions{}); err != nil {
		return nil, fmt.Errorf("[name=%s] PersistentVolume create fail", k8sPV.Name)
	}
	pv := pv.NewPV(in)
	// 写入到数据库
	_, err := i.col.InsertOne(ctx, pv)
	if err != nil {
		return nil, fmt.Errorf("[name=%s] PersistentVolume insert mongodb fail, err: %s", k8sPV.Name, err.Error())
	}
	return pv, nil
}

// 删除PersistentVolume
func (i *impl) DeletePV(ctx context.Context, in *pv.DeletePVRequest) (*pv.PV, error) {
	pvApi := i.clientSet.CoreV1().PersistentVolumes()
	_, err := pvApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[name=%s] PersistentVolume not found", in.Name)
	}
	err = pvApi.Delete(ctx, in.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[name=%s] PersistentVolume delete fail", in.Name)
	}
	filter := bson.M{"name": in.Name}
	pvReq := pv.NewDefaultPV()
	err = i.col.FindOne(ctx, filter).Decode(pvReq)
	if err != nil {
		return nil, fmt.Errorf("[name=%s] is not found in mongodb", in.Name)
	}
	_, err = i.col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("[name=%s] delete from mongodb fail", in.Name)
	}
	return pvReq, nil
}

// 查询PersistentVolume集合
func (i *impl) QueryPV(ctx context.Context, in *pv.QueryPVRequest) (*pv.PVSet, error) {
	pvApi := i.clientSet.CoreV1().PersistentVolumes()
	k8sPVList, err := pvApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get PersistentVolume list error, err: %s", err.Error())
	}
	pvResList := pv.NewPVSet()
	for _, k8sPV := range k8sPVList.Items {
		pvRes := i.PVK8s2Res(k8sPV)
		if strings.Contains(pvRes.Name, in.Keyword) {
			pvResList.AddItems(pvRes)
		}
	}
	pvResList.Total = int64(len(pvResList.Items))
	return pvResList, nil
}
