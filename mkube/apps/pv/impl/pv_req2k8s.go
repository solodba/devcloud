package impl

import (
	"strconv"

	"github.com/solodba/devcloud/mkube/apps/pv"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	VOLUME_TYPE = "nfs"
)

// 自定义PersistentVolume结构体转换成K8S结构体
func (i *impl) PVReq2K8s(pvReq *pv.CreatePVRequest) *corev1.PersistentVolume {
	return &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvReq.Name,
			Labels: i.getK8sPVLabels(pvReq.Labels),
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity:                      i.getK8sPVCapacity(pvReq.Capacity),
			AccessModes:                   i.getK8sPVAccessModes(pvReq.AccessModes),
			PersistentVolumeReclaimPolicy: corev1.PersistentVolumeReclaimPolicy(pvReq.ReClaimPolicy),
			PersistentVolumeSource:        i.getK8sPersistentVolumeSource(pvReq.VolumeSource),
		},
	}
}

// AccessModes转换
func (i *impl) getK8sPVAccessModes(pvReqAccessModes []string) []corev1.PersistentVolumeAccessMode {
	k8sPVAccessModes := make([]corev1.PersistentVolumeAccessMode, 0)
	for _, item := range pvReqAccessModes {
		k8sPVAccessModes = append(k8sPVAccessModes, corev1.PersistentVolumeAccessMode(item))
	}
	return k8sPVAccessModes
}

// Labels转换
func (i *impl) getK8sPVLabels(pvReqLabels []*pv.ListMapItem) map[string]string {
	k8sPVLabels := make(map[string]string)
	for _, label := range pvReqLabels {
		k8sPVLabels[label.Key] = label.Value
	}
	return k8sPVLabels
}

// Capacity转换
func (i *impl) getK8sPVCapacity(pvReqCapacity int32) corev1.ResourceList {
	return corev1.ResourceList{
		corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(pvReqCapacity)) + "Mi"),
	}
}

// PersistentVolumeSource转换
func (i *impl) getK8sPersistentVolumeSource(pvReqVolumeSource *pv.VolumeSource) corev1.PersistentVolumeSource {
	var k8sPVSource corev1.PersistentVolumeSource
	switch pvReqVolumeSource.Type {
	case VOLUME_TYPE:
		k8sPVSource.NFS = &corev1.NFSVolumeSource{
			Server:   pvReqVolumeSource.NfsVolumeSource.NfsServer,
			Path:     pvReqVolumeSource.NfsVolumeSource.NfsPath,
			ReadOnly: pvReqVolumeSource.NfsVolumeSource.NfsReadOnly,
		}
	}
	return k8sPVSource
}
