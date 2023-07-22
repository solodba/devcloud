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
