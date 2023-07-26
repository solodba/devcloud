package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mkube/apps/daemonset"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建DaemonSet
func (i *impl) CreateDaemonSet(ctx context.Context, in *daemonset.CreateDaemonSetRequest) (*daemonset.DaemonSet, error) {
	// DaemonSet转换
	k8sDaemonSet := i.DaemonSetReq2K8sConvert(in)
	daemonSetApi := i.clientSet.AppsV1().DaemonSets(k8sDaemonSet.Namespace)
	// 判断DaemonSet是否存在
	if getK8sDaemonSet, err := daemonSetApi.Get(ctx, k8sDaemonSet.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] daemonset already exists", getK8sDaemonSet.Namespace, getK8sDaemonSet.Name)
	}
	// 创建DaemonSet
	_, err := daemonSetApi.Create(ctx, k8sDaemonSet, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] daemonset create fail", k8sDaemonSet.Namespace, k8sDaemonSet.Name)
	}
	daemonset := daemonset.NewDaemonSet(in)
	// 入库
	_, err = i.col.InsertOne(ctx, daemonset)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] daemonset insert mongodb fail, err: %s", k8sDaemonSet.Namespace, k8sDaemonSet.Name, err.Error())
	}
	return daemonset, nil
}

// 删除DaemonSet
func (i *impl) DeleteDaemonSet(ctx context.Context, in *daemonset.DeleteDaemonSetRequest) (*daemonset.CreateDaemonSetRequest, error) {
	req := daemonset.NewDescribeDaemonSetRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	daemonset, err := i.DescribeDaemonSet(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] daemonset not found", in.Namespace, in.Name)
	}
	daemonSetApi := i.clientSet.AppsV1().DaemonSets(daemonset.Base.Namespace)
	err = daemonSetApi.Delete(ctx, daemonset.Base.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] daemonset delete fail", daemonset.Base.Namespace, daemonset.Base.Name)
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"base.name": daemonset.Base.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", daemonset.Base.Namespace, daemonset.Base.Name)
	}
	return daemonset, nil
}

// 更新DaemonSet
func (i *impl) UpdateDaemonSet(ctx context.Context, in *daemonset.UpdateDaemonSetRequest) (*daemonset.DaemonSet, error) {
	k8sDaemonSet := i.DaemonSetReq2K8sConvert(in.DaemonSet)
	daemonSetApi := i.clientSet.AppsV1().DaemonSets(in.DaemonSet.Base.Namespace)
	if _, err := daemonSetApi.Get(ctx, in.DaemonSet.Base.Name, metav1.GetOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] daemonset not found", in.DaemonSet.Base.Namespace, in.DaemonSet.Base.Name)
	}
	_, err := daemonSetApi.Update(ctx, k8sDaemonSet, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] daemonset update error, err: %s", in.DaemonSet.Base.Namespace, in.DaemonSet.Base.Name, err.Error())
	}
	daemonset := daemonset.NewDefaultDaemonSet()
	err = i.col.FindOne(ctx, bson.M{"base.name": k8sDaemonSet.Name}).Decode(daemonset)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sDaemonSet.Namespace, k8sDaemonSet.Name, err.Error())
	}
	daemonset.Meta.UpdatedAt = time.Now().Unix()
	daemonset.DaemonSet = in.DaemonSet
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"base.name": k8sDaemonSet.Name}, bson.M{"$set": daemonset})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sDaemonSet.Namespace, k8sDaemonSet.Name, err.Error())
	}
	return daemonset, nil
}

// 查询DaemonSet
func (i *impl) QueryDaemonSet(ctx context.Context, in *daemonset.QueryDaemonSetRequest) (*daemonset.DaemonSetList, error) {
	daemonsetApi := i.clientSet.AppsV1().DaemonSets(in.Namespace)
	k8sDaemonSetList, err := daemonsetApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get daemonset list error, err: %s", err.Error())
	}
	daemonsetList := daemonset.NewDaemonSetList()
	for _, k8sDamonSet := range k8sDaemonSetList.Items {
		// 数据转换
		daemonsetRes := i.DaemonSetK8s2ResConvert(k8sDamonSet)
		if strings.Contains(daemonsetRes.Name, in.Keyword) {
			daemonsetList.AddItems(daemonsetRes)
		}
	}
	daemonsetList.Total = int64(len(daemonsetList.Items))
	return daemonsetList, nil
}

// 查询DaemonSet详情
func (i *impl) DescribeDaemonSet(ctx context.Context, in *daemonset.DescribeDaemonSetRequest) (*daemonset.CreateDaemonSetRequest, error) {
	daemonSetApi := i.clientSet.AppsV1().DaemonSets(in.Namespace)
	k8sDaemonSet, err := daemonSetApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get daemonset detail error, err: %s", err.Error())
	}
	return i.DaemonSetK8s2ReqConvert(k8sDaemonSet), nil
}
