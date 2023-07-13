package token

import (
	context "context"

	"github.com/solodba/devcloud/tree/main/mcenter/common/validator"
)

// 定义模块名称
const (
	AppName = "token"
)

// 定义token业务接口
type Service interface {
	// 颁发令牌
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// GRP服务端接口
	RPCServer
}

// IssueTokenRequest结构体校验
func (req *IssueTokenRequest) Validate() error {
	return validator.V().Struct(req)
}

// IssueTokenRequest初始化函数
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{
		GrantType: GRANT_TYPE_PASSWORD,
	}
}
