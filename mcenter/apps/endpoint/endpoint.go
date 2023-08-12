package endpoint

import "github.com/solodba/mcube/pb/meta"

// EndpointSet构造函数
func NewEndpointSet() *EndpointSet {
	return &EndpointSet{
		Items: make([]*Endpoint, 0),
	}
}

// EndpointSet添加方法
func (e *EndpointSet) AddItems(items ...*Endpoint) {
	e.Items = append(e.Items, items...)
}

// Endpoint构造函数
func NewEndpoint(req *CreateEndpointRequest) *Endpoint {
	return &Endpoint{
		Meta: meta.NewMeta(),
		Spec: req,
	}
}
