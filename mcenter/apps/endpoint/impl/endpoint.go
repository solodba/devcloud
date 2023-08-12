package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/endpoint"
)

// 注册Endpoint服务
func (i *impl) RegistryEndpoint(ctx context.Context, in *endpoint.RegistryEndpointRequest) (*endpoint.EndpointSet, error) {
	return nil, nil
}

// 查询Endpoint服务
func (i *impl) QueryEndpoint(ctx context.Context, in *endpoint.QueryEndpointRequest) (*endpoint.EndpointSet, error) {
	return nil, nil
}
