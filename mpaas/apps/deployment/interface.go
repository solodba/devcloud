package deployment

// 模块名称
const (
	AppName = "deployment"
)

// Deployment相关功能接口
type Service interface {
	// 嵌套Deployment GRPC接口
	RPCServer
}
