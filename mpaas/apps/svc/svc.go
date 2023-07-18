package svc

import "github.com/solodba/mcube/pb/meta"

// Service结构体初始化函数
func NewService(req *CreateServiceRequest) *Service {
	return &Service{
		Meta:    meta.NewMeta(),
		Service: req,
	}
}

// Service结构体默认初始化函数
func NewDefaultService() *Service {
	return &Service{
		Meta:    meta.NewMeta(),
		Service: NewCreateServiceRequest(),
	}
}
