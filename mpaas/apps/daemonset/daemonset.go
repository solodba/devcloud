package daemonset

import "github.com/solodba/mcube/pb/meta"

// DaemonSet初始化函数
func NewDaemonSet(req *CreateDaemonSetRequest) *DaemonSet {
	return &DaemonSet{
		Meta:      meta.NewMeta(),
		DaemonSet: req,
	}
}

// DaemonSet默认初始化函数
func NewDefaultDaemonSet() *DaemonSet {
	return &DaemonSet{
		Meta:      meta.NewMeta(),
		DaemonSet: NewCreateDaemonSetRequest(),
	}
}

// DaemonSetList初始化函数
func NewDaemonSetList() *DaemonSetList {
	return &DaemonSetList{
		Items: make([]*DaemonSetListItem, 0),
	}
}

// DaemonSetList添加方法
func (d *DaemonSetList) AddItems(items ...*DaemonSetListItem) {
	d.Items = append(d.Items, items...)
}
