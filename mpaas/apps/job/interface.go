package job

import "github.com/solodba/devcloud/mpaas/apps/pod"

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
