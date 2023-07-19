package configmap

// 模块名称
const (
	AppName = "configmap"
)

// ConfigMap业务接口
type Service interface {
	// 嵌套ConfigMap GRPC接口
	RPCServer
}

// CreateConfigMapRequest初始化函数
func NewCreateConfigMapRequest() *CreateConfigMapRequest {
	return &CreateConfigMapRequest{}
}

// UpdateConfigMapRequest初始化函数
func NewUpdateConfigMapRequest(req *CreateConfigMapRequest) *UpdateConfigMapRequest {
	return &UpdateConfigMapRequest{
		ConfigMap: req,
	}
}

// QueryConfigMapRequest初始化函数
func NewQueryConfigMapRequest() *QueryConfigMapRequest {
	return &QueryConfigMapRequest{}
}

// DescribeConfigMapRequest初始化函数
func NewDescribeConfigMapRequest() *DescribeConfigMapRequest {
	return &DescribeConfigMapRequest{}
}

// DeleteConfigMapRequest初始化函数
func NewDeleteConfigMapRequest() *DeleteConfigMapRequest {
	return &DeleteConfigMapRequest{}
}
