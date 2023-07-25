package cronjob

import (
	"github.com/solodba/devcloud/mpaas/apps/pod"
	"github.com/solodba/mcube/pb/meta"
)

// 模块名称
const (
	AppName = "cronjob"
)

// CronJob相关功能接口
type Service interface {
	// 嵌套CronJob GRPC接口
	RPCServer
}

// CreateCronJobRequest结构体初始化函数
func NewCreateCronJobRequest() *CreateCronJobRequest {
	return &CreateCronJobRequest{
		Base: &CronJobBase{
			Labels: make([]*ListMapItem, 0),
		},
		Template: &pod.Pod{
			Meta: meta.NewMeta(),
			Base: &pod.Base{
				Labels: make([]*pod.ListMapItem, 0),
			},
			Volumes: make([]*pod.Volume, 0),
			NetWorking: &pod.NetWorking{
				DnsConfig:   &pod.DnsConfig{},
				HostAliases: make([]*pod.ListMapItem, 0),
			},
			InitContainers: make([]*pod.Container, 0),
			Containers:     make([]*pod.Container, 0),
			Tolerations:    make([]*pod.Tolerations, 0),
		},
	}
}
