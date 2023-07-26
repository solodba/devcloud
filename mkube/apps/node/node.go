package node

// NodeSet初始化函数
func NewNodeSet() *NodeSet {
	return &NodeSet{
		Items: make([]*NodeSetItem, 0),
	}
}

// NodeSet添加方法
func (n *NodeSet) AddItems(items ...*NodeSetItem) {
	n.Items = append(n.Items, items...)
}
