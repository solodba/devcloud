package deployment

import "github.com/solodba/mcube/pb/meta"

// Deployment初始化函数
func NewDeployment(req *CreateDeploymentRequest) *Deployment {
	return &Deployment{
		Meta:       meta.NewMeta(),
		Deployment: req,
	}
}
