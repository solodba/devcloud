package audit

import (
	context "context"
	"strconv"

	"github.com/emicklei/go-restful/v3"
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

// 从restful解析分页参数构造结构体QueryAuditLogRequest
func NewQueryAuditLogRequestFromRestful(r *restful.Request) *QueryAuditLogRequest {
	pageReq := page.NewPageRequest()
	pageNumber, _ := strconv.Atoi(r.QueryParameter("page_number"))
	pageReq.PageNumber = int64(pageNumber)
	pageSize, _ := strconv.Atoi(r.QueryParameter("page_size"))
	pageReq.PageSize = int64(pageSize)
	return &QueryAuditLogRequest{
		Page: pageReq,
	}
}
