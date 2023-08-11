package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/service"
)

// 创建服务
func (i *impl) CreateService(ctx context.Context, in *service.CreateServiceRequest) (*service.Service, error) {
	serviceIns := service.NewService(in)
	_, err := i.col.InsertOne(ctx, serviceIns)
	if err != nil {
		return nil, err
	}
	return serviceIns, nil
}

// 查询服务
func (i *impl) QueryService(ctx context.Context, in *service.QueryServiceRequest) (*service.ServiceSet, error) {
	return nil, nil
}

// 查询服务详情
func (i *impl) DescribeService(ctx context.Context, in *service.DescribeServiceRequest) (*service.Service, error) {
	return nil, nil
}
