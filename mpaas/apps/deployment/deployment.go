package deployment

import "github.com/solodba/mcube/pb/meta"

// Deployment初始化函数
func NewDeployment(req *CreateDeploymentRequest) *Deployment {
	return &Deployment{
		Meta:       meta.NewMeta(),
		Deployment: req,
	}
}

// Deployment默认初始化函数
func NewDefaultDeployment() *Deployment {
	return &Deployment{
		Meta:       meta.NewMeta(),
		Deployment: NewCreateDeploymentRequest(),
	}
}

// DeploymentSet初始化函数
func NewDeploymentSet() *DeploymentSet {
	return &DeploymentSet{
		Items: make([]*DeploymentSetItem, 0),
	}
}

// DeploymentSet添加方法
func (d *DeploymentSet) AddItems(items ...*DeploymentSetItem) {
	d.Items = append(d.Items, items...)
}
