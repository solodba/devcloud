package ingroute

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
