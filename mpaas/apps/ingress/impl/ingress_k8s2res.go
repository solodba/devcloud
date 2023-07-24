package impl

import (
	"strings"

	"github.com/solodba/devcloud/mpaas/apps/ingress"
	networkingv1 "k8s.io/api/networking/v1"
)

// k8s中Ingress结构体转换成自定义结构体
func (i *impl) IngressK8s2ResConvert(k8sIngress networkingv1.Ingress) *ingress.IngressSetItem {
	return &ingress.IngressSetItem{
		Name:      k8sIngress.Name,
		Namespace: k8sIngress.Namespace,
		Hosts:     i.getIngressResHosts(k8sIngress.Spec.Rules),
		Age:       k8sIngress.CreationTimestamp.Unix(),
	}
}

// k8s中Hosts结构体转换成自定义结构体
func (i *impl) getIngressResHosts(k8sIngressRules []networkingv1.IngressRule) string {
	ingressRuleHosts := make([]string, 0)
	for _, rule := range k8sIngressRules {
		ingressRuleHosts = append(ingressRuleHosts, rule.Host)
	}
	return strings.Join(ingressRuleHosts, ",")
}
