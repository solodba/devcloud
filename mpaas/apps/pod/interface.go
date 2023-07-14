package pod

// 模块名称
const (
	AppName = "pod"
)

// 定义业务接口
type Service interface {
	RPCServer
}
