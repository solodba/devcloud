package configmap

// 模块名称
const (
	AppName = "configmap"
)

// ConfigMap业务接口
type Service interface {
	// 嵌套ConfigMap GRPC接口
	RPCServer
}
