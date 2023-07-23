package pvc

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
