package service

import (
	"time"

	"github.com/rs/xid"
	"github.com/solodba/mcube/pb/meta"
)

// Service构造函数
func NewService(req *CreateServiceRequest) *Service {
	return &Service{
		Meta:       meta.NewMeta(),
		Spec:       req,
		Credential: NewCredential(),
	}
}

// Credential构造函数
func NewCredential() *Credential {
	return &Credential{
		CreateAt:     time.Now().Unix(),
		ClientId:     xid.New().String(),
		ClientSecret: xid.New().String(),
	}
}
