package impl

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/token/provider"
)

// 颁发令牌
func (i *impl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	svc := provider.Get(in.GrantType).(provider.Issuer)
	tk, err := svc.IssueToken(ctx, in)
	if err != nil {
		return nil, err
	}
	_, err = i.col.InsertOne(ctx, tk)
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 令牌校验
func (i *impl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {
	return nil, nil
}
