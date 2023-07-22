package pv

import (
	"github.com/solodba/mcube/pb/meta"
)

// PV结构体初始化函数
func NewPV(req *CreatePVRequest) *PV {
	return &PV{
		Meta: meta.NewMeta(),
		PV:   req,
	}
}
