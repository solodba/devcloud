package token

import context "context"

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
