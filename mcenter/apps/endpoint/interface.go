package endpoint

// 模块名称
const (
	AppName = "endpoint"
)

// Endpoint服务接口
type Service interface {
	// 嵌套Endpoint GRPC接口
	RPCServer
}
