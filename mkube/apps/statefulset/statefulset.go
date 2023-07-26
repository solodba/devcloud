package statefulset

import "github.com/solodba/mcube/pb/meta"

// StatefulSet初始化函数
func NewStatefulSet(req *CreateStatefulSetRequest) *StatefulSet {
	return &StatefulSet{
		Meta:        meta.NewMeta(),
		StatefulSet: req,
	}
}

// StatefulSet默认初始化函数
func NewDefaultStatefulSet() *StatefulSet {
	return &StatefulSet{
		Meta:        meta.NewMeta(),
		StatefulSet: NewCreateStatefulSetRequest(),
	}
}

// StatefulSetSet初始化函数
func NewStatefulSetSet() *StatefulSetSet {
	return &StatefulSetSet{
		Items: make([]*StatefulSetSetItem, 0),
	}
}

// StatefulSetSet添加方法
func (s *StatefulSetSet) AddItems(items ...*StatefulSetSetItem) {
	s.Items = append(s.Items, items...)
}
