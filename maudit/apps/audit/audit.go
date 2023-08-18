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
func (a *AuditLog) MustToJsonByte() []byte {
	dj, err := json.Marshal(a)
	if err != nil {
		logger.L().Panic().Msgf("auditlog parse to json failed, error: %s", err.Error())
	}
	return dj
}

// 通过json解析数据到结构体AuditLog
func (a *AuditLog) LoadAuditLogFromJson(data []byte) error {
	err := json.Unmarshal(data, a)
	if err != nil {
		return err
	}
	return nil
}

// AuditLogSet构造函数
func NewAuditLogSet() *AuditLogSet {
	return &AuditLogSet{
		Items: make([]*AuditLog, 0),
	}
}

// AuditLogSet添加方法
func (a *AuditLogSet) AddItems(items ...*AuditLog) {
	a.Items = append(a.Items, items...)
}
