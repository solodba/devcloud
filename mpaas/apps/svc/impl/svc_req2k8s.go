package impl

import (
	"fmt"

	"github.com/solodba/devcloud/mpaas/apps/svc"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Service转换
func (i *impl) SVCReq2K8sConvert(svcReq *svc.CreateServiceRequest) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svcReq.Name,
			Namespace: svcReq.Namespace,
			Labels:    i.getK8sSVCLabels(svcReq.Labels),
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceType(svcReq.Type),
			Selector: i.getK8sSVCSelector(svcReq.Selector),
			Ports:    i.getK8sSVCPorts(svcReq.Ports),
		},
	}
}

// 标签转换
func (i *impl) getK8sSVCLabels(svcReqLables []*svc.ListMapItem) map[string]string {
	k8sSVCLabels := make(map[string]string)
	for _, label := range svcReqLables {
		k8sSVCLabels[label.Key] = label.Value
	}
	return k8sSVCLabels
}

// 标签选择器转换
func (i *impl) getK8sSVCSelector(svcReqSelector []*svc.ListMapItem) map[string]string {
	k8sSVCSelector := make(map[string]string)
	for _, selector := range svcReqSelector {
		k8sSVCSelector[selector.Key] = selector.Value
	}
	return k8sSVCSelector
}

// ServicePort转换
func (i *impl) getK8sSVCPorts(svcReqPorts []*svc.ServicePort) []corev1.ServicePort {
	k8sSVCPorts := make([]corev1.ServicePort, 0)
	for _, port := range svcReqPorts {
		k8sSVCPorts = append(k8sSVCPorts, corev1.ServicePort{
			Name:       port.Name,
			Port:       port.Port,
			TargetPort: intstr.FromInt(int(port.TargetPort)),
			NodePort:   port.NodePort,
		})
	}
	fmt.Println(k8sSVCPorts)
	return k8sSVCPorts
}
