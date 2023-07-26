package pod

import (
	"fmt"

	"github.com/solodba/mcube/pb/meta"
	"github.com/solodba/mcube/validator"
)

// 定义常量
const (
	IMAGE_PULL_POLICY_IFNOTPRESENT = "IfNotPresent"
	RESTART_POLICY_ALWAYS          = "Always"
)

// Base结构体必要参数校验
func (req *Base) Validate() error {
	return validator.V().Struct(req)
}

// Container结构体校验
func (req *Container) Validate() error {
	return validator.V().Struct(req)
}

// Pod结构必要参数校验
func (req *Pod) Validate() error {
	if err := req.Base.Validate(); err != nil {
		return fmt.Errorf("请定义Pod的名字!")
	}
	if len(req.Containers) == 0 {
		return fmt.Errorf("请定义Pod的容器信息!")
	}
	if len(req.InitContainers) > 0 {
		for _, container := range req.InitContainers {
			if err := container.Validate(); err != nil {
				return fmt.Errorf("InitContainers中发现没有定义名称或者镜像的容器!")
			}
			if container.ImagePullPolicy == "" {
				container.ImagePullPolicy = IMAGE_PULL_POLICY_IFNOTPRESENT
			}
		}
	}
	if len(req.Containers) > 0 {
		for _, container := range req.Containers {
			if err := container.Validate(); err != nil {
				return fmt.Errorf("Containers中发现没有定义名称或者镜像的容器!")
			}
			if container.ImagePullPolicy == "" {
				container.ImagePullPolicy = IMAGE_PULL_POLICY_IFNOTPRESENT
			}
		}
	}
	if req.Base.RestartPolicy == "" {
		req.Base.RestartPolicy = RESTART_POLICY_ALWAYS
	}
	return nil
}

// Pod默认初始化函数
func NewDefaultPod() *Pod {
	return &Pod{
		Meta:           meta.NewMeta(),
		Base:           NewBase(),
		Volumes:        make([]*Volume, 0),
		NetWorking:     NewNetWorking(),
		InitContainers: make([]*Container, 0),
		Containers:     make([]*Container, 0),
		Tolerations:    make([]*Tolerations, 0),
		NodeScheduling: NewNodeScheduling(),
	}
}

// Pod初始化函数
func NewPod(in *CreatePodRequest) (*Pod, error) {
	if err := in.Pod.Validate(); err != nil {
		return nil, err
	}
	return &Pod{
		Meta:           meta.NewMeta(),
		Base:           in.Pod.Base,
		Volumes:        in.Pod.Volumes,
		NetWorking:     in.Pod.NetWorking,
		InitContainers: in.Pod.InitContainers,
		Containers:     in.Pod.Containers,
		Tolerations:    in.Pod.Tolerations,
		NodeScheduling: in.Pod.NodeScheduling,
	}, nil
}

// Base初始化函数
func NewBase() *Base {
	return &Base{
		Name:          "test",
		Namespace:     "test",
		Labels:        make([]*ListMapItem, 0),
		RestartPolicy: RESTART_POLICY_ALWAYS,
	}
}

// NetWorking初始化函数
func NewNetWorking() *NetWorking {
	return &NetWorking{}
}

// NodeScheduling初始化函数
func NewNodeScheduling() *NodeScheduling {
	return &NodeScheduling{}
}

// PodSet初始化参数
func NewPodSet() *PodSet {
	return &PodSet{
		PodListItem: make([]*PodListItem, 0),
	}
}

// PodSet添加方法
func (p *PodSet) AddItems(items ...*PodListItem) {
	p.PodListItem = append(p.PodListItem, items...)
}
