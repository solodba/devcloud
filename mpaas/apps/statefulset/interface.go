package statefulset

// 模块名
const (
	AppName = "statefulset"
)

// StatefulSet功能接口
type Service interface {
	// 嵌套StatefulSet GRPC接口
	RPCServer
}
