package impl

import (
	"github.com/solodba/devcloud/mkube/apps/ingress"
	networkingv1 "k8s.io/api/networking/v1"
)

// K8S中结构体转换成自定义结构体
func (i *impl) IngressK8s2ReqConvert(k8sIngress *networkingv1.Ingress) *ingress.CreateIngressRequest {
	return &ingress.CreateIngressRequest{
		Name:      k8sIngress.Name,
		Namespace: k8sIngress.Namespace,
		Labels:    i.getIngressReqLabels(k8sIngress.Labels),
		Rules:     i.getIngressReqRules(k8sIngress.Spec.Rules),
	}
}

// 标签转换
func (i *impl) getIngressReqLabels(k8sIngressLabels map[string]string) []*ingress.ListMapItem {
	ingressReqLabels := make([]*ingress.ListMapItem, 0)
	for k, v := range k8sIngressLabels {
		ingressReqLabels = append(ingressReqLabels, &ingress.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return ingressReqLabels
}

// Rule转换
func (i *impl) getIngressReqRules(k8sIngressRules []networkingv1.IngressRule) []*ingress.IngressRules {
	ingressReqRules := make([]*ingress.IngressRules, 0)
	for _, rule := range k8sIngressRules {
		ingressReqRules = append(ingressReqRules, &ingress.IngressRules{
			Host:  rule.Host,
			Value: rule.IngressRuleValue,
		})
	}
	return ingressReqRules
}
