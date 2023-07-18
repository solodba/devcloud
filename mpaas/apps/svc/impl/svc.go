package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mpaas/apps/svc"
	"go.mongodb.org/mongo-driver/bson"
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
		return nil, fmt.Errorf("[namespace=%s, name=%s] service insert mongodb fail, err: %s", k8sSVC.Namespace, k8sSVC.Name, err.Error())
	}
	return service, nil
}

// 删除Service
func (i *impl) DeleteService(ctx context.Context, in *svc.DeleteServiceRequest) (*svc.Service, error) {
	return nil, nil
}

// 修改Service
func (i *impl) UpdateService(ctx context.Context, in *svc.UpdateServiceRequest) (*svc.Service, error) {
	k8sSVC := i.SVCReq2K8sConvert(in.Service)
	svcApi := i.clientSet.CoreV1().Services(in.Service.Namespace)
	if _, err := svcApi.Get(ctx, in.Service.Name, metav1.GetOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] service not exists", in.Service.Namespace, in.Service.Name)
	}
	_, err := svcApi.Update(ctx, k8sSVC, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] service update error, err: %s", in.Service.Namespace, in.Service.Name, err.Error())
	}
	service := svc.NewDefaultService()
	err = i.col.FindOne(ctx, bson.M{"name": k8sSVC.Name}).Decode(service)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sSVC.Namespace, k8sSVC.Name, err.Error())
	}
	service.Meta.UpdatedAt = time.Now().Unix()
	service.Service = in.Service
	_, err = i.col.UpdateOne(ctx, bson.M{"name": k8sSVC.Name}, bson.M{"$set": service})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sSVC.Namespace, k8sSVC.Name, err.Error())
	}
	return service, nil
}

// 查询Service
func (i *impl) QueryService(ctx context.Context, in *svc.QueryServiceRequest) (*svc.ServiceSet, error) {
	svcApi := i.clientSet.CoreV1().Services(in.Namespace)
	k8sSVCList, err := svcApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get service list error, err: %s", err.Error())
	}
	serviceSet := svc.NewServiceSet()
	for _, k8sSvc := range k8sSVCList.Items {
		// 数据转换
		svcRes := i.SVCK8s2ResConvert(k8sSvc)
		if strings.Contains(svcRes.Name, in.Keyword) {
			serviceSet.AddItems(svcRes)
		}
	}
	serviceSet.Total = int64(len(serviceSet.Items))
	return serviceSet, nil
}

// 查询Service详情
func (i *impl) DescribeService(ctx context.Context, in *svc.DescribeServiceRequest) (*svc.Service, error) {
	svcApi := i.clientSet.CoreV1().Services(in.Namespace)
	k8sSVC, err := svcApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get service detail error, err: %s", err.Error())
	}
	service := i.SVCK8s2Req(k8sSVC)
	return service, nil
}
