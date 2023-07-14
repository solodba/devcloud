package pod

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pod自定义结构体转换成K8s中定义的结构体
func (p *Pod) PodReq2K8s() *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Base.Name,
			Namespace: p.Base.Namespace,
			Labels:    p.getK8sLabels(p.Base.Labels),
		},
		Spec: corev1.PodSpec{},
	}
}

// 标签转换
func (p *Pod) getK8sLabels(podReqLabels []*ListMapItem) map[string]string {
	podK8sLabels := make(map[string]string)
	for _, label := range podReqLabels {
		podK8sLabels[label.Key] = label.Value
	}
	return podK8sLabels
}
