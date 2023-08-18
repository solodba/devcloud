package audit

// AuditLogSet结构体
type AuditLogSet struct {
	Total int64       `json:"total"`
	Items []*AuditLog `json:"items"`
}
