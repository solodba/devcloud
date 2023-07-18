package impl

import (
	"strings"

	"github.com/solodba/devcloud/mpaas/apps/svc"
	corev1 "k8s.io/api/core/v1"
)

// k8s中service结构体转换成自定义结构体
func (i *impl) SVCK8s2ResConvert(k8sSVC corev1.Service) *svc.ServiceListItem {
	return &svc.ServiceListItem{
		Name:       k8sSVC.Name,
		Type:       string(k8sSVC.Spec.Type),
		ClusterIP:  k8sSVC.Spec.ClusterIP,
		ExternalIP: strings.Join(k8sSVC.Spec.ExternalIPs, ","),
		Age:        k8sSVC.CreationTimestamp.Unix(),
	}
}
