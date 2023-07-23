package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/sc"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 从自定义StorageClass结构体转换成K8S中结构体
func (i *impl) SCReq2K8s(scReq *sc.CreateSCRequest) *storagev1.StorageClass {
	return &storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:      scReq.Name,
			Namespace: scReq.Namespace,
			Labels:    i.getK8sSCLabels(scReq.Labels),
		},
		Provisioner:          scReq.Provisioner,
		MountOptions:         scReq.MountOptions,
		Parameters:           i.getK8sSCParameters(scReq.Parameters),
		ReclaimPolicy:        i.getK8sSCReclaimPolicy(scReq.ReClaimPolicy),
		AllowVolumeExpansion: &scReq.AllowVolumeExpansion,
		VolumeBindingMode:    i.getK8sSCVolumeBindingMode(scReq.VolumeBindingMode),
	}
}

// ReclaimPolicy转换
func (i *impl) getK8sSCReclaimPolicy(scReqReclaimPolicy string) *corev1.PersistentVolumeReclaimPolicy {
	var k8sSCReclaimPolicy corev1.PersistentVolumeReclaimPolicy
	k8sSCReclaimPolicy = corev1.PersistentVolumeReclaimPolicy(scReqReclaimPolicy)
	return &k8sSCReclaimPolicy
}

// VolumeBindingMode转换
func (i *impl) getK8sSCVolumeBindingMode(scReqVolumeBindingMode string) *storagev1.VolumeBindingMode {
	var k8sSCVolumeBindingMode storagev1.VolumeBindingMode
	k8sSCVolumeBindingMode = storagev1.VolumeBindingMode(scReqVolumeBindingMode)
	return &k8sSCVolumeBindingMode
}

// Labels转换
func (i *impl) getK8sSCLabels(scReqLabels []*sc.ListMapItem) map[string]string {
	k8sSCLabels := make(map[string]string)
	for _, label := range scReqLabels {
		k8sSCLabels[label.Key] = label.Value
	}
	return k8sSCLabels
}

// Parameters转换
func (i *impl) getK8sSCParameters(scReqParameters []*sc.ListMapItem) map[string]string {
	k8sSCParameters := make(map[string]string)
	for _, parameter := range scReqParameters {
		k8sSCParameters[parameter.Key] = parameter.Value
	}
	return k8sSCParameters
}
