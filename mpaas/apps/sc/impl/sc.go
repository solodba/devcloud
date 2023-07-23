package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/solodba/devcloud/mpaas/apps/sc"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建StorageClass
func (i *impl) CreateSC(ctx context.Context, in *sc.CreateSCRequest) (*sc.SC, error) {
	if !in.ValidatePlugin(i.conf) {
		return nil, fmt.Errorf("[%s] provisioner is not support", in.Provisioner)
	}
	k8sSC := i.SCReq2K8s(in)
	scApi := i.clientSet.StorageV1().StorageClasses()
	if _, err := scApi.Create(ctx, k8sSC, metav1.CreateOptions{}); err != nil {
		return nil, fmt.Errorf("[name=%s] StorageClass create fail", k8sSC.Name)
	}
	sc := sc.NewSC(in)
	// 写入到数据库
	_, err := i.col.InsertOne(ctx, sc)
	if err != nil {
		return nil, fmt.Errorf("[name=%s] StorageClass insert mongodb fail, err: %s", k8sSC.Name, err.Error())
	}
	return sc, nil
}

// 删除StorageClass
func (i *impl) DeleteSC(ctx context.Context, in *sc.DeleteSCRequest) (*sc.SC, error) {
	scApi := i.clientSet.StorageV1().StorageClasses()
	_, err := scApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[name=%s] StorageClass not found", in.Name)
	}
	err = scApi.Delete(ctx, in.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[name=%s] StorageClass delete fail", in.Name)
	}
	filter := bson.M{"name": in.Name}
	scReq := sc.NewDefaultSC()
	err = i.col.FindOne(ctx, filter).Decode(scReq)
	if err != nil {
		return nil, fmt.Errorf("[name=%s] is not found in mongodb", in.Name)
	}
	_, err = i.col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("[name=%s] delete from mongodb fail", in.Name)
	}
	return scReq, nil
}

// 查询StorageClass集合
func (i *impl) QuerySC(ctx context.Context, in *sc.QuerySCRequest) (*sc.SCSet, error) {
	scApi := i.clientSet.StorageV1().StorageClasses()
	k8sSCList, err := scApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get StorageClass list error, err: %s", err.Error())
	}
	scResList := sc.NewSCSet()
	for _, k8sSC := range k8sSCList.Items {
		scRes := i.SCK8s2Res(k8sSC)
		if strings.Contains(scRes.Name, in.Keyword) {
			scResList.AddItems(scRes)
		}
	}
	scResList.Total = int64(len(scResList.Items))
	return scResList, nil
}
