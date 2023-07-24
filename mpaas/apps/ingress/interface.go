package ingress

import "github.com/emicklei/go-restful/v3"

// 模块名
const (
	AppName = "ingress"
)

// Ingress相关功能接口
type Service interface {
	// 嵌套Ingress GRPC接口
	RPCServer
}

// CreateIngressRequest结构体初始化函数
func NewCreateIngressRequest() *CreateIngressRequest {
	return &CreateIngressRequest{
		Labels: make([]*ListMapItem, 0),
		Rules:  make([]*IngressRules, 0),
	}
}

// UpdateIngressRequest结构体初始化函数
func NewUpdateIngressRequest(req *CreateIngressRequest) *UpdateIngressRequest {
	return &UpdateIngressRequest{
		Ingress: req,
	}
}

// QueryIngressRequest初始化结构体函数
func NewQueryIngressRequest() *QueryIngressRequest {
	return &QueryIngressRequest{}
}

// DescribeIngressRequest初始化结构体函数
func NewDescribeIngressRequest() *DescribeIngressRequest {
	return &DescribeIngressRequest{}
}

// DeleteIngressRequest结构体初始化函数
func NewDeleteIngressRequest() *DeleteIngressRequest {
	return &DeleteIngressRequest{}
}

// 从restful解析参数初始化DeleteIngressRequest
func NewDeleteIngressRequestFromRestful(r *restful.Request) *DeleteIngressRequest {
	return &DeleteIngressRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从restful解析参数初始化QueryIngressRequest
func NewQueryIngressRequestFromRestful(r *restful.Request) *QueryIngressRequest {
	return &QueryIngressRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful解析参数初始化DescribeIngressRequest
func NewDescribeIngressRequestFromRestful(r *restful.Request) *DescribeIngressRequest {
	return &DescribeIngressRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}
