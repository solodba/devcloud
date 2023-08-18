package audit

import (
	"encoding/json"

	"github.com/solodba/mcube/logger"
)

// AuditLog结构体构造函数
func NewAuditLog() *AuditLog {
	return &AuditLog{}
}

// AuditLog添加json序列化方法
func (a *AuditLog) MustToJson() string {
	dj, err := json.Marshal(a)
	if err != nil {
		logger.L().Panic().Msgf("auditlog parse to json failed, error: %s", err.Error())
	}
	return string(dj)
}
