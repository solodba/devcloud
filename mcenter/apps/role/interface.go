package role

import context "context"

// 模块名称
const (
	AppName = "role"
)

// Role相关业务接口
type Service interface {
	// 创建role
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	// 嵌套role grpc接口
	RPCServer
}

// CreateRoleRequest构造函数
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Feature: make([]*Feature, 0),
	}
}

// CreateRoleRequest添加方法
func (req *CreateRoleRequest) AddItems(items ...*Feature) {
	req.Feature = append(req.Feature, items...)
}

// Feature构造方法
func NewFeature() *Feature {
	return &Feature{}
}
