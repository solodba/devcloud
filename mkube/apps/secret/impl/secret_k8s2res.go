package impl

import (
	"github.com/solodba/devcloud/mkube/apps/secret"
	corev1 "k8s.io/api/core/v1"
)

// K8S中Secret结构体转换成自定义结构体
func (i *impl) SecretK8s2ResItemConvert(k8sSecret corev1.Secret) *secret.SecretSetItem {
	return &secret.SecretSetItem{
		Name:      k8sSecret.Name,
		Namespace: k8sSecret.Namespace,
		Type:      string(k8sSecret.Type),
		DataNum:   int64(len(k8sSecret.Data)),
		Age:       k8sSecret.CreationTimestamp.Unix(),
	}
}

// K8S中Secret结构体转换成自定义详情结构体
func (i *impl) SecretK8s2ResDetailConvert(secret corev1.Secret) *secret.SecretSetItem {
	secretRes := i.SecretK8s2ResItemConvert(secret)
	secretRes.Data = i.getSecretResData(secret.Data)
	secretRes.Labels = i.getSecretResLabels(secret.Labels)
	return secretRes
}

// Secret Data转换
func (i *impl) getSecretResData(k8sSecretData map[string][]byte) []*secret.ListMapItem {
	secretResData := make([]*secret.ListMapItem, 0)
	for k, v := range k8sSecretData {
		secretResData = append(secretResData, &secret.ListMapItem{
			Key:   k,
			Value: string(v),
		})
	}
	return secretResData
}

// 标签转换
func (i *impl) getSecretResLabels(k8sSecretLabels map[string]string) []*secret.ListMapItem {
	secretResLabels := make([]*secret.ListMapItem, 0)
	for k, v := range k8sSecretLabels {
		secretResLabels = append(secretResLabels, &secret.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return secretResLabels
}
