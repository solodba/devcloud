package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
	url := fmt.Sprintf(INGRESS_ROUTE_LIST_URL, in.Namespace)
	k8sIngRouteList := ingroute.NewK8sIngressRouteList()
	ingRouteSet := ingroute.NewIngressRouteSet()
	raw, err := i.clientSet.RESTClient().Get().AbsPath(url).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("get ingress route resource failed, err: %s", err.Error())
	}
	err = json.Unmarshal(raw, k8sIngRouteList)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed, err: %s", err.Error())
	}
	for i := range k8sIngRouteList.Items {
		if strings.Contains(k8sIngRouteList.Items[i].Metadata.Name, in.Keyword) {
			ingRouteSetItem := ingroute.NewIngressRouteSetItem()
			ingRouteSetItem.Name = k8sIngRouteList.Items[i].Metadata.Name
			ingRouteSetItem.Namespace = k8sIngRouteList.Items[i].Metadata.Namespace
			ingRouteSetItem.Age = k8sIngRouteList.Items[i].Metadata.CreationTimestamp.Unix()
			ingRouteSet.AddItems(ingRouteSetItem)
		}
	}
	ingRouteSet.Total = int64(len(ingRouteSet.Items))
	return ingRouteSet, nil
}

// 查询IngressRoute详情
func (i *impl) DescribeIngressRoute(ctx context.Context, in *ingroute.DescribeIngressRouteRequest) (*ingroute.CreateIngressRouteRequest, error) {
	ingRouteReq := ingroute.NewCreateIngressRouteRequest()
	k8sIngRoute := ingroute.NewK8sIngressRoute()
	url := fmt.Sprintf(INGRESS_ROUTE_DETAIL_URL, in.Namespace, in.Name)
	raw, err := i.clientSet.RESTClient().Get().AbsPath(url).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("get ingress route resource failed, err: %s", err.Error())
	}
	err = json.Unmarshal(raw, k8sIngRoute)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed, err: %s", err.Error())
	}
	ingRouteReq.Name = k8sIngRoute.Metadata.Name
	ingRouteReq.Namespace = k8sIngRoute.Metadata.Namespace
	ingRouteReqLabels := make([]*ingroute.ListMapItem, 0)
	for k, v := range k8sIngRoute.Metadata.Labels {
		ingRouteReqLabels = append(ingRouteReqLabels, &ingroute.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	ingRouteReq.Labels = ingRouteReqLabels
	ingRouteReq.IngressRouteSpec = &k8sIngRoute.Spec
	return ingRouteReq, nil
}

// 查询IngressRoute中间件列表
func (i *impl) QueryIngRouteMiddlewareList(ctx context.Context, in *ingroute.QueryIngRouteMwRequest) (*ingroute.MiddlewareList, error) {
	return nil, nil
}
