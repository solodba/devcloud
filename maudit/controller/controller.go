package controller

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/solodba/devcloud/maudit/apps/audit"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
)

// auditlog controller结构体
type AuditLogSaveController struct {
	brokerAddress []string
	groupId       string
	topic         string
	r             *kafka.Reader
	audit         audit.Service
}

// AuditLogSaveController构造函数
func NewAuditLogSaveController(brokerAddress []string, groupId, topic string) *AuditLogSaveController {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokerAddress,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &AuditLogSaveController{
		brokerAddress: brokerAddress,
		groupId:       groupId,
		topic:         topic,
		r:             r,
		audit:         apps.GetInternalApp(audit.AppName).(audit.Service),
	}
}

// 启动Consumser消费端
func (a *AuditLogSaveController) Run(ctx context.Context) {
	for {
		m, err := a.r.ReadMessage(ctx)
		if err != nil {
			break
		}
		logger.L().Info().Msgf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		auditLogIns := audit.NewAuditLog()
		err = auditLogIns.LoadAuditLogFromJson(m.Value)
		if err != nil {
			logger.L().Error().Msgf("load audit log error, err: %s", err.Error())
			continue
		}
		_, err = a.audit.SaveAuditLog(ctx, auditLogIns)
		if err != nil {
			logger.L().Error().Msgf("save audit log error, err: %s", err.Error())
			continue
		}
	}
}

// 停止Consumer消费端
func (a *AuditLogSaveController) Stop() error {
	return a.r.Close()
}
