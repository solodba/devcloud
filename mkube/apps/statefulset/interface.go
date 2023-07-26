package statefulset

import (
	"github.com/emicklei/go-restful/v3"
	pod "github.com/solodba/devcloud/mkube/apps/pod"
)

// 模块名
const (
	AppName = "statefulset"
)

// StatefulSet功能接口
type Service interface {
	// 嵌套StatefulSet GRPC接口
	RPCServer
}

// CreateStatefulSetRequest初始化函数
func NewCreateStatefulSetRequest() *CreateStatefulSetRequest {
	return &CreateStatefulSetRequest{
		Base: &Base{
			Labels:   make([]*ListMapItem, 0),
			Selector: make([]*ListMapItem, 0),
		},
		Template: &pod.Pod{
			Volumes:        make([]*pod.Volume, 0),
			InitContainers: make([]*pod.Container, 0),
			Containers:     make([]*pod.Container, 0),
			Tolerations:    make([]*pod.Tolerations, 0),
		},
	}
}

// DeleteStatefulSetRequest初始化函数
func NewDeleteStatefulSetRequest() *DeleteStatefulSetRequest {
	return &DeleteStatefulSetRequest{}
}

// UpdateStatefulSetRequest初始化函数
func NewUpdateStatefulSetRequest() *UpdateStatefulSetRequest {
	return &UpdateStatefulSetRequest{
		StatefulSet: NewCreateStatefulSetRequest(),
	}
}

// QueryStatefulSetRequest初始化函数
func NewQueryStatefulSetRequest() *QueryStatefulSetRequest {
	return &QueryStatefulSetRequest{}
}

// DescribeStatefulSetRequest初始化函数
func NewDescribeStatefulSetRequest() *DescribeStatefulSetRequest {
	return &DescribeStatefulSetRequest{}
}

// 从restful解析参数初始化DeleteStatefulSetRequest
func NewDeleteStatefulSetRequestFromRestful(r *restful.Request) *DeleteStatefulSetRequest {
	return &DeleteStatefulSetRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从restful解析参数初始化QueryStatefulSetRequest
func NewQueryStatefulSetRequestFromRestful(r *restful.Request) *QueryStatefulSetRequest {
	return &QueryStatefulSetRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful解析参数初始化DescribeStatefulSetRequest
func NewDescribeStatefulSetRequestFromRestful(r *restful.Request) *DescribeStatefulSetRequest {
	return &DescribeStatefulSetRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}
