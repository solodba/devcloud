package daemonset

import "github.com/solodba/devcloud/mpaas/apps/pod"

// 模块名称
const (
	AppName = "daemonset"
)

// DaemonSet功能接口
type Service interface {
	// 嵌套DaemonSet GRPC接口
	RPCServer
}

// CreateDaemonSetRequest初始化函数
func NewCreateDaemonSetRequest() *CreateDaemonSetRequest {
	return &CreateDaemonSetRequest{
		Base:     &Base{},
		Template: &pod.Pod{},
	}
}

// DeleteDaemonSetRequest初始化函数
func NewDeleteDaemonSetRequest() *DeleteDaemonSetRequest {
	return &DeleteDaemonSetRequest{}
}

// UpdateDaemonSetRequest初始化函数
func NewUpdateDaemonSetRequest() *UpdateDaemonSetRequest {
	return &UpdateDaemonSetRequest{
		DaemonSet: NewCreateDaemonSetRequest(),
	}
}

// QueryDaemonSetRequest初始化函数
func NewQueryDaemonSetRequest() *QueryDaemonSetRequest {
	return &QueryDaemonSetRequest{}
}

// DescribeDaemonSetRequest初始化函数
func NewDescribeDaemonSetRequest() *DescribeDaemonSetRequest {
	return &DescribeDaemonSetRequest{}
}
