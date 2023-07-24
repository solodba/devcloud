package ingress

import "github.com/solodba/mcube/pb/meta"

// Ingress初始化函数
func NewIngress(req *CreateIngressRequest) *Ingress {
	return &Ingress{
		Meta:    meta.NewMeta(),
		Ingress: req,
	}
}

// Ingress默认初始化函数
func NewDefaultIngress() *Ingress {
	return &Ingress{
		Meta:    meta.NewMeta(),
		Ingress: NewCreateIngressRequest(),
	}
}
