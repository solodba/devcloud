package user

import (
	context "context"

	"github.com/infraboard/mcube/validator"
	"github.com/solodba/devcloud/mcenter/apps/domain"
	"github.com/solodba/mcube/pb/page"
	"golang.org/x/crypto/bcrypt"
)

// 业务模块名称
const (
	AppName = "user"
)

// 定义业务接口
type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	// 更新用户
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// GRPC业务接口
	RPCServer
}

// CreateUserRequest结构体必要参数校验
func (req *CreateUserRequest) Validate() error {
	if req.Domain == "" {
		req.Domain = domain.DEFAULT_DOMAIN
	}
	return validator.V().Struct(req)
}

// CreateUserRequest结构体中密码加密
func (req *CreateUserRequest) HashPassword() error {
	hp, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return err
	}
	req.Password = string(hp)
	return nil
}

// CreateUserRequest初始化函数
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

// QueryUserRequest初始化函数
func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		Page: page.NewPageRequest(),
	}
}

// DescribeUserRequest结构体
func NewDescribeUserRequest() *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeType: DESCRIBE_BY_USERNAME,
	}
}

// DeleteUserRequest结构体添加校验方法
func (req *DeleteUserRequest) Validate() error {
	return validator.V().Struct(req)
}

// DeleteUserRequest结构体初始化方法
func NewDeleteUserRequest() *DeleteUserRequest {
	return &DeleteUserRequest{}
}

// UpdateUserRequest结构体添加校验方法
func (req *UpdateUserRequest) Validate() error {
	return validator.V().Struct(req)
}

// UpdateUserRequest结构体密码加密方法
func (req *UpdateUserRequest) HashPassword() error {
	hp, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return err
	}
	req.Password = string(hp)
	return nil
}

// UpdateUserRequest结构体初始化方法
func NewUpdateUserRequest() *UpdateUserRequest {
	return &UpdateUserRequest{}
}
