package pod

import (
	context "context"
)

// 模块名称
const (
	AppName = "pod"
)

// 定义业务接口
type Service interface {
	// 创建Pod
	CreatePod(context.Context, *CreatePodRequest) (*Pod, error)
	// 删除Pod
	DeletePod(context.Context, *DeletePodRequest) (*Pod, error)
	// 修改Pod
	UpdatePod(context.Context, *UpdatePodRequest) (*Pod, error)
	// 查询Pod
	QueryPod(context.Context, *QueryPodRequest) (*PodSet, error)
	// 查询Pod详情
	DescribePod(context.Context, *DescribePodRequest) (*Pod, error)
	RPCServer
}
