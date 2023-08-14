package impl_test

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/policy"
	"github.com/solodba/devcloud/mcenter/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc policy.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(policy.AppName).(policy.Service)
}
