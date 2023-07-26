package impl

import (
	"github.com/solodba/devcloud/mkube/apps/pvc"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8S中PVC结构体转换成自定义结构体
func (i *impl) PVCK8s2Res(k8sPVC corev1.PersistentVolumeClaim) *pvc.PVCSetItem {
	return &pvc.PVCSetItem{
		Name:             k8sPVC.Name,
		Namespace:        k8sPVC.Namespace,
		Labels:           i.getPVCReqLabels(k8sPVC.Labels),
		AccessModes:      i.getPVCReqAccessModes(k8sPVC.Spec.AccessModes),
		Capacity:         i.getPVCReqCapacity(k8sPVC.Spec.Resources),
		Selector:         i.getPVCReqSelector(k8sPVC.Spec.Selector),
		StorageClassName: *k8sPVC.Spec.StorageClassName,
		Age:              k8sPVC.CreationTimestamp.Unix(),
		Status:           string(k8sPVC.Status.Phase),
		Volume:           k8sPVC.Spec.VolumeName,
	}
}

func (i *impl) PVCK8s2Req(k8sPVC corev1.PersistentVolumeClaim) *pvc.CreatePVCRequest {
	return &pvc.CreatePVCRequest{
		Name:             k8sPVC.Name,
		Namespace:        k8sPVC.Namespace,
		Labels:           i.getPVCReqLabels(k8sPVC.Labels),
		AccessModes:      i.getPVCReqAccessModes(k8sPVC.Spec.AccessModes),
		Capacity:         i.getPVCReqCapacity(k8sPVC.Spec.Resources),
		Selector:         i.getPVCReqSelector(k8sPVC.Spec.Selector),
		StorageClassName: *k8sPVC.Spec.StorageClassName,
	}
}

// AccessModes转换
func (i *impl) getPVCReqAccessModes(k8sPVCAccessModes []corev1.PersistentVolumeAccessMode) []string {
	pvcReqAccessModes := make([]string, 0)
	for _, item := range k8sPVCAccessModes {
		pvcReqAccessModes = append(pvcReqAccessModes, string(item))
	}
	return pvcReqAccessModes
}

// Labels转换
func (i *impl) getPVCReqLabels(k8sPVCLabels map[string]string) []*pvc.ListMapItem {
	pvcReqLabels := make([]*pvc.ListMapItem, 0)
	for k, v := range k8sPVCLabels {
		pvcReqLabels = append(pvcReqLabels, &pvc.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return pvcReqLabels
}

// Capacity转换
func (i *impl) getPVCReqCapacity(k8sPVCResources corev1.ResourceRequirements) int32 {
	var pvcReqCapacity int32
	if k8sPVCResources.Requests != nil {
		pvcReqCapacity = int32(k8sPVCResources.Requests.Storage().Value() / (1024 * 1024))
	}
	return pvcReqCapacity
}

// Selector转换
func (i *impl) getPVCReqSelector(k8sPVCSelector *metav1.LabelSelector) []*pvc.ListMapItem {
	pvcReqLabels := make([]*pvc.ListMapItem, 0)
	if k8sPVCSelector != nil {
		for k, v := range k8sPVCSelector.MatchLabels {
			pvcReqLabels = append(pvcReqLabels, &pvc.ListMapItem{
				Key:   k,
				Value: v,
			})
		}
	}
	return pvcReqLabels
}
