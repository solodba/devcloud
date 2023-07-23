package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/sc"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

// k8s中StorageClass结构体转换成自定义结构体
func (i *impl) SCK8s2Res(k8sSC storagev1.StorageClass) *sc.SCSetItem {
	return &sc.SCSetItem{
		Name:                 k8sSC.Name,
		Namespace:            k8sSC.Namespace,
		Labels:               i.getSCResLabels(k8sSC.Labels),
		Provisioner:          k8sSC.Provisioner,
		MountOptions:         k8sSC.MountOptions,
		Parameters:           i.getSCResParameters(k8sSC.Parameters),
		ReClaimPolicy:        i.getSCResReclaimPolicy(k8sSC.ReclaimPolicy),
		AllowVolumeExpansion: i.getSCResAllowVolumeExpansion(k8sSC.AllowVolumeExpansion),
		VolumeBindingMode:    i.getSCResVolumeBindingMode(k8sSC.VolumeBindingMode),
		Age:                  k8sSC.CreationTimestamp.Unix(),
	}
}

// Labels转换
func (i *impl) getSCResLabels(k8sSCLabels map[string]string) []*sc.ListMapItem {
	scResLabels := make([]*sc.ListMapItem, 0)
	for k, v := range k8sSCLabels {
		scResLabels = append(scResLabels, &sc.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return scResLabels
}

// Parameters转换
func (i *impl) getSCResParameters(k8sSCParameters map[string]string) []*sc.ListMapItem {
	scResPrameters := make([]*sc.ListMapItem, 0)
	for k, v := range k8sSCParameters {
		scResPrameters = append(scResPrameters, &sc.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return scResPrameters
}

// ReclaimPolicy转换
func (i *impl) getSCResReclaimPolicy(k8sSCReclaimPolicy *corev1.PersistentVolumeReclaimPolicy) string {
	var scResReclaimPolicy string
	if k8sSCReclaimPolicy != nil {
		scResReclaimPolicy = string(*k8sSCReclaimPolicy)
	}
	return scResReclaimPolicy
}

// AllowVolumeExpansion转换
func (i *impl) getSCResAllowVolumeExpansion(k8sSCAllowVolumeExpansion *bool) bool {
	var scResAllowVolumeExpansion bool
	if k8sSCAllowVolumeExpansion != nil {
		scResAllowVolumeExpansion = *k8sSCAllowVolumeExpansion
	}
	return scResAllowVolumeExpansion
}

// VolumeBindingMode转换
func (i *impl) getSCResVolumeBindingMode(k8sSCVolumeBindingMode *storagev1.VolumeBindingMode) string {
	var scResVolumeBindingMode string
	if k8sSCVolumeBindingMode != nil {
		scResVolumeBindingMode = string(*k8sSCVolumeBindingMode)
	}
	return scResVolumeBindingMode
}
