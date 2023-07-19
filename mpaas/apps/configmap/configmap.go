package configmap

import "github.com/solodba/mcube/pb/meta"

// ConfigMap初始化函数
func NewConfigMap(req *CreateConfigMapRequest) *ConfigMap {
	return &ConfigMap{
		Meta:      meta.NewMeta(),
		ConfigMap: req,
	}
}
