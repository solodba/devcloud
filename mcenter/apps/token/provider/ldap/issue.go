package ldap

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/apps/token/provider"
	"github.com/solodba/devcloud/mcenter/apps/user"
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
	return nil, nil
}

// 注册实例类
func init() {
	provider.Registry(token.GRANT_TYPE_LDAP, &issue{})
}
