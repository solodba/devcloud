package impl

import (
	"context"

	"github.com/solodba/devcloud/mpaas/apps/deployment"
)

// 创建Deployment
func (i *impl) CreateDeployment(ctx context.Context, in *deployment.CreateDeploymentRequest) (*deployment.Deployment, error) {
	return nil, nil
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
