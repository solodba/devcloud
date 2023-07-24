package daemonset

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mpaas/apps/pod"
)

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

// 从restful获取参数初始化DeleteDaemonSetRequest
func NewDeleteDaemonSetRequestFromRestful(r *restful.Request) *DeleteDaemonSetRequest {
	return &DeleteDaemonSetRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从restful获取参数初始化QueryDaemonSetRequest
func NewQueryDaemonSetRequestFromRestful(r *restful.Request) *QueryDaemonSetRequest {
	return &QueryDaemonSetRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful获取参数初始化DescribeDaemonSetRequest
func NewDescribeDaemonSetRequestFromRestful(r *restful.Request) *DescribeDaemonSetRequest {
	return &DescribeDaemonSetRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}
