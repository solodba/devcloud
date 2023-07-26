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

// PVC结构体默认初始化函数
func NewDefaultPVC() *PVC {
	return &PVC{
		Meta: meta.NewMeta(),
		PVC:  NewCreatePVCRequest(),
	}
}

// PVCSet初始化函数
func NewPVCSet() *PVCSet {
	return &PVCSet{
		Items: make([]*PVCSetItem, 0),
	}
}

// PVCSet添加方法
func (p *PVCSet) AddItems(items ...*PVCSetItem) {
	p.Items = append(p.Items, items...)
}
