package password

import (
	"context"

	"github.com/rs/xid"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/token/provider"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/user"
	"github.com/solodba/mcube/apps"
)

// 结构体实现类
type issue struct {
	// 引用user模块服务
	svc user.Service
}

// 实现Conf方法
func (i *issue) Conf() error {
	i.svc = apps.GetInternalApp(user.AppName).(user.Service)
	return nil
}

// 实现IssueToken方法
func (i *issue) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	req := user.NewDescribeUserRequest()
	req.DescribeValue = in.Username
	userIns, err := i.svc.DescribeUser(ctx, req)
	if err != nil {
		return nil, err
	}
	err = userIns.CheckPassword(in.Password)
	if err != nil {
		return nil, err
	}
	tk := token.NewToken()
	tk.AccessToken = xid.New().String()
	tk.RefreshToken = xid.New().String()
	return tk, nil
}

// 注册实例类
func init() {
	provider.Registry(token.GRANT_TYPE_PASSWORD, &issue{})
}
