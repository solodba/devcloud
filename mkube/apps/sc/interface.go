package sc

import (
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/conf"
)

// 模块名称
const (
	AppName = "sc"
)

// StorageClass相关功能接口
type Service interface {
	// 嵌套StorageClass GRPC接口
	RPCServer
}

// CreateSCRequest初始化函数
func NewCreateSCRequest() *CreateSCRequest {
	return &CreateSCRequest{
		Labels:       make([]*ListMapItem, 0),
		MountOptions: make([]string, 0),
		Parameters:   make([]*ListMapItem, 0),
	}
}

// DeleteSCRequest初始化函数
func NewDeleteSCRequest(name string) *DeleteSCRequest {
	return &DeleteSCRequest{
		Name: name,
	}
}

// QuerySCRequest初始化函数
func NewQuerySCRequest(keyword string) *QuerySCRequest {
	return &QuerySCRequest{
		Keyword: keyword,
	}
}

// CreateSCRequest添加方法
func (c *CreateSCRequest) ValidatePlugin(conf *conf.Config) bool {
	provisionerList := strings.Split(conf.Plugin.Provisionser, ",")
	for _, provisioner := range provisionerList {
		if c.Provisioner == provisioner {
			return true
		}
	}
	return false
}

// 从restful解析参数初始化DeleteSCRequest结构体
func NewDeleteSCRequestFromRestful(r *restful.Request) *DeleteSCRequest {
	return &DeleteSCRequest{
		Name: r.PathParameter("name"),
	}
}

// 从restful解析参数初始化QuerySCRequest结构体
func NewQuerySCRequestFromRestful(r *restful.Request) *QuerySCRequest {
	return &QuerySCRequest{
		Keyword: r.QueryParameter("keyword"),
	}
}
