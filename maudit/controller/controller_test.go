package controller_test

import (
	"context"

	"github.com/solodba/devcloud/maudit/client/broker"
	"github.com/solodba/devcloud/maudit/controller"
	"github.com/solodba/devcloud/maudit/test/tools"
)

var (
	c   *controller.AuditLogSaveController
	ctx = context.Background()
)

func init() {
	c = controller.NewAuditLogSaveController(
		[]string{"192.168.1.130:9092"},
		"consumer-group-test",
		broker.DEFAULT_TOPIC_NAME,
	)
	tools.DevelopmentSet()
}
