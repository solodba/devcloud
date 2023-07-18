package svc

import "github.com/solodba/mcube/pb/meta"

// Service结构体初始化函数
func NewService(req *CreateServiceRequest) *Service {
	return &Service{
		Meta:    meta.NewMeta(),
		Service: req,
	}
}
