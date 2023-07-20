package secret

// 模块名称
const (
	AppName = "secret"
)

// Secret业务接口
type Service interface {
	// 嵌套Secret GRPC接口
	RPCServer
}

// CreateSecretRequest结构体
func NewCreateSecretRequest() *CreateSecretRequest {
	return &CreateSecretRequest{
		Labels: make([]*ListMapItem, 0),
		Data:   make([]*ListMapItem, 0),
	}
}

// UpdateSecretRequest结构体
func NewUpdateSecretRequest(req *CreateSecretRequest) *UpdateSecretRequest {
	return &UpdateSecretRequest{
		Secret: req,
	}
}

// QuerySecretRequest结构体
func NewQuerySecretRequest() *QuerySecretRequest {
	return &QuerySecretRequest{}
}

// DescribeSecretRequest结构体
func NewDescribeSecretRequest() *DescribeSecretRequest {
	return &DescribeSecretRequest{}
}
