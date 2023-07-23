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
