package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/daemonset"
	"github.com/solodba/devcloud/mpaas/apps/pod"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8S DaemonSet结构体转换成自定义结构体
func (i *impl) DaemonSetK8s2ReqConvert(k8sDaemonSet *appsv1.DaemonSet) *daemonset.CreateDaemonSetRequest {
	return &daemonset.CreateDaemonSetRequest{
		Base: &daemonset.Base{
			Name:      k8sDaemonSet.Name,
			Namespace: k8sDaemonSet.Namespace,
			Labels:    i.getDaemonSetReqLabels(k8sDaemonSet.Labels),
			Selector:  i.getDaemonSetReqSelector(k8sDaemonSet.Spec.Selector),
		},
		Template: i.getDaemonSetReqTemplate(k8sDaemonSet.Spec.Template),
	}
}

// 标签转换
func (i *impl) getDaemonSetReqLabels(k8sDaemonSetLabels map[string]string) []*daemonset.ListMapItem {
	daemonSetReqLabels := make([]*daemonset.ListMapItem, 0)
	for k, v := range k8sDaemonSetLabels {
		daemonSetReqLabels = append(daemonSetReqLabels, &daemonset.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return daemonSetReqLabels
}

// 标签选择器转换
func (i *impl) getDaemonSetReqSelector(k8sDaemonSetSelector *metav1.LabelSelector) []*daemonset.ListMapItem {
	daemonSetReqSelector := make([]*daemonset.ListMapItem, 0)
	if k8sDaemonSetSelector == nil {
		return daemonSetReqSelector
	}
	if k8sDaemonSetSelector.MatchLabels == nil {
		return daemonSetReqSelector
	}
	for k, v := range k8sDaemonSetSelector.MatchLabels {
		daemonSetReqSelector = append(daemonSetReqSelector, &daemonset.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return daemonSetReqSelector
}

// Template转换
func (i *impl) getDaemonSetReqTemplate(k8sDaemonSetTemplate corev1.PodTemplateSpec) *pod.Pod {
	var k8sPod corev1.Pod
	k8sPod.TypeMeta.APIVersion = "v1"
	k8sPod.TypeMeta.Kind = "Pod"
	k8sPod.ObjectMeta = k8sDaemonSetTemplate.ObjectMeta
	k8sPod.Spec = k8sDaemonSetTemplate.Spec
	podReq := i.podSvc.PodK8s2Req(&k8sPod)
	return podReq
}
