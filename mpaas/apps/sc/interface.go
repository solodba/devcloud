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
