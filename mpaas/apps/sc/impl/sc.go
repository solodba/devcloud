package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/sc"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建StorageClass
func (i *impl) CreateSC(ctx context.Context, in *sc.CreateSCRequest) (*sc.SC, error) {
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
	return nil, nil
}