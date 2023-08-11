package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/service"
	"github.com/solodba/devcloud/mcenter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc service.ServiceManager
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(service.AppName).(service.ServiceManager)
}
