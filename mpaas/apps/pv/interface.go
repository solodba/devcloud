package pv

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
