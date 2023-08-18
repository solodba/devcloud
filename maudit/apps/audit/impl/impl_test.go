package impl_test

import (
	"context"

	"github.com/solodba/devcloud/maudit/apps/audit"
	"github.com/solodba/devcloud/maudit/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc audit.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(audit.AppName).(audit.Service)
}
