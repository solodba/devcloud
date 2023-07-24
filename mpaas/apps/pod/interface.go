package pod

import (
	"github.com/emicklei/go-restful/v3"
	corev1 "k8s.io/api/core/v1"
)

// 模块名称
const (
	AppName = "pod"
)

// 定义业务接口
type Service interface {
	RPCServer
	PodReq2K8s(*Pod) *corev1.Pod
	PodK8s2Req(*corev1.Pod) *Pod
}

// CreatePodRequest初始化函数
func NewCreatePodRequest() *CreatePodRequest {
	return &CreatePodRequest{
		Pod: NewDefaultPod(),
	}
}

// QueryPodRequest初始化函数
func NewQueryPodRequest() *QueryPodRequest {
	return &QueryPodRequest{}
}

// DescribePodRequest初始化函数
func NewDescribePodRequest() *DescribePodRequest {
	return &DescribePodRequest{}
}

// DeletePodRequest初始化函数
func NewDeletePodRequest() *DeletePodRequest {
	return &DeletePodRequest{}
}

// UpdatePodRequest初始化函数
func NewUpdatePodRequest(pod *Pod) *UpdatePodRequest {
	return &UpdatePodRequest{
		Pod: pod,
	}
}

// 从前端获取参数初始化DeletePodRequest
func NewDeletePodRequestFromRestful(r *restful.Request) *DeletePodRequest {
	return &DeletePodRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从前端获取参数初始化DeletePodRequest
func NewQueryPodRequestFromRestful(r *restful.Request) *QueryPodRequest {
	return &QueryPodRequest{
		Namespace: r.PathParameter("namespace"),
		NodeName:  r.QueryParameter("nodename"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从前端获取参数初始化DescribePodRequest
func NewDescribePodRequestFromRestful(r *restful.Request) *DescribePodRequest {
	return &DescribePodRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}
