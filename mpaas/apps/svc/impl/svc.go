package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/svc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建Service
func (i *impl) CreateService(ctx context.Context, in *svc.CreateServiceRequest) (*svc.Service, error) {
	// Service转换
	k8sSVC := i.SVCReq2K8sConvert(in)
	svcApi := i.clientSet.CoreV1().Services(k8sSVC.Namespace)
	// 判断Service是否存在
	if getSVC, err := svcApi.Get(ctx, k8sSVC.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] service already exists", getSVC.Namespace, getSVC.Name)
	}
	// 创建Service
	_, err := svcApi.Create(ctx, k8sSVC, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] service create fail", k8sSVC.Namespace, k8sSVC.Name)
	}
	service := svc.NewService(in)
	// 入库
	_, err = i.col.InsertOne(ctx, service)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] service insert mongodb fail", k8sSVC.Namespace, k8sSVC.Name)
	}
	return service, nil
}

// 删除Service
func (i *impl) DeleteService(ctx context.Context, in *svc.DeleteServiceRequest) (*svc.Service, error) {
	return nil, nil
}

// 修改Service
func (i *impl) UpdateService(ctx context.Context, in *svc.UpdateServiceRequest) (*svc.Service, error) {
	return nil, nil
}

// 查询Service
func (i *impl) QueryService(ctx context.Context, in *svc.QueryServiceRequest) (*svc.ServiceSet, error) {
	return nil, nil
}

// 查询Service详情
func (i *impl) DescribeService(ctx context.Context, in *svc.DescribeServiceRequest) (*svc.Service, error) {
	return nil, nil
}
