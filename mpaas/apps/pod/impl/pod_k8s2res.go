package impl

import (
	"fmt"

	"github.com/solodba/devcloud/tree/main/mpaas/apps/pod"
	corev1 "k8s.io/api/core/v1"
)

// k8s中Pod结构体转换成自定义Pod结构体
func (i *impl) PodK8s2ItemRes(k8sPod corev1.Pod) *pod.PodListItem {
	var totalC, readyC, restartC int32
	for _, containerStatus := range k8sPod.Status.ContainerStatuses {
		if containerStatus.Ready {
			readyC++
		}
		restartC += containerStatus.RestartCount
		totalC++
	}
	var podStatus string
	if k8sPod.Status.Phase != "Running" {
		podStatus = "Error"
	} else {
		podStatus = "Running"
	}
	return &pod.PodListItem{
		Name:     k8sPod.Name,
		Ready:    fmt.Sprintf("%d/%d", readyC, totalC),
		Status:   podStatus,
		Restarts: restartC,
		Age:      k8sPod.CreationTimestamp.Unix(),
		IP:       k8sPod.Status.PodIP,
		Node:     k8sPod.Spec.NodeName,
	}
}
