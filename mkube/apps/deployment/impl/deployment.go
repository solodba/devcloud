package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mkube/apps/deployment"
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
	req := deployment.NewDescribeDeploymentRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	deployment, err := i.DescribeDeployment(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] deployment not found", in.Namespace, in.Name)
	}
	deploymentApi := i.clientSet.AppsV1().Deployments(deployment.Base.Namespace)
	err = deploymentApi.Delete(ctx, deployment.Base.Name, metav1.DeleteOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] deployment delete fail", deployment.Base.Namespace, deployment.Base.Name)
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"base.name": deployment.Base.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", deployment.Base.Namespace, deployment.Base.Name)
	}
	return deployment, nil
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
	deploymentApi := i.clientSet.AppsV1().Deployments(in.Namespace)
	k8sDeploymentList, err := deploymentApi.List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get deoployment list error, err: %s", err.Error())
	}
	deploymentSet := deployment.NewDeploymentSet()
	for _, k8sDeployment := range k8sDeploymentList.Items {
		// 数据转换
		deploymentRes := i.DeploymentK8s2ResConvert(&k8sDeployment)
		if strings.Contains(deploymentRes.Name, in.Keyword) {
			deploymentSet.AddItems(deploymentRes)
		}
	}
	deploymentSet.Total = int64(len(deploymentSet.Items))
	return deploymentSet, nil
}

// 查询Deployment详情
func (i *impl) DescribeDeployment(ctx context.Context, in *deployment.DescribeDeploymentRequest) (*deployment.CreateDeploymentRequest, error) {
	deploymentApi := i.clientSet.AppsV1().Deployments(in.Namespace)
	k8sDeployment, err := deploymentApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get deployment detail error, err: %s", err.Error())
	}
	return i.DeploymentK8s2ReqConvert(k8sDeployment), nil
}
