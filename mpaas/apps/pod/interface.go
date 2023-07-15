package pod

import "github.com/solodba/mcube/pb/meta"

// 模块名称
const (
	AppName = "pod"
)

// 定义业务接口
type Service interface {
	RPCServer
}

// CreatePodRequest初始化函数
func NewCreatePodRequest() *CreatePodRequest {
	return &CreatePodRequest{
		Meta: meta.NewMeta(),
		Pod:  NewPod(),
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
