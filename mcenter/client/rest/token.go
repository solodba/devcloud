package rest

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
)

// 数据返回结构体
type Response struct {
	Code int          `json:"code"`
	Data *token.Token `json:"data"`
}

// 颁发令牌
func (c *Client) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	req := token.NewIssueTokenRequest()
	req.Username = "test2"
	req.Password = "123456"
	resp := &Response{Data: &token.Token{}}
	err := c.c.Post("token").Body(in).Do(ctx).Into(resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
