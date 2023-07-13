package user

import (
	"encoding/json"

	"github.com/solodba/devcloud/tree/main/mcenter/common/pb/meta"
)

// User结构体初始化函数
func NewUser(req *CreateUserRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	err := req.HashPassword()
	if err != nil {
		return nil, err
	}
	return &User{
		Meta: meta.NewMeta(),
		Spec: req,
	}, nil
}

// User结构体格式化输出方法
func (u *User) MarshalJSON() ([]byte, error) {
	data := struct {
		*meta.Meta
		*CreateUserRequest
	}{
		u.Meta,
		u.Spec,
	}
	return json.Marshal(data)
}
