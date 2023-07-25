package job

// 模块名称
const (
	AppName = "job"
)

// Job相关功能接口
type Service interface {
	// 嵌套Job GRPC接口
	RPCServer
}
