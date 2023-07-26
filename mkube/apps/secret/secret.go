package secret

import "github.com/solodba/mcube/pb/meta"

// Secret初始化函数
func NewSecret(req *CreateSecretRequest) *Secret {
	return &Secret{
		Meta:   meta.NewMeta(),
		Secret: req,
	}
}

// Secret初始化默认函数
func NewDefaultSecret() *Secret {
	return &Secret{
		Meta:   meta.NewMeta(),
		Secret: NewCreateSecretRequest(),
	}
}

// SecretSet初始化函数
func NewSecretSet() *SecretSet {
	return &SecretSet{
		Items: make([]*SecretSetItem, 0),
	}
}

// SecretSet添加方法
func (s *SecretSet) AddItems(items ...*SecretSetItem) {
	s.Items = append(s.Items, items...)
}
