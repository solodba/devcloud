package impl

import (
	"context"

	"github.com/solodba/devcloud/maudit/apps/audit"
)

// 保存审计日志
func (i *impl) SaveAuditLog(ctx context.Context, in *audit.AuditLog) (*audit.AuditLog, error) {
	return nil, nil
}

// 查询审计日志
func (i *impl) QueryAuditLog(ctx context.Context, in *audit.QueryAuditLogRequest) (*audit.AuditLogSet, error) {
	return nil, nil
}
