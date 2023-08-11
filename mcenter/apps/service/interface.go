package service

import context "context"

// 模块名称
const (
	AppName = "service"
)

// Service服务管理接口
type ServiceManager interface {
	// 创建服务
	CreateService(context.Context, *CreateServiceRequest) (*Service, error)
	// 嵌套Service GRPC接口
	RPCServer
}

// CreateServiceRequest构造函数
func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{}
}

// DescribeServiceRequest构造函数
func NewDescribeServiceRequest() *DescribeServiceRequest {
	return &DescribeServiceRequest{
		DescribeType: DESCRIBE_BY_SERVICE_ID,
	}
}
