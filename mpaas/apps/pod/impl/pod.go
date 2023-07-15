package impl

import (
	"context"
	"fmt"
	"strings"

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
	k8sPod := i.PodReq2K8s(in.Pod)
	podApi := i.clientSet.CoreV1().Pods(k8sPod.Namespace)
	if getPod, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); err == nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] pod already exists", getPod.Namespace, getPod.Name)
	}
	_, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] create failed, err: %s", k8sPod.Namespace, k8sPod.Name, err.Error())
	}
	return in.Pod, nil
}

// 删除Pod
func (i *impl) DeletePod(ctx context.Context, in *pod.DeletePodRequest) (*pod.Pod, error) {
	req := pod.NewDescribePodRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	pod, err := i.DescribePod(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] pod not found, err: %s", in.Namespace, in.Name, err.Error())
	}
	var gracePeriodSeconds int64 = 0
	background := metav1.DeletePropagationBackground
	podApi := i.clientSet.CoreV1().Pods(pod.Base.Namespace)
	err = podApi.Delete(ctx, pod.Base.Name, metav1.DeleteOptions{
		GracePeriodSeconds: &gracePeriodSeconds,
		PropagationPolicy:  &background,
	})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete failed, err: %s", pod.Base.Namespace, pod.Base.Name, err.Error())
	}
	return pod, nil
}

// 修改Pod
func (i *impl) UpdatePod(ctx context.Context, in *pod.UpdatePodRequest) (*pod.Pod, error) {
	return nil, nil
}

// 查询Pod
func (i *impl) QueryPod(ctx context.Context, in *pod.QueryPodRequest) (*pod.PodSet, error) {
	list, err := i.clientSet.CoreV1().Pods(in.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("get pod list error, err: %s", err.Error())
	}
	podSet := pod.NewPodSet()
	for _, item := range list.Items {
		if in.NodeName != "" && item.Spec.NodeName != in.NodeName {
			continue
		}
		if strings.Contains(item.Name, in.Keyword) {
			// K8S中Pod转换成自定义Pod
			podItem := i.PodK8s2ItemRes(item)
			podSet.AddItems(podItem)
		}
	}
	podSet.Total = int64(len(podSet.PodListItem))
	return podSet, nil
}

// 查询Pod详情
func (i *impl) DescribePod(ctx context.Context, in *pod.DescribePodRequest) (*pod.Pod, error) {
	podApi := i.clientSet.CoreV1().Pods(in.Namespace)
	k8sGetPod, err := podApi.Get(ctx, in.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s,name=%s] get pod detail error, err: %s", in.Namespace, in.Name, err.Error())
	}
	// K8s中Pod转换成自定义Pod
	return i.PodK8s2Req(k8sGetPod), nil
}
