package node

import "github.com/emicklei/go-restful/v3"

// 模块名称
const (
	AppName = "node"
)

// Node相关功能接口
type Service interface {
	// 嵌套Node GRPC接口
	RPCServer
}

// QueryNodeRequest结构体初始化函数
func NewQueryNodeRequest(keyword string) *QueryNodeRequest {
	return &QueryNodeRequest{
		Keyword: keyword,
	}
}

// DescribeNodeRequest结构体初始化函数
func NewDescribeNodeRequest(nodeName string) *DescribeNodeRequest {
	return &DescribeNodeRequest{
		Name: nodeName,
	}
}

// UpdatedLabelRequest结构体初始化函数
func NewUpdatedLabelRequest() *UpdatedLabelRequest {
	return &UpdatedLabelRequest{}
}

// UpdatedTaintRequest结构体初始化函数
func NewUpdatedTaintRequest() *UpdatedTaintRequest {
	return &UpdatedTaintRequest{}
}

// 从restful获取参数初始化QueryNodeRequest结构体
func NewQueryNodeRequestFromRestful(r *restful.Request) *QueryNodeRequest {
	return &QueryNodeRequest{
		Keyword: r.QueryParameter("keyword"),
	}
}

// 从restful获取参数初始化DescribeNodeRequest结构体
func NewDescribeNodeRequestFromRestful(r *restful.Request) *DescribeNodeRequest {
	return &DescribeNodeRequest{
		Name: r.QueryParameter("name"),
	}
}
