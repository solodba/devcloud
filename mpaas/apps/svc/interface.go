package svc

import "github.com/emicklei/go-restful/v3"

// 模块名称
const (
	AppName = "svc"
)

// Service业务接口
type SvcService interface {
	// 嵌套Service RPC接口
	RPCServer
}

// CreateServiceRequest初始化函数
func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{
		Labels:   make([]*ListMapItem, 0),
		Selector: make([]*ListMapItem, 0),
		Ports:    make([]*ServicePort, 0),
	}
}

// UpdateServiceRequest结构体初始化函数
func NewUpdateServiceRequest(req *CreateServiceRequest) *UpdateServiceRequest {
	return &UpdateServiceRequest{
		Service: req,
	}
}

// QueryServiceRequest结构体初始化函数
func NewQueryServiceRequest() *QueryServiceRequest {
	return &QueryServiceRequest{}
}

// DescribeServiceRequest结构体初始化函数
func NewDescribeServiceRequest() *DescribeServiceRequest {
	return &DescribeServiceRequest{}
}

// DeleteServiceRequest结构体初始化函数
func NewDeleteServiceRequest() *DeleteServiceRequest {
	return &DeleteServiceRequest{}
}

// 从restful获取参数初始化DeleteServiceRequest
func NewDeleteServiceRequestFromRestful(r *restful.Request) *DeleteServiceRequest {
	return &DeleteServiceRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// UpdateServiceRequest结构体默认初始化函数
func NewDefaultUpdateServiceRequest() *UpdateServiceRequest {
	return &UpdateServiceRequest{
		Service: NewCreateServiceRequest(),
	}
}
