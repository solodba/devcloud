package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/permission"
	"github.com/solodba/devcloud/mcenter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc permission.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(permission.AppName).(permission.Service)
}
