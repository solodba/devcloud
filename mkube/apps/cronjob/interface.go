package cronjob

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/solodba/devcloud/mkube/apps/pod"
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

// UpdateCronJobRequest初始化函数
func NewUpdateCronJobRequest() *UpdateCronJobRequest {
	return &UpdateCronJobRequest{
		CronJob: NewCreateCronJobRequest(),
	}
}

// QueryCronJobRequest初始化函数
func NewQueryCronJobRequest() *QueryCronJobRequest {
	return &QueryCronJobRequest{}
}

// DescribeCronJobRequest初始化函数
func NewDescribeCronJobRequest() *DescribeCronJobRequest {
	return &DescribeCronJobRequest{}
}

// DeleteCronJobRequest初始化函数
func NewDeleteCronJobRequest() *DeleteCronJobRequest {
	return &DeleteCronJobRequest{}
}

// 从restful解析参数初始化DeleteCronJobRequest
func NewDeleteCronJobRequestFromRestful(r *restful.Request) *DeleteCronJobRequest {
	return &DeleteCronJobRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.PathParameter("name"),
	}
}

// 从restful解析参数初始化QueryCronJobRequest
func NewQueryCronJobRequestFromRestful(r *restful.Request) *QueryCronJobRequest {
	return &QueryCronJobRequest{
		Namespace: r.PathParameter("namespace"),
		Keyword:   r.QueryParameter("keyword"),
	}
}

// 从restful解析参数初始化DescribeCronJobRequest
func NewDescribeCronJobRequestFromRestful(r *restful.Request) *DescribeCronJobRequest {
	return &DescribeCronJobRequest{
		Namespace: r.PathParameter("namespace"),
		Name:      r.QueryParameter("name"),
	}
}
