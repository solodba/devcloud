package audit

import (
	context "context"

	"github.com/solodba/mcube/pb/page"
)

// 模块名称
const (
	AppName = "audit"
)

// 审计业务功能接口
type Service interface {
	// 保存审计日志
	SaveAuditLog(context.Context, *AuditLog) (*AuditLog, error)
	// 查询审计日志
	QueryAuditLog(context.Context, *QueryAuditLogRequest) (*AuditLogSet, error)
	// 嵌套maudit grpc接口
	RPCServer
}

// QueryAuditLogRequest构造函数
func NewQueryAuditLogRequest() *QueryAuditLogRequest {
	return &QueryAuditLogRequest{
		Page: page.NewPageRequest(),
	}
}
