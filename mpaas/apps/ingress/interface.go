package ingress

// 模块名
const (
	AppName = "ingress"
)

// Ingress相关功能接口
type Service interface {
	// 嵌套Ingress GRPC接口
	RPCServer
}
