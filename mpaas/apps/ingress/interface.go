package ingress

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
