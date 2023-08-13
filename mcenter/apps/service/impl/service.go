package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/service"
	"go.mongodb.org/mongo-driver/bson"
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
	filter := bson.M{}
	switch in.DescribeType {
	case service.DESCRIBE_BY_SERVICE_ID:
		filter["_id"] = in.DescribeValue
	case service.DESCRIBE_BY_SERVICE_CREDENTIAL_ID:
		filter["client_id"] = in.DescribeValue
	}
	serviceIns := service.NewDefaultService()
	err := i.col.FindOne(ctx, filter).Decode(serviceIns)
	if err != nil {
		return nil, err
	}
	return serviceIns, nil
}

// 通过服务名称找到service id
func (i *impl) QueryServiceIdByClientId(ctx context.Context, in *service.QueryServiceIdByClientIdRequest) (*service.Service, error) {
	serviceIns := service.NewDefaultService()
	filter := bson.M{"client_id": in.ClientId}
	err := i.col.FindOne(ctx, filter).Decode(serviceIns)
	if err != nil {
		return nil, err
	}
	return serviceIns, nil
}
