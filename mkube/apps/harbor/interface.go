package harbor

import (
	"strconv"

	"github.com/emicklei/go-restful/v3"
)

// 模块名称
const (
	AppName = "harbor"
)

// Harbor功能接口
type Service interface {
	// 嵌套Harbor GRPC接口
	RPCServer
}

// QueryProjectsRequest初始化函数
func NewQueryProjectsRequest() *QueryProjectsRequest {
	return &QueryProjectsRequest{}
}

// 从restful解析参数初始化QueryProjectsRequest
func NewQueryProjectsRequestFromRestful(r *restful.Request) (*QueryProjectsRequest, error) {
	curPage := r.QueryParameter("currentPage")
	if curPage == "" {
		curPage = "1"
	}
	newCurPage, err := strconv.Atoi(curPage)
	if err != nil {
		return nil, err
	}
	pageSize := r.QueryParameter("pageSize")
	if pageSize == "" {
		pageSize = "10"
	}
	newPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}
	return &QueryProjectsRequest{
		CurPage:  int64(newCurPage),
		PageSize: int64(newPageSize),
		Keyword:  r.QueryParameter("keyword"),
	}, nil
}
