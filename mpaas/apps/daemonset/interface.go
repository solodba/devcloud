package daemonset

// 模块名称
const (
	AppName = "daemonset"
)

// DaemonSet功能接口
type Service interface {
	// 嵌套DaemonSet GRPC接口
	RPCServer
}
