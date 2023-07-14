package impl_test

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/user"
	"github.com/solodba/devcloud/tree/main/mcenter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc user.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(user.AppName).(user.Service)
}
