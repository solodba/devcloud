package pv

import "github.com/emicklei/go-restful/v3"

// 模块名称
const (
	AppName = "pv"
)

// PersistentVolume相关功能接口
type Service interface {
	// 嵌套PersistentVolume GRPC接口
	RPCServer
}

// CreatePVRequest结构体初始化函数
func NewCreatePVRequest() *CreatePVRequest {
	return &CreatePVRequest{
		Labels:      make([]*ListMapItem, 0),
		AccessModes: make([]string, 0),
	}
}

// DeletePVRequest结构体初始化函数
func NewDeletePVRequest(name string) *DeletePVRequest {
	return &DeletePVRequest{
		Name: name,
	}
}

// PVSet结构体初始化函数
func NewPVSet() *PVSet {
	return &PVSet{
		Items: make([]*PVSetItem, 0),
	}
}

// PVSet结构体添加方法
func (p *PVSet) AddItems(items ...*PVSetItem) {
	p.Items = append(p.Items, items...)
}

// QueryPVRequest初始化函数
func NewQueryPVRequest(keyword string) *QueryPVRequest {
	return &QueryPVRequest{
		Keyword: keyword,
	}
}

// 从restful获取参数初始化DeletePVRequest结构体
func NewDeletePVRequestFromRestful(r *restful.Request) *DeletePVRequest {
	return &DeletePVRequest{
		Name: r.PathParameter("name"),
	}
}

// 从restful获取参数初始化QueryPVRequest结构体
func NewQueryPVRequestFromRestful(r *restful.Request) *QueryPVRequest {
	return &QueryPVRequest{
		Keyword: r.QueryParameter("keyword"),
	}
}
