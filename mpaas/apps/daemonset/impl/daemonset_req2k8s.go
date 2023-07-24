package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/daemonset"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义DaemonSet结构体转换成K8S中结构体
func (i *impl) DaemonSetReq2K8sConvert(daemonSetReq *daemonset.CreateDaemonSetRequest) *appsv1.DaemonSet {
	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      daemonSetReq.Base.Name,
			Namespace: daemonSetReq.Base.Namespace,
			Labels:    i.getk8sDaemonSetLabels(daemonSetReq.Base.Labels),
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: i.getk8sDaemonSetSelector(daemonSetReq.Base.Selector),
			Template: i.getk8sDaemonSetTemplate(daemonSetReq.Template),
		},
	}
}

// 标签转换
func (i *impl) getk8sDaemonSetLabels(daemonSetReqLabels []*daemonset.ListMapItem) map[string]string {
	k8sDaemonSetLabels := make(map[string]string)
	for _, label := range daemonSetReqLabels {
		k8sDaemonSetLabels[label.Key] = label.Value
	}
	return k8sDaemonSetLabels
}

// 标签选择器转换
func (i *impl) getk8sDaemonSetSelector(daemonSetReqSelector []*daemonset.ListMapItem) *metav1.LabelSelector {
	var k8sDaemonSetLabelSelector metav1.LabelSelector
	k8sDaemonSetLabelSelector.MatchLabels = make(map[string]string)
	for _, selector := range daemonSetReqSelector {
		k8sDaemonSetLabelSelector.MatchLabels[selector.Key] = selector.Value
	}
	return &k8sDaemonSetLabelSelector
}

// Template转换
func (i *impl) getk8sDaemonSetTemplate(daemonSetReqPodTemplate *pod.Pod) corev1.PodTemplateSpec {
	var podTemplateSpec corev1.PodTemplateSpec
	k8sPod := i.podSvc.PodReq2K8s(daemonSetReqPodTemplate)
	podTemplateSpec.ObjectMeta = k8sPod.ObjectMeta
	podTemplateSpec.Spec = k8sPod.Spec
	return podTemplateSpec
}
