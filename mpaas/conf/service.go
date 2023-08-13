package conf

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/endpoint"
	"github.com/solodba/devcloud/mcenter/apps/service"
	"github.com/solodba/devcloud/mcenter/client/rpc"
	"github.com/solodba/mcube/logger"
)

// 获取service id
func (c *Config) GetServiceIdByClientId() (serviceId string, err error) {
	logger.L().Info().Msgf("%v", c.Mcenter)
	rpcClient := rpc.NewClient(c.Mcenter)
	req := service.NewQueryServiceIdByClientIdRequest(c.Mcenter.ClientID)
	serviceIns, err := rpcClient.NewServiceRPCClient().QueryServiceIdByClientId(context.Background(), req)
	if err != nil {
		return
	}
	serviceId = serviceIns.Meta.Id
	return
}

// endpointSet通过grpc写入到mcenter
func (c *Config) WriteEndpointSetToMcenter(es *endpoint.EndpointSet) error {
	rpcClient := rpc.NewClient(c.Mcenter)
	esReq := endpoint.NewRegistryEndpointRequest()
	for _, item := range es.Items {
		esCreateReq := endpoint.NewCreateEndpointRequest()
		esCreateReq.Auth = item.Spec.Auth
		esCreateReq.Method = item.Spec.Method
		esCreateReq.Perm = item.Spec.Perm
		esCreateReq.Operation = item.Spec.Operation
		esCreateReq.ServiceId = item.Spec.ServiceId
		esCreateReq.Path = item.Spec.Path
		esReq.AddItems(esCreateReq)
	}
	_, err := rpcClient.NewEndpointRPCClient().RegistryEndpoint(context.Background(), esReq)
	if err != nil {
		return err
	}
	return nil
}
