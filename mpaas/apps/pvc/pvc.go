package pvc

import (
	"github.com/solodba/mcube/pb/meta"
)

// PVC结构体初始化函数
func NewPVC(req *CreatePVCRequest) *PVC {
	return &PVC{
		Meta: meta.NewMeta(),
		PVC:  req,
	}
}
