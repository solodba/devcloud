package broker_test

import (
	"testing"

	"github.com/rs/xid"
	"github.com/solodba/devcloud/maudit/apps/audit"
)

func TestSendAuditLog(t *testing.T) {
	auditIns := audit.NewAuditLog()
	auditIns.Id = xid.New().String()
	auditIns.Operate = "POST"
	auditIns.Request = "/api/v1/cluster"
	auditIns.ServiceId = "abcdefghijk"
	auditIns.Time = 1231312345
	auditIns.Username = "admin"
	defer c.StopSendAuditLog()
	err := c.SendAuditLog(ctx, auditIns)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Producer发送日志到kafka成功!")
}
