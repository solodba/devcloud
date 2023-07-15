package impl

import (
	"context"
	"fmt"

	"github.com/solodba/devcloud/tree/main/mpaas/apps/pod"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 创建Pod
func (i *impl) CreatePod(ctx context.Context, in *pod.CreatePodRequest) (*pod.Pod, error) {
	// Pod必要参数校验
	if err := in.Pod.Validate(); err != nil {
		return nil, err
	}
	// Pod结构体转换
	k8sPod := in.Pod.PodReq2K8s()
	podApi := i.clientSet.CoreV1().Pods(k8sPod.Namespace)
	if getPod, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s]Pod已经存在!", getPod.Namespace, getPod.Name)
	}
	_, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s]Pod创建失败!", k8sPod.Namespace, k8sPod.Name)
	}
	return in.Pod, nil
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
