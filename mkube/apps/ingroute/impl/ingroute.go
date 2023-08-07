package impl

import (
	"context"

	"github.com/solodba/devcloud/mkube/apps/ingroute"
)

// 创建IngressRoute
func (i *impl) CreateIngressRoute(ctx context.Context, in *ingroute.CreateIngressRouteRequest) (*ingroute.IngressRoute, error) {
	return nil, nil
}

// 更新IngressRoute
func (i *impl) UpdateIngressRoute(ctx context.Context, in *ingroute.UpdateIngressRouteRequest) (*ingroute.IngressRoute, error) {
	return nil, nil
}

// 删除IngressRoute
func (i *impl) DeleteIngressRoute(ctx context.Context, in *ingroute.DeleteIngressRouteRequest) (*ingroute.CreateIngressRouteRequest, error) {
	return nil, nil
}

// 查询IngressRoute
func (i *impl) QueryIngressRoute(ctx context.Context, in *ingroute.QueryIngressRouteRequest) (*ingroute.IngressRouteSet, error) {
	return nil, nil
}

// 查询IngressRoute详情
func (i *impl) DescribeIngressRoute(ctx context.Context, in *ingroute.DescribeIngressRouteRequest) (*ingroute.CreateIngressRouteRequest, error) {
	return nil, nil
}

// 查询IngressRoute中间件列表
func (i *impl) QueryIngRouteMiddlewareList(ctx context.Context, in *ingroute.QueryIngRouteMwRequest) (*ingroute.MiddlewareList, error) {
	return nil, nil
}
