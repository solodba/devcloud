package namespace

// NamespaceSet结构体初始化函数
func NewNamespaceSet() *NamespaceSet {
	return &NamespaceSet{
		Items: make([]*NamespaceSetItem, 0),
	}
}

// NamespaceSet结构体添加方法
func (n *NamespaceSet) AddItems(item ...*NamespaceSetItem) {
	n.Items = append(n.Items, item...)
}
