package impl

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/token/provider"
	"go.mongodb.org/mongo-driver/bson"
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
	if err := in.Validate(); err != nil {
		return nil, err
	}
	filter := bson.M{}
	filter["_id"] = in.AccessToken
	tk := token.NewToken()
	err := i.col.FindOne(ctx, filter).Decode(tk)
	if err != nil {
		return nil, err
	}
	err = tk.ValidateToken()
	if err != nil {
		return nil, err
	}
	return tk, nil
}
