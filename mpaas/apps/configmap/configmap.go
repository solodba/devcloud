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
