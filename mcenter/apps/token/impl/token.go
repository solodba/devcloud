package impl

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
)

// 颁发令牌
func (i *impl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	return nil, nil
}

// 令牌校验
func (i *impl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {
	return nil, nil
}
