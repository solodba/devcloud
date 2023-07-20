package secret

import "github.com/solodba/mcube/pb/meta"

// Secret初始化函数
func NewSecret(req *CreateSecretRequest) *Secret {
	return &Secret{
		Meta:   meta.NewMeta(),
		Secret: req,
	}
}
