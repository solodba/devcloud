package svc

// 模块名称
const (
	AppName = "svc"
)

// Service业务接口
type SvcService interface {
	// 嵌套Service RPC接口
	RPCServer
}
