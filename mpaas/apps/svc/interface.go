package svc

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
