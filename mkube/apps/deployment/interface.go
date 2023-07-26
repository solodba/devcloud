package deployment

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/pod"
)

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

// DeleteDeploymentRequest初始化函数
func NewDeleteDeploymentRequest() *DeleteDeploymentRequest {
	return &DeleteDeploymentRequest{}
}

// 从restful解析参数初始化DeleteDeploymentRequest
func NewDeleteDeploymentRequestFromRestful(r *restful.Request) *DeleteDeploymentRequest {
	return &DeleteDeploymentRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从restful解析参数初始化QueryDeploymentRequest
func NewQueryDeploymentRequestFromRestful(r *restful.Request) *QueryDeploymentRequest {
	return &QueryDeploymentRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful解析参数初始化DescribeDeploymentRequest
func NewDescribeDeploymentRequestFromRestful(r *restful.Request) *DescribeDeploymentRequest {
	return &DescribeDeploymentRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}
