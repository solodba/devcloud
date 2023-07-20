package secret

// 模块名称
const (
	AppName = "secret"
)

// Secret业务接口
type Service interface {
	// 嵌套Secret GRPC接口
	RPCServer
}
