package sc

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
