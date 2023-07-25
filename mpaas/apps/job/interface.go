package job

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/pod"
)

// 模块名称
const (
	AppName = "job"
)

// Job相关功能接口
type Service interface {
	// 嵌套Job GRPC接口
	RPCServer
}

// CreateJobRequest初始化函数
func NewCreateJobRequest() *CreateJobRequest {
	return &CreateJobRequest{
		Base: &JobBase{
			Labels: make([]*ListMapItem, 0),
		},
		Template: &pod.Pod{
			Volumes:        make([]*pod.Volume, 0),
			InitContainers: make([]*pod.Container, 0),
			Containers:     make([]*pod.Container, 0),
			Tolerations:    make([]*pod.Tolerations, 0),
		},
	}
}

// UpdateJobRequest初始化函数
func NewUpdateJobRequest() *UpdateJobRequest {
	return &UpdateJobRequest{
		Job: NewCreateJobRequest(),
	}
}

// QueryJobRequest初始化函数
func NewQueryJobRequest() *QueryJobRequest {
	return &QueryJobRequest{}
}

// DescribeJobRequest初始化函数
func NewDescribeJobRequest() *DescribeJobRequest {
	return &DescribeJobRequest{}
}

// DeleteJobRequest初始化函数
func NewDeleteJobRequest() *DeleteJobRequest {
	return &DeleteJobRequest{}
}

// 从restful解析参数初始化DeleteJobRequest
func NewDeleteJobRequestFromRestful(r *restful.Request) *DeleteJobRequest {
	return &DeleteJobRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从restful解析参数初始化QueryJobRequest
func NewQueryJobRequestFromRestful(r *restful.Request) *QueryJobRequest {
	return &QueryJobRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful解析参数初始化DescribeJobRequest
func NewDescribeJobRequestFromRestful(r *restful.Request) *DescribeJobRequest {
	return &DescribeJobRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}
