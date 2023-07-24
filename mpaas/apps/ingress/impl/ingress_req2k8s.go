package impl

import (
	"github.com/solodba/devcloud/mpaas/apps/ingress"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 自定义结构体转换成k8s中结构体
func (i *impl) IngressReq2K8sConvert(ingressReq *ingress.CreateIngressRequest) *networkingv1.Ingress {
	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressReq.Name,
			Namespace: ingressReq.Namespace,
			Labels:    i.getK8sIngressLabels(ingressReq.Labels),
		},
		Spec: networkingv1.IngressSpec{
			Rules: i.getK8sIngressRules(ingressReq.Rules),
		},
	}
}

// 标签转换
func (i *impl) getK8sIngressLabels(ingressReqLabels []*ingress.ListMapItem) map[string]string {
	k8sIngressLabels := make(map[string]string)
	for _, label := range ingressReqLabels {
		k8sIngressLabels[label.Key] = label.Value
	}
	return k8sIngressLabels
}

// Rule转换
func (i *impl) getK8sIngressRules(ingressReqRules []*ingress.IngressRules) []networkingv1.IngressRule {
	k8sIngressRules := make([]networkingv1.IngressRule, 0)
	for _, rule := range ingressReqRules {
		k8sIngressRules = append(k8sIngressRules, networkingv1.IngressRule{
			Host:             rule.Host,
			IngressRuleValue: rule.Value,
		})
	}
	return k8sIngressRules
}
