package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/ingress"
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
	return nil, nil
}

// 删除Ingress
func (i *impl) DeleteIngress(ctx context.Context, in *ingress.DeleteIngressRequest) (*ingress.Ingress, error) {
	return nil, nil
}

// 查询Ingress
func (i *impl) QueryIngress(ctx context.Context, in *ingress.QueryIngressRequest) (*ingress.IngressSet, error) {
	return nil, nil
}

// 查询Ingress详情
func (i *impl) DescribeIngress(ctx context.Context, in *ingress.DescribeIngressRequest) (*ingress.Ingress, error) {
	return nil, nil
}
