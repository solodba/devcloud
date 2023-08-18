package impl_test

import (
	"testing"

	"github.com/solodba/devcloud/maudit/apps/audit"
	"github.com/solodba/devcloud/maudit/test/tools"
)

func TestSaveAuditLog(t *testing.T) {
	auditLogIns := audit.NewAuditLog()
	auditLogIns.Id = "asdasda123123123"
	auditLogIns.Operate = "POST"
	auditLogIns.Request = "/api/v1/cluster"
	auditLogIns.ServiceId = "aaaaaaaaaaaaaa"
	auditLogIns.Time = 23453451
	auditLogIns.Username = "admin"
	auditLogIns, err := svc.SaveAuditLog(ctx, auditLogIns)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(auditLogIns))
}

func TestQueryAuditLog(t *testing.T) {
	req := audit.NewQueryAuditLogRequest()
	// req.Page.PageNumber = 2
	// req.Page.PageSize = 5
	auditLogSet, err := svc.QueryAuditLog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(auditLogSet))
}
