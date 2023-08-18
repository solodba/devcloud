package impl

import (
	"context"

	"github.com/solodba/devcloud/maudit/apps/audit"
	"github.com/solodba/mcube/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 保存审计日志
func (i *impl) SaveAuditLog(ctx context.Context, in *audit.AuditLog) (*audit.AuditLog, error) {
	_, err := i.col.InsertOne(ctx, in)
	if err != nil {
		return nil, err
	}
	return in, nil
}

// 查询审计日志
func (i *impl) QueryAuditLog(ctx context.Context, in *audit.QueryAuditLogRequest) (*audit.AuditLogSet, error) {
	auditLogSet := audit.NewAuditLogSet()
	filter := bson.M{}
	opts := &options.FindOptions{}
	opts.SetLimit(in.Page.PageSize)
	in.Page.ComputeOffSet()
	opts.SetSkip(in.Page.Offset)
	cursor, err := i.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		auditLogIns := audit.NewAuditLog()
		err = cursor.Decode(auditLogIns)
		if err != nil {
			return nil, err
		}
		auditLogSet.AddItems(auditLogIns)
	}
	logger.L().Info().Msgf("%v", auditLogSet)
	auditLogSet.Total, err = i.col.CountDocuments(ctx, filter)
	return auditLogSet, nil
}
