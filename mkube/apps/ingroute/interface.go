package ingroute

import "github.com/emicklei/go-restful/v3"

// 模块名称
const (
	AppName = "ingroute"
)

// IngressRoute业务接口
type Service interface {
	// 嵌套IngressRoute GRPC接口
	RPCServer
}

// CreateIngressRouteRequest结构体初始化函数
func NewCreateIngressRouteRequest() *CreateIngressRouteRequest {
	return &CreateIngressRouteRequest{
		Labels: make([]*ListMapItem, 0),
		IngressRouteSpec: &IngressRouteSpec{
			EntryPoints: make([]string, 0),
			Routes:      make([]*Routes, 0),
			Tls:         &Tls{},
		},
	}
}

// DescribeIngressRouteRequest初始化函数
func NewDescribeIngressRouteRequest() *DescribeIngressRouteRequest {
	return &DescribeIngressRouteRequest{}
}

// QueryIngressRouteRequest初始化函数
func NewQueryIngressRouteRequest() *QueryIngressRouteRequest {
	return &QueryIngressRouteRequest{}
}

// UpdateIngressRouteRequest初始化函数
func NewUpdateIngressRouteRequest() *UpdateIngressRouteRequest {
	return &UpdateIngressRouteRequest{
		IngressRoute: NewCreateIngressRouteRequest(),
	}
}

// DeleteIngressRouteRequest初始化函数
func NewDeleteIngressRouteRequest() *DeleteIngressRouteRequest {
	return &DeleteIngressRouteRequest{}
}

// QueryIngRouteMwRequest初始化函数
func NewQueryIngRouteMwRequest(namespace string) *QueryIngRouteMwRequest {
	return &QueryIngRouteMwRequest{
		Namespace: namespace,
	}
}

// 从restful中解析参数DeleteIngressRouteRequest初始化函数
func NewDeleteIngressRouteRequestFromRestful(r *restful.Request) *DeleteIngressRouteRequest {
	return &DeleteIngressRouteRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从restful中解析参数初始化QueryIngressRouteRequest
func NewQueryIngressRouteRequestFromRestful(r *restful.Request) *QueryIngressRouteRequest {
	return &QueryIngressRouteRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful中解析参数初始化DescribeIngressRouteRequest
func NewDescribeIngressRouteRequestFromRestful(r *restful.Request) *DescribeIngressRouteRequest {
	return &DescribeIngressRouteRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}

// 从restful中解析参数初始化QueryIngRouteMwRequest
func NewQueryIngRouteMwRequestFromRestful(r *restful.Request) *QueryIngRouteMwRequest {
	return &QueryIngRouteMwRequest{
		Namespace: r.PathParameter("namespace"),
	}
}
