package ingroute

// 模块名称
const (
	AppName = "ingroute"
)

// IngressRoute业务接口
type Service interface {
	// 嵌套IngressRoute GRPC接口
	RPCServer
}
