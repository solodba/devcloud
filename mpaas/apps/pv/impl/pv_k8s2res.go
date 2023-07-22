package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/pv"
	corev1 "k8s.io/api/core/v1"
)

// K8S PersistentVolume结构体转换成自定义结构体
func (i *impl) PVK8s2Res(k8sPV corev1.PersistentVolume) *pv.PVSetItem {
	return &pv.PVSetItem{
		Name:             k8sPV.Name,
		Labels:           i.getPVResLabels(k8sPV.Labels),
		Capacity:         i.getPVResCapacity(k8sPV.Spec.Capacity),
		AccessModes:      i.getPVResAccessModes(k8sPV.Spec.AccessModes),
		ReClaimPolicy:    string(k8sPV.Spec.PersistentVolumeReclaimPolicy),
		Status:           string(k8sPV.Status.Phase),
		Claim:            i.getPVResClaim(k8sPV.Spec.ClaimRef),
		StorageClassName: k8sPV.Spec.StorageClassName,
		Age:              k8sPV.CreationTimestamp.Unix(),
		Reason:           k8sPV.Status.Reason,
	}
}

// AccessModes转换
func (i *impl) getPVResAccessModes(k8sPVAccessModes []corev1.PersistentVolumeAccessMode) []string {
	pvResAccessModes := make([]string, 0)
	for _, item := range k8sPVAccessModes {
		pvResAccessModes = append(pvResAccessModes, string(item))
	}
	return pvResAccessModes
}

// Labels转换
func (i *impl) getPVResLabels(k8sPVLables map[string]string) []*pv.ListMapItem {
	pvResLables := make([]*pv.ListMapItem, 0)
	for k, v := range k8sPVLables {
		pvResLables = append(pvResLables, &pv.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return pvResLables
}

// Capacity转换
func (i *impl) getPVResCapacity(k8sPVCapacity corev1.ResourceList) int32 {
	var pvResCapacity int32
	if k8sPVCapacity != nil {
		pvResCapacity = int32(k8sPVCapacity.Storage().Value() / (1024 * 1024))
	}
	return pvResCapacity
}

// Claim转换
func (i *impl) getPVResClaim(k8sPVClaim *corev1.ObjectReference) string {
	pvResClaim := ""
	if k8sPVClaim != nil {
		pvResClaim = k8sPVClaim.Name
	}
	return pvResClaim
}
