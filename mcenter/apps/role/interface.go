package role

// 模块名称
const (
	AppName = "role"
)

// Role相关业务接口
type Service interface {
	// 嵌套role grpc接口
	RPCServer
}
