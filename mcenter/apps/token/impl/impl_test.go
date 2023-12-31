package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/token"
	"github.com/solodba/devcloud/mcenter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc token.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(token.AppName).(token.Service)
}
