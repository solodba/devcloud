package endpoint

// 模块名称
const (
	AppName = "endpoint"
)

// Endpoint服务接口
type Service interface {
	// 嵌套Endpoint GRPC接口
	RPCServer
}

// RegistryEndpointRequest构造函数
func NewRegistryEndpointRequest() *RegistryEndpointRequest {
	return &RegistryEndpointRequest{
		Items: make([]*CreateEndpointRequest, 0),
	}
}

// RegistryEndpointRequest结构体添加方法
func (r *RegistryEndpointRequest) AddItems(item ...*CreateEndpointRequest) {
	r.Items = append(r.Items, item...)
}

// CreateEndpointRequest构造函数
func NewCreateEndpointRequest() *CreateEndpointRequest {
	return &CreateEndpointRequest{}
}
