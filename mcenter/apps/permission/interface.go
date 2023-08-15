package permission

// 模块名称
const (
	AppName = "permission"
)

// Permission业务接口
type Service interface {
	// 嵌套Permission Grpc接口
	RPCServer
}

// CheckPermissionRequest构造函数
func NewCheckPermissionRequest() *CheckPermissionRequest {
	return &CheckPermissionRequest{}
}
