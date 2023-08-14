package policy

import "github.com/solodba/mcube/pb/page"

// 模块名称
const (
	AppName = "policy"
)

// policy接口服务
type Service interface {
	// 嵌套policy grpc接口
	RPCServer
}

// CreatePolicyRequest构造函数
func NewCreatePolicyRequest() *CreatePolicyRequest {
	return &CreatePolicyRequest{}
}

// QueryPolicyRequest构造函数
func NewQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page: page.NewPageRequest(),
	}
}
