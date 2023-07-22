package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/namespace"
	corev1 "k8s.io/api/core/v1"
)

// k8s中namespace结构体转换成自定义结构体
func (i *impl) NamespaceK8s2ResItemConvert(k8sNamespace corev1.Namespace) *namespace.NamespaceSetItem {
	return &namespace.NamespaceSetItem{
		Name:              k8sNamespace.Name,
		CreationTimestamp: k8sNamespace.CreationTimestamp.Unix(),
		Status:            string(k8sNamespace.Status.Phase),
	}
}
