package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/configmap"
	corev1 "k8s.io/api/core/v1"
)

// K8S中ConfigMap结构体转换成自定义ConfigMap结构体
func (i *impl) GetCmResItem(configMap *corev1.ConfigMap) *configmap.ConfigMapSetItem {
	return &configmap.ConfigMapSetItem{
		Name:      configMap.Name,
		Namespace: configMap.Namespace,
		DataNum:   int64(len(configMap.Data)),
		Age:       configMap.CreationTimestamp.Unix(),
	}
}

// K8S中ConfgMap结构体转换成ConfigMap详情
func (i *impl) GetCmResDetail(configMap *corev1.ConfigMap) *configmap.ConfigMapSetItem {
	cm := i.GetCmResItem(configMap)
	configMapResLables := make([]*configmap.ListMapItem, 0)
	for k, v := range configMap.Labels {
		configMapResLables = append(configMapResLables, &configmap.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	cm.Labels = configMapResLables
	configMapResData := make([]*configmap.ListMapItem, 0)
	for k, v := range configMap.Data {
		configMapResData = append(configMapResData, &configmap.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	cm.Data = configMapResData
	return cm
}
