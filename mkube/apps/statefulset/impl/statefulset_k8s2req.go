package impl

import (
	"github.com/solodba/devcloud/mkube/apps/pod"
	"github.com/solodba/devcloud/mkube/apps/pvc"
	"github.com/solodba/devcloud/mkube/apps/statefulset"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// k8s中结构体转换成自定义结构体
func (i *impl) StatefulSetlK8s2ReqConvert(k8sStatefulSet *appsv1.StatefulSet) *statefulset.CreateStatefulSetRequest {
	return &statefulset.CreateStatefulSetRequest{
		Base: &statefulset.Base{
			Name:                 k8sStatefulSet.Name,
			Namespace:            k8sStatefulSet.Namespace,
			Labels:               i.getStatefulSetReqLabels(k8sStatefulSet.Labels),
			Replicas:             i.getStatefulSetReqReplicas(k8sStatefulSet.Spec.Replicas),
			Selector:             i.getStatefulSetReqSelector(k8sStatefulSet.Spec.Selector),
			ServiceName:          k8sStatefulSet.Spec.ServiceName,
			VolumeClaimTemplates: i.getStatefulSetReqVolumeClaimTemplates(k8sStatefulSet.Spec.VolumeClaimTemplates),
		},
		Template: i.getStatefulSetReqTemplate(k8sStatefulSet.Spec.Template),
	}
}

// 标签转换
func (i *impl) getStatefulSetReqLabels(k8sStatefulSetLabels map[string]string) []*statefulset.ListMapItem {
	statefulSetReqLabels := make([]*statefulset.ListMapItem, 0)
	for k, v := range k8sStatefulSetLabels {
		statefulSetReqLabels = append(statefulSetReqLabels, &statefulset.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return statefulSetReqLabels
}

// Replicas转换
func (i *impl) getStatefulSetReqReplicas(k8sStatefulSetReplicas *int32) int32 {
	var statefulSetReqReplicas int32
	if k8sStatefulSetReplicas == nil {
		return statefulSetReqReplicas
	}
	statefulSetReqReplicas = *k8sStatefulSetReplicas
	return statefulSetReqReplicas
}

// Selector转换
func (i *impl) getStatefulSetReqSelector(k8sStatefulSetSelector *metav1.LabelSelector) []*statefulset.ListMapItem {
	statefulSetSelector := make([]*statefulset.ListMapItem, 0)
	if k8sStatefulSetSelector == nil {
		return statefulSetSelector
	}
	if k8sStatefulSetSelector.MatchLabels == nil {
		return statefulSetSelector
	}
	for k, v := range k8sStatefulSetSelector.MatchLabels {
		statefulSetSelector = append(statefulSetSelector, &statefulset.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return statefulSetSelector
}

// VolumeClaim转换
func (i *impl) getStatefulSetReqVolumeClaimTemplates(k8sStatefulSetVolumeClaimTemplates []corev1.PersistentVolumeClaim) []*pvc.CreatePVCRequest {
	persistenVolumeClaimReq := make([]*pvc.CreatePVCRequest, 0)
	for _, item := range k8sStatefulSetVolumeClaimTemplates {
		// 数据转换
		pvc := i.pvcSvc.PVCK8s2Req(item)
		persistenVolumeClaimReq = append(persistenVolumeClaimReq, pvc)
	}
	return persistenVolumeClaimReq
}

// PodTemplate转换
func (i *impl) getStatefulSetReqTemplate(k8sStatefulSetTemplate corev1.PodTemplateSpec) *pod.Pod {
	var k8sPod corev1.Pod
	k8sPod.TypeMeta.APIVersion = "v1"
	k8sPod.TypeMeta.Kind = "Pod"
	k8sPod.ObjectMeta = k8sStatefulSetTemplate.ObjectMeta
	k8sPod.Spec = k8sStatefulSetTemplate.Spec
	podReq := i.podSvc.PodK8s2Req(&k8sPod)
	return podReq
}
