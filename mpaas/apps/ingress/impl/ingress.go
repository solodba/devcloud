package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/ingress"
)

// 创建Ingress
func (i *impl) CreateIngress(ctx context.Context, in *ingress.CreateIngressRequest) (*ingress.Ingress, error) {
	return nil, nil
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
