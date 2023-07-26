package impl

import (
	"github.com/solodba/devcloud/mkube/apps/configmap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义ConfigMap结构体转换成K8S中ConfigMap结构体
func (i *impl) CmReq2K8s(configMap *configmap.CreateConfigMapRequest) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMap.Name,
			Namespace: configMap.Namespace,
			Labels:    i.getK8sCmLabels(configMap.Labels),
		},
		Data: i.getK8sCmData(configMap.Data),
	}
}

// 标签转换
func (i *impl) getK8sCmLabels(cmReqLabels []*configmap.ListMapItem) map[string]string {
	k8sCmLabels := make(map[string]string)
	for _, label := range cmReqLabels {
		k8sCmLabels[label.Key] = label.Value
	}
	return k8sCmLabels
}

// Data转换
func (i *impl) getK8sCmData(cmReqData []*configmap.ListMapItem) map[string]string {
	k8sCmData := make(map[string]string)
	for _, data := range cmReqData {
		k8sCmData[data.Key] = data.Value
	}
	return k8sCmData
}
