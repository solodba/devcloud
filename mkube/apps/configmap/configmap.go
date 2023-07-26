package configmap

import "github.com/solodba/mcube/pb/meta"

// ConfigMap初始化函数
func NewConfigMap(req *CreateConfigMapRequest) *ConfigMap {
	return &ConfigMap{
		Meta:      meta.NewMeta(),
		ConfigMap: req,
	}
}

// ConfigMap默认初始化函数
func NewDefaultConfigMap() *ConfigMap {
	return &ConfigMap{
		Meta:      meta.NewMeta(),
		ConfigMap: NewCreateConfigMapRequest(),
	}
}

// ConfigMapSet初始化函数
func NewConfigMapSet() *ConfigMapSet {
	return &ConfigMapSet{
		Items: make([]*ConfigMapSetItem, 0),
	}
}

// ConfigMapSet添加方法
func (c *ConfigMapSet) AddItems(items ...*ConfigMapSetItem) {
	c.Items = append(c.Items, items...)
}
