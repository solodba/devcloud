package cronjob

// 模块名称
const (
	AppName = "cronjob"
)

// CronJob相关功能接口
type Service interface {
	// 嵌套CronJob GRPC接口
	RPCServer
}
