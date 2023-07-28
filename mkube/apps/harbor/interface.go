package harbor

// 模块名称
const (
	AppName = "harbor"
)

// Harbor功能接口
type Service interface {
	// 嵌套Harbor GRPC接口
	RPCServer
}
