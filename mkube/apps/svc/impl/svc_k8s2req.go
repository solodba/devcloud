package impl

import (
	"github.com/solodba/devcloud/mkube/apps/svc"
	corev1 "k8s.io/api/core/v1"
)

// k8s中Service结构体转换成自定义结构体
func (i *impl) SVCK8s2Req(k8sSVC *corev1.Service) *svc.Service {
	return &svc.Service{
		Service: &svc.CreateServiceRequest{
			Name:      k8sSVC.Name,
			Namespace: k8sSVC.Namespace,
			Labels:    i.getSVCReqLabels(k8sSVC.Labels),
			Type:      string(k8sSVC.Spec.Type),
			Selector:  i.getSVCReqSelector(k8sSVC.Spec.Selector),
			Ports:     i.getSVCReqPorts(k8sSVC.Spec.Ports),
		},
	}
}

// 标签转换
func (i *impl) getSVCReqLabels(k8sSVCLabels map[string]string) []*svc.ListMapItem {
	svcReqLabels := make([]*svc.ListMapItem, 0)
	for k, v := range k8sSVCLabels {
		svcReqLabels = append(svcReqLabels, &svc.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return svcReqLabels
}

// 标签选择器转换
func (i *impl) getSVCReqSelector(k8sSVCSelector map[string]string) []*svc.ListMapItem {
	svcReqSelector := make([]*svc.ListMapItem, 0)
	for k, v := range k8sSVCSelector {
		svcReqSelector = append(svcReqSelector, &svc.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return svcReqSelector
}

// 端口转换
func (i *impl) getSVCReqPorts(k8sSVCPorts []corev1.ServicePort) []*svc.ServicePort {
	svcReqPorts := make([]*svc.ServicePort, 0)
	for _, port := range k8sSVCPorts {
		svcReqPorts = append(svcReqPorts, &svc.ServicePort{
			Name:       port.Name,
			Port:       port.Port,
			TargetPort: port.TargetPort.IntVal,
			NodePort:   port.NodePort,
		})
	}
	return svcReqPorts
}
