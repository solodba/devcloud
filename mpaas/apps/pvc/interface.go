package pvc

import "github.com/emicklei/go-restful/v3"

// 模块名称
const (
	AppName = "pvc"
)

// PersistentVolumeClaim相关功能接口
type Service interface {
	// 嵌套PersistentVolumeClaim GRPC接口
	RPCServer
}

// CreatePVCRequest初始化函数
func NewCreatePVCRequest() *CreatePVCRequest {
	return &CreatePVCRequest{
		Labels:      make([]*ListMapItem, 0),
		AccessModes: make([]string, 0),
		Selector:    make([]*ListMapItem, 0),
	}
}

// DeletePVCRequest初始化函数
func NewDeletePVCRequest() *DeletePVCRequest {
	return &DeletePVCRequest{}
}

// QueryPVCRequest初始化函数
func NewQueryPVCRequest() *QueryPVCRequest {
	return &QueryPVCRequest{}
}

// 从restful解析参数初始化QueryPVCRequest结构体
func NewQueryPVCRequestFromRestful(r *restful.Request) *QueryPVCRequest {
	return &QueryPVCRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful解析参数初始化DeletePVCRequest结构体
func NewDeletePVCRequestFromRestful(r *restful.Request) *DeletePVCRequest {
	return &DeletePVCRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}
