package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/pod"
	"github.com/solodba/devcloud/mpaas/apps/pvc"
	"github.com/solodba/devcloud/mpaas/apps/statefulset"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义StatefulSet结构体转换成K8S中结构体
func (i *impl) StatefulSetReq2K8sConvert(statefulSetReq *statefulset.CreateStatefulSetRequest) *appsv1.StatefulSet {
	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      statefulSetReq.Base.Name,
			Namespace: statefulSetReq.Base.Namespace,
			Labels:    i.getk8sStatefulSetLabels(statefulSetReq.Base.Labels),
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas:             &statefulSetReq.Base.Replicas,
			Selector:             i.getk8sStatefulSetSelector(statefulSetReq.Base.Selector),
			ServiceName:          statefulSetReq.Base.ServiceName,
			Template:             i.getk8sStatefulSetTemplate(statefulSetReq.Template),
			VolumeClaimTemplates: i.getk8sStatefulSetVolumeClaimTemplates(statefulSetReq.Base.VolumeClaimTemplates),
		},
	}
}

// 标签转换
func (i *impl) getk8sStatefulSetLabels(statefulSetReqLables []*statefulset.ListMapItem) map[string]string {
	k8sStatefulSetLables := make(map[string]string)
	for _, label := range statefulSetReqLables {
		k8sStatefulSetLables[label.Key] = label.Value
	}
	return k8sStatefulSetLables
}

// 标签选择器转换
func (i *impl) getk8sStatefulSetSelector(statefulSetReqSelector []*statefulset.ListMapItem) *metav1.LabelSelector {
	var labelSelector metav1.LabelSelector
	k8sStatefulSetSelector := make(map[string]string)
	for _, label := range statefulSetReqSelector {
		k8sStatefulSetSelector[label.Key] = label.Value
	}
	labelSelector.MatchLabels = k8sStatefulSetSelector
	return &labelSelector
}

// Template转换
func (i *impl) getk8sStatefulSetTemplate(statefulSetReqTemplate *pod.Pod) corev1.PodTemplateSpec {
	var podTemplateSepc corev1.PodTemplateSpec
	k8sPodTemplate := i.podSvc.PodReq2K8s(statefulSetReqTemplate)
	if k8sPodTemplate == nil {
		return podTemplateSepc
	}
	podTemplateSepc.ObjectMeta = k8sPodTemplate.ObjectMeta
	podTemplateSepc.Spec = k8sPodTemplate.Spec
	return podTemplateSepc
}

// VolumeClaim转换
func (i *impl) getk8sStatefulSetVolumeClaimTemplates(volumeClaimTemplates []*pvc.CreatePVCRequest) []corev1.PersistentVolumeClaim {
	k8sPVC := make([]corev1.PersistentVolumeClaim, 0)
	for _, item := range volumeClaimTemplates {
		pvc := i.pvcSvc.PVCReq2K8s(item)
		if pvc != nil {
			k8sPVC = append(k8sPVC, *pvc)
		}
	}
	return k8sPVC
}
