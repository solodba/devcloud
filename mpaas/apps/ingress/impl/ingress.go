package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mpaas/apps/ingress"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建Ingress
func (i *impl) CreateIngress(ctx context.Context, in *ingress.CreateIngressRequest) (*ingress.Ingress, error) {
	k8sIngress := i.IngressReq2K8sConvert(in)
	ingressApi := i.clientSet.NetworkingV1().Ingresses(k8sIngress.Namespace)
	if getPod, err := ingressApi.Get(ctx, k8sIngress.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] ingress already exists", getPod.Namespace, getPod.Name)
	}
	_, err := ingressApi.Create(ctx, k8sIngress, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] create failed, err: %s", k8sIngress.Namespace, k8sIngress.Name, err.Error())
	}
	// 新增Pod数据入库]
	ingress := ingress.NewIngress(in)
	_, err = i.col.InsertOne(ctx, ingress)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] insert into mongodb failed, err: %s", k8sIngress.Namespace, k8sIngress.Name, err.Error())
	}
	return ingress, nil
}

// 更新Ingress
func (i *impl) UpdateIngress(ctx context.Context, in *ingress.UpdateIngressRequest) (*ingress.Ingress, error) {
	k8sIngress := i.IngressReq2K8sConvert(in.Ingress)
	ingressApi := i.clientSet.NetworkingV1().Ingresses(k8sIngress.Namespace)
	if _, err := ingressApi.Get(ctx, in.Ingress.Name, metav1.GetOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] ingress not found", in.Ingress.Namespace, in.Ingress.Name)
	}
	_, err := ingressApi.Update(ctx, k8sIngress, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] ingress update error, err: %s", in.Ingress.Namespace, in.Ingress.Name, err.Error())
	}
	ingress := ingress.NewDefaultIngress()
	err = i.col.FindOne(ctx, bson.M{"name": k8sIngress.Name}).Decode(ingress)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sIngress.Namespace, k8sIngress.Name, err.Error())
	}
	ingress.Meta.UpdatedAt = time.Now().Unix()
	ingress.Ingress = in.Ingress
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"name": k8sIngress.Name}, bson.M{"$set": ingress})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sIngress.Namespace, k8sIngress.Name, err.Error())
	}
	return ingress, nil
}

// 删除Ingress
func (i *impl) DeleteIngress(ctx context.Context, in *ingress.DeleteIngressRequest) (*ingress.Ingress, error) {
	return nil, nil
}

// 查询Ingress
func (i *impl) QueryIngress(ctx context.Context, in *ingress.QueryIngressRequest) (*ingress.IngressSet, error) {
	ingressApi := i.clientSet.NetworkingV1().Ingresses(in.Namespace)
	k8sIngressList, err := ingressApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get ingress list error, err: %s", err.Error())
	}
	ingressSet := ingress.NewIngressSet()
	for _, k8sIngress := range k8sIngressList.Items {
		// 数据转换
		ingressRes := i.IngressK8s2ResConvert(k8sIngress)
		if strings.Contains(ingressRes.Name, in.Keyword) {
			ingressSet.AddItems(ingressRes)
		}
	}
	ingressSet.Total = int64(len(ingressSet.Items))
	return ingressSet, nil
}

// 查询Ingress详情
func (i *impl) DescribeIngress(ctx context.Context, in *ingress.DescribeIngressRequest) (*ingress.Ingress, error) {
	return nil, nil
}
