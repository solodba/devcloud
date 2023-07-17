package provider

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/token"
)

// 定义Issuer接口
type Issuer interface {
	IssuerToken
	Conf() error
}

// 定义令牌颁发接口
type IssuerToken interface {
	// 颁发令牌
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}
