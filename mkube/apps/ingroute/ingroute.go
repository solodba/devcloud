package ingroute

import (
	"github.com/solodba/mcube/pb/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sIngressRoute结构体
type K8sIngressRoute struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata"`
	Spec            *IngressRouteSpec `json:"spec"`
}

// K8sIngressRouteList结构体
type K8sIngressRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []*K8sIngressRoute `json:"items"`
}

// IngressRouteSet结构体初始化函数
func NewIngressRouteSet() *IngressRouteSet {
	return &IngressRouteSet{
		Items: make([]*IngressRouteSetItem, 0),
	}
}

// IngressRouteSetItem结构体初始化函数
func NewIngressRouteSetItem() *IngressRouteSetItem {
	return &IngressRouteSetItem{}
}

// IngressRouteSet结构体添加方法
func (i *IngressRouteSet) AddItems(items ...*IngressRouteSetItem) {
	i.Items = append(i.Items, items...)
}

// K8sIngressRoute结构体初始化函数
func NewK8sIngressRoute() *K8sIngressRoute {
	return &K8sIngressRoute{
		TypeMeta: metav1.TypeMeta{},
		Metadata: metav1.ObjectMeta{},
		Spec: &IngressRouteSpec{
			EntryPoints: make([]string, 0),
			Routes:      make([]*Routes, 0),
			Tls:         &Tls{},
		},
	}
}

// K8sIngressRouteList初始化函数
func NewK8sIngressRouteList() *K8sIngressRouteList {
	return &K8sIngressRouteList{
		TypeMeta: metav1.TypeMeta{},
		ListMeta: metav1.ListMeta{},
		Items:    make([]*K8sIngressRoute, 0),
	}
}

// K8sIngressRouteList添加方法
func (k *K8sIngressRouteList) AddItems(items ...*K8sIngressRoute) {
	k.Items = append(k.Items, items...)
}

// IngressRoute结构体初始化函数
func NewIngressRoute(req *CreateIngressRouteRequest) *IngressRoute {
	return &IngressRoute{
		Meta:         meta.NewMeta(),
		IngressRoute: req,
	}
}

// IngressRoute默认初始化函数
func NewDefaultIngressRoute() *IngressRoute {
	return &IngressRoute{
		Meta:         meta.NewMeta(),
		IngressRoute: NewCreateIngressRouteRequest(),
	}
}
