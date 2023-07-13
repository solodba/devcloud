package impl_test

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps"
	"github.com/solodba/devcloud/tree/main/mcenter/apps/token"
	"github.com/solodba/devcloud/tree/main/mcenter/test/tools"
)

var (
	svc token.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(token.AppName).(token.Service)
}
