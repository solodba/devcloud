package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/role"
	"github.com/solodba/devcloud/mcenter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc role.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(role.AppName).(role.Service)
}
