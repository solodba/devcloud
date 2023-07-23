package sc

import "github.com/solodba/mcube/pb/meta"

// SC初始化函数
func NewSC(req *CreateSCRequest) *SC {
	return &SC{
		Meta: meta.NewMeta(),
		SC:   req,
	}
}
