package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/secret"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义Secret结构体转换成K8S结构体
func (i *impl) SecretReq2K8sConvert(secret *secret.CreateSecretRequest) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret.Name,
			Namespace: secret.Namespace,
			Labels:    i.getK8sSecretLabels(secret.Labels),
		},
		Type:       corev1.SecretType(secret.Type),
		StringData: i.getK8sSecretData(secret.Data),
	}
}

// 标签转换
func (i *impl) getK8sSecretLabels(secretReqLabels []*secret.ListMapItem) map[string]string {
	k8sSecretLabels := make(map[string]string)
	for _, secret := range secretReqLabels {
		k8sSecretLabels[secret.Key] = secret.Value
	}
	return k8sSecretLabels
}

// Data转换
func (i *impl) getK8sSecretData(secretReqData []*secret.ListMapItem) map[string]string {
	k8sSecretData := make(map[string]string)
	for _, data := range secretReqData {
		k8sSecretData[data.Key] = data.Value
	}
	return k8sSecretData
}
