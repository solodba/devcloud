package audit

import "context"

// 模块名称
const (
	AppName = "audit"
)

// 审计业务功能接口
type Service interface {
	SaveAuditLog(context.Context, *AuditLog) (*AuditLog, error)
	QueryAuditLog(context.Context, *QueryAuditLogRequest) (*AuditLogSet, error)
}

// AuditLog结构体
type AuditLog struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Time      int64  `json:"time"`
	ServiceId string `json:"service_id"`
	Operate   string `json:"operate"`
	Request   string `json:"request"`
}

// QueryAuditLogRequest结构体
type QueryAuditLogRequest struct {
}
