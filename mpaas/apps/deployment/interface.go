package deployment

import "github.com/solodba/devcloud/mpaas/apps/pod"

// 模块名称
const (
	AppName = "deployment"
)

// Deployment相关功能接口
type Service interface {
	// 嵌套Deployment GRPC接口
	RPCServer
}

// CreateDeploymentRequest初始化函数
func NewCreateDeploymentRequest() *CreateDeploymentRequest {
	return &CreateDeploymentRequest{
		Base:     &Base{},
		Template: &pod.Pod{},
	}
}

// UpdateDeploymentRequest初始化函数
func NewUpdateDeploymentRequest() *UpdateDeploymentRequest {
	return &UpdateDeploymentRequest{
		Deployment: &CreateDeploymentRequest{},
	}
}

// QueryDeploymentRequest初始化函数
func NewQueryDeploymentRequest() *QueryDeploymentRequest {
	return &QueryDeploymentRequest{}
}

// DescribeDeploymentRequest初始化函数
func NewDescribeDeploymentRequest() *DescribeDeploymentRequest {
	return &DescribeDeploymentRequest{}
}
