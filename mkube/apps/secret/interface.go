package secret

import "github.com/emicklei/go-restful/v3"

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

// DeleteSecretRequest结构体
func NewDeleteSecretRequest() *DeleteSecretRequest {
	return &DeleteSecretRequest{}
}

// 从restful获取参数初始化QuerySecretRequest结构体
func NewQuerySecretRequestFromRestful(r *restful.Request) *QuerySecretRequest {
	return &QuerySecretRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful获取参数初始化DescribeSecretRequest结构体
func NewDescribeSecretRequestFromRestful(r *restful.Request) *DescribeSecretRequest {
	return &DescribeSecretRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}

// 从restful获取参数初始化DeleteSecretRequest结构体
func NewDeleteSecretRequestFromRestful(r *restful.Request) *DeleteSecretRequest {
	return &DeleteSecretRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}
