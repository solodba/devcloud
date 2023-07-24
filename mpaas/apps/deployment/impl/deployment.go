package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/deployment"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建Deployment
func (i *impl) CreateDeployment(ctx context.Context, in *deployment.CreateDeploymentRequest) (*deployment.Deployment, error) {
	// Deployment转换
	k8sDeployment := i.DeploymentReq2K8sConvert(in)
	deploymentApi := i.clientSet.AppsV1().Deployments(k8sDeployment.Namespace)
	// 判断Deployment是否存在
	if getK8sDeployment, err := deploymentApi.Get(ctx, k8sDeployment.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] deployment already exists", getK8sDeployment.Namespace, getK8sDeployment.Name)
	}
	// 创建Deployment
	_, err := deploymentApi.Create(ctx, k8sDeployment, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] deployment create fail", k8sDeployment.Namespace, k8sDeployment.Name)
	}
	deployment := deployment.NewDeployment(in)
	// 入库
	_, err = i.col.InsertOne(ctx, deployment)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] deployment insert mongodb fail, err: %s", k8sDeployment.Namespace, k8sDeployment.Name, err.Error())
	}
	return deployment, nil
}

// 删除Deployment
func (i *impl) DeleteDeployment(ctx context.Context, in *deployment.DeleteDeploymentRequest) (*deployment.CreateDeploymentRequest, error) {
	return nil, nil
}

// 更新Deployment
func (i *impl) UpdateDeployment(ctx context.Context, in *deployment.UpdateDeploymentRequest) (*deployment.Deployment, error) {
	return nil, nil
}

// 查询Deployment
func (i *impl) QueryDeployment(ctx context.Context, in *deployment.QueryDeploymentRequest) (*deployment.DeploymentSet, error) {
	return nil, nil
}

// 查询Deployment详情
func (i *impl) DescribeDeployment(ctx context.Context, in *deployment.DescribeDeploymentRequest) (*deployment.CreateDeploymentRequest, error) {
	return nil, nil
}
