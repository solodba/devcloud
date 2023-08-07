package ingroute

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sIngressRoute结构体
type K8sIngressRoute struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata"`
	Spec            IngressRouteSpec  `json:"spec"`
}

// K8sIngressRouteList结构体
type K8sIngressRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []IngressRoute `json:"items"`
}

// IngressRouteSet结构体初始化函数
func NewIngressRouteSet() *IngressRouteSet {
	return &IngressRouteSet{
		Items: make([]*IngressRouteSetItem, 0),
	}
}

// IngressRouteSet结构体添加方法
func (i *IngressRouteSet) AddItems(items ...*IngressRouteSetItem) {
	i.Items = append(i.Items, items...)
}

// K8sIngressRoute结构体初始化函数
func NewK8sIngressRoute() *K8sIngressRoute {
	return &K8sIngressRoute{
		TypeMeta: metav1.TypeMeta{},
		Spec: IngressRouteSpec{
			EntryPoints: make([]string, 0),
			Routes:      make([]*Routes, 0),
			Tls:         &Tls{},
		},
	}
}
