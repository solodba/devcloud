package sc

import "github.com/solodba/mcube/pb/meta"

// SC初始化函数
func NewSC(req *CreateSCRequest) *SC {
	return &SC{
		Meta: meta.NewMeta(),
		SC:   req,
	}
}

// SC默认初始化函数
func NewDefaultSC() *SC {
	return &SC{
		Meta: meta.NewMeta(),
		SC:   NewCreateSCRequest(),
	}
}

// SCSet初始化函数
func NewSCSet() *SCSet {
	return &SCSet{
		Items: make([]*SCSetItem, 0),
	}
}

// SCSet添加方法
func (s *SCSet) AddItems(items ...*SCSetItem) {
	s.Items = append(s.Items, items...)
}
