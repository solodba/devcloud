package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mkube/apps/statefulset"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建StatefulSet
func (i *impl) CreateStatefulSet(ctx context.Context, in *statefulset.CreateStatefulSetRequest) (*statefulset.StatefulSet, error) {
	// StatefulSet转换
	k8sStatefulSet := i.StatefulSetReq2K8sConvert(in)
	statefulSetApi := i.clientSet.AppsV1().StatefulSets(k8sStatefulSet.Namespace)
	// 判断StatefulSet是否存在
	if getStatefulSet, err := statefulSetApi.Get(ctx, k8sStatefulSet.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] statefulset already exists", getStatefulSet.Namespace, getStatefulSet.Name)
	}
	// 创建StatefulSet
	_, err := statefulSetApi.Create(ctx, k8sStatefulSet, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] statefulset create fail", k8sStatefulSet.Namespace, k8sStatefulSet.Name)
	}
	statefulset := statefulset.NewStatefulSet(in)
	// 入库
	_, err = i.col.InsertOne(ctx, statefulset)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] statefulset insert mongodb fail, err: %s", k8sStatefulSet.Namespace, k8sStatefulSet.Name, err.Error())
	}
	return statefulset, nil
}

// 删除StatefulSet
func (i *impl) DeleteStatefulSet(ctx context.Context, in *statefulset.DeleteStatefulSetRequest) (*statefulset.CreateStatefulSetRequest, error) {
	req := statefulset.NewDescribeStatefulSetRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	statefulset, err := i.DescribeStatefulSet(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] statefulset not found", in.Namespace, in.Name)
	}
	statefulsetApi := i.clientSet.AppsV1().StatefulSets(statefulset.Base.Namespace)
	err = statefulsetApi.Delete(ctx, statefulset.Base.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] statefulset delete fail", statefulset.Base.Namespace, statefulset.Base.Name)
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"base.name": statefulset.Base.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", statefulset.Base.Namespace, statefulset.Base.Name)
	}
	return statefulset, nil
}

// 更新StatefulSet
func (i *impl) UpdateStatefulSet(ctx context.Context, in *statefulset.UpdateStatefulSetRequest) (*statefulset.StatefulSet, error) {
	k8sStatefulSet := i.StatefulSetReq2K8sConvert(in.StatefulSet)
	statefulsetApi := i.clientSet.AppsV1().StatefulSets(in.StatefulSet.Base.Namespace)
	if _, err := statefulsetApi.Get(ctx, in.StatefulSet.Base.Name, metav1.GetOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] statefulset not found", in.StatefulSet.Base.Namespace, in.StatefulSet.Base.Name)
	}
	_, err := statefulsetApi.Update(ctx, k8sStatefulSet, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] statefulset update error, err: %s", in.StatefulSet.Base.Namespace, in.StatefulSet.Base.Name, err.Error())
	}
	statefulset := statefulset.NewDefaultStatefulSet()
	err = i.col.FindOne(ctx, bson.M{"base.name": k8sStatefulSet.Name}).Decode(statefulset)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sStatefulSet.Namespace, k8sStatefulSet.Name, err.Error())
	}
	statefulset.Meta.UpdatedAt = time.Now().Unix()
	statefulset.StatefulSet = in.StatefulSet
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"base.name": k8sStatefulSet.Name}, bson.M{"$set": statefulset})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sStatefulSet.Namespace, k8sStatefulSet.Name, err.Error())
	}
	return statefulset, nil
}

// 查询StatefulSet
func (i *impl) QueryStatefulSet(ctx context.Context, in *statefulset.QueryStatefulSetRequest) (*statefulset.StatefulSetSet, error) {
	statefulsetApi := i.clientSet.AppsV1().StatefulSets(in.Namespace)
	k8sStatefulSetList, err := statefulsetApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get statefulset list error, err: %s", err.Error())
	}
	statefulSetSet := statefulset.NewStatefulSetSet()
	for _, k8sStatefulSet := range k8sStatefulSetList.Items {
		// 数据转换
		statefulSetRes := i.StatefulSetlK8s2ResConvert(k8sStatefulSet)
		if strings.Contains(statefulSetRes.Name, in.Keyword) {
			statefulSetSet.AddItems(statefulSetRes)
		}
	}
	statefulSetSet.Total = int64(len(statefulSetSet.Items))
	return statefulSetSet, nil
}

// 查询StatefulSet详情
func (i *impl) DescribeStatefulSet(ctx context.Context, in *statefulset.DescribeStatefulSetRequest) (*statefulset.CreateStatefulSetRequest, error) {
	statefulsetApi := i.clientSet.AppsV1().StatefulSets(in.Namespace)
	k8sStatefulSet, err := statefulsetApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get statefulset detail error, err: %s", err.Error())
	}
	return i.StatefulSetlK8s2ReqConvert(k8sStatefulSet), nil
}
