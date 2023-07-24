package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/solodba/devcloud/mpaas/apps/deployment"
	"go.mongodb.org/mongo-driver/bson"
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
	k8sDeployment := i.DeploymentReq2K8sConvert(in.Deployment)
	deploymentApi := i.clientSet.AppsV1().Deployments(in.Deployment.Base.Namespace)
	if _, err := deploymentApi.Get(ctx, in.Deployment.Base.Name, metav1.GetOptions{}); err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] deployment not found", in.Deployment.Base.Namespace, in.Deployment.Base.Name)
	}
	_, err := deploymentApi.Update(ctx, k8sDeployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] deployment update error, err: %s", in.Deployment.Base.Namespace, in.Deployment.Base.Name, err.Error())
	}
	deployment := deployment.NewDefaultDeployment()
	err = i.col.FindOne(ctx, bson.M{"base.name": k8sDeployment.Name}).Decode(deployment)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", k8sDeployment.Namespace, k8sDeployment.Name, err.Error())
	}
	deployment.Meta.UpdatedAt = time.Now().Unix()
	deployment.Deployment = in.Deployment
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"base.name": k8sDeployment.Name}, bson.M{"$set": deployment})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", k8sDeployment.Namespace, k8sDeployment.Name, err.Error())
	}
	return deployment, nil
}

// 查询Deployment
func (i *impl) QueryDeployment(ctx context.Context, in *deployment.QueryDeploymentRequest) (*deployment.DeploymentSet, error) {
	return nil, nil
}

// 查询Deployment详情
func (i *impl) DescribeDeployment(ctx context.Context, in *deployment.DescribeDeploymentRequest) (*deployment.CreateDeploymentRequest, error) {
	return nil, nil
}
