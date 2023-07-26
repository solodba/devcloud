package impl

import (
	"strconv"

	"github.com/solodba/devcloud/mkube/apps/pvc"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义PersistentVolumeClaim结构体转换成K8S中结构体
func (i *impl) PVCReq2K8s(pvcReq *pvc.CreatePVCRequest) *corev1.PersistentVolumeClaim {
	return &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcReq.Name,
			Namespace: pvcReq.Namespace,
			Labels:    i.getK8sPVCLabels(pvcReq.Labels),
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes:      i.getK8sPVCAccessModes(pvcReq.AccessModes),
			Resources:        i.getK8sPVCResources(pvcReq.Capacity),
			Selector:         i.getK8sPVCSelector(pvcReq.Selector),
			StorageClassName: &pvcReq.StorageClassName,
		},
	}
}

// AccessModes转换
func (i *impl) getK8sPVCAccessModes(pvcReqAccessModes []string) []corev1.PersistentVolumeAccessMode {
	k8sPVCAccessModes := make([]corev1.PersistentVolumeAccessMode, 0)
	for _, item := range pvcReqAccessModes {
		k8sPVCAccessModes = append(k8sPVCAccessModes, corev1.PersistentVolumeAccessMode(item))
	}
	return k8sPVCAccessModes
}

// Labels转换
func (i *impl) getK8sPVCLabels(pvcReqLabels []*pvc.ListMapItem) map[string]string {
	k8sPVCLabels := make(map[string]string)
	for _, label := range pvcReqLabels {
		k8sPVCLabels[label.Key] = label.Value
	}
	return k8sPVCLabels
}

// Resources转换
func (i *impl) getK8sPVCResources(pvcReqCapacity int32) corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(pvcReqCapacity)) + "Mi"),
		},
	}
}

// Selector转换
func (i *impl) getK8sPVCSelector(pvcReqSelector []*pvc.ListMapItem) *metav1.LabelSelector {
	k8sPVCMatchLabels := make(map[string]string)
	if len(pvcReqSelector) == 0 {
		return nil
	}
	for _, selector := range pvcReqSelector {
		k8sPVCMatchLabels[selector.Key] = selector.Value
	}
	return &metav1.LabelSelector{
		MatchLabels: k8sPVCMatchLabels,
	}
}
