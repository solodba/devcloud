package impl

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mpaas/apps/pod"
)

// 创建Pod
func (i *impl) CreatePod(ctx context.Context, in *pod.CreatePodRequest) (*pod.Pod, error) {
	return nil, nil
}

// 删除Pod
func (i *impl) DeletePod(ctx context.Context, in *pod.DeletePodRequest) (*pod.Pod, error) {
	return nil, nil
}

// 修改Pod
func (i *impl) UpdatePod(ctx context.Context, in *pod.UpdatePodRequest) (*pod.Pod, error) {
	return nil, nil
}

// 查询Pod
func (i *impl) QueryPod(ctx context.Context, in *pod.QueryPodRequest) (*pod.PodSet, error) {
	return nil, nil
}

// 查询Pod详情
func (i *impl) DescribePod(ctx context.Context, in *pod.DescribePodRequest) (*pod.Pod, error) {
	return nil, nil
}
