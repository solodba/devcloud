package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/solodba/devcloud/mkube/apps/ingroute"
	"go.mongodb.org/mongo-driver/bson"
)

// 创建IngressRoute
func (i *impl) CreateIngressRoute(ctx context.Context, in *ingroute.CreateIngressRouteRequest) (*ingroute.IngressRoute, error) {
	k8sIngRoute := ingroute.NewK8sIngressRoute()
	k8sIngRoute.APIVersion = "traefik.io/v1alpha1"
	k8sIngRoute.Kind = "IngressRoute"
	k8sIngRoute.Metadata.Name = in.Name
	k8sIngRoute.Metadata.Namespace = in.Namespace
	k8sIngRouteLabels := make(map[string]string)
	for _, label := range in.Labels {
		k8sIngRouteLabels[label.Key] = label.Value
	}
	k8sIngRoute.Metadata.Labels = k8sIngRouteLabels
	k8sIngRoute.Spec = in.IngressRouteSpec
	result, err := json.Marshal(k8sIngRoute)
	if err != nil {
		return nil, fmt.Errorf("json marshal failed, err: %s", err.Error())
	}
	url := fmt.Sprintf(INGRESS_ROUTE_CREATE_URL, k8sIngRoute.Metadata.Namespace)
	_, err = i.clientSet.RESTClient().Get().AbsPath(url).Name(k8sIngRoute.Metadata.Name).DoRaw(ctx)
	if err == nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] ingress route is already exist", k8sIngRoute.Metadata.Namespace, k8sIngRoute.Metadata.Name)
	}
	_, err = i.clientSet.RESTClient().Post().AbsPath(url).Body(result).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] ingress route create error, err: %s", k8sIngRoute.Metadata.Namespace, k8sIngRoute.Metadata.Name, err.Error())
	}
	ingroute := ingroute.NewIngressRoute(in)
	// 新增ingRoute数据入库
	_, err = i.col.InsertOne(ctx, ingroute)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] insert into mongodb failed, err: %s", k8sIngRoute.Metadata.Namespace, k8sIngRoute.Metadata.Name, err.Error())
	}
	return ingroute, nil
}

// 更新IngressRoute
func (i *impl) UpdateIngressRoute(ctx context.Context, in *ingroute.UpdateIngressRouteRequest) (*ingroute.IngressRoute, error) {
	k8sIngRoute := ingroute.NewK8sIngressRoute()
	k8sIngRoute.APIVersion = "traefik.io/v1alpha1"
	k8sIngRoute.Kind = "IngressRoute"
	k8sIngRoute.Metadata.Name = in.IngressRoute.Name
	k8sIngRoute.Metadata.Namespace = in.IngressRoute.Namespace
	k8sIngRouteLabels := make(map[string]string)
	for _, label := range in.IngressRoute.Labels {
		k8sIngRouteLabels[label.Key] = label.Value
	}
	k8sIngRoute.Metadata.Labels = k8sIngRouteLabels
	k8sIngRoute.Spec = in.IngressRoute.IngressRouteSpec
	result, err := json.Marshal(k8sIngRoute)
	if err != nil {
		return nil, fmt.Errorf("json marshal failed, err: %s", err.Error())
	}
	url := fmt.Sprintf(INGRESS_ROUTE_UPDATE_URL, k8sIngRoute.Metadata.Namespace)
	raw, err := i.clientSet.RESTClient().Get().AbsPath(url).Name(k8sIngRoute.Metadata.Name).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] ingress route is not found", k8sIngRoute.Metadata.Namespace, k8sIngRoute.Metadata.Name)
	}
	ingressRoute := ingroute.NewK8sIngressRoute()
	err = json.Unmarshal(raw, ingressRoute)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed, err: %s", err.Error())
	}
	ingressRoute.Spec = k8sIngRoute.Spec
	result, err = json.Marshal(ingressRoute)
	if err != nil {
		return nil, fmt.Errorf("json marshal failed, err: %s", err.Error())
	}
	_, err = i.clientSet.RESTClient().Put().Name(ingressRoute.Metadata.Name).AbsPath(url).Body(result).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s name=%s] ingress route update error, err: %s", ingressRoute.Metadata.Namespace, ingressRoute.Metadata.Name, err.Error())
	}
	ing := ingroute.NewDefaultIngressRoute()
	err = i.col.FindOne(ctx, bson.M{"name": ingressRoute.Metadata.Name}).Decode(ing)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] not found in mongodb, err: %s", ingressRoute.Metadata.Namespace, ingressRoute.Metadata.Name, err.Error())
	}
	ing.Meta.UpdatedAt = time.Now().Unix()
	ing.IngressRoute = in.IngressRoute
	// 入库
	_, err = i.col.UpdateOne(ctx, bson.M{"name": ingressRoute.Metadata.Name}, bson.M{"$set": ing})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] update in mongodb, err: %s", ingressRoute.Metadata.Namespace, ingressRoute.Metadata.Name, err.Error())
	}
	return ing, nil
}

// 删除IngressRoute
func (i *impl) DeleteIngressRoute(ctx context.Context, in *ingroute.DeleteIngressRouteRequest) (*ingroute.CreateIngressRouteRequest, error) {
	req := ingroute.NewDescribeIngressRouteRequest()
	req.Namespace = in.Namespace
	req.Name = in.Name
	ingressRoute, err := i.DescribeIngressRoute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] ingress route not found", in.Namespace, in.Name)
	}
	url := fmt.Sprintf(INGRESS_ROUTE_DELETE_URL, ingressRoute.Namespace, ingressRoute.Name)
	_, err = i.clientSet.RESTClient().Delete().AbsPath(url).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] ingress delete fail", ingressRoute.Namespace, ingressRoute.Name)
	}
	// 从库中删除
	_, err = i.col.DeleteOne(ctx, bson.M{"name": ingressRoute.Name})
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s, name=%s] delete from mongodb fail", ingressRoute.Namespace, ingressRoute.Name)
	}
	return ingressRoute, nil
}

// 查询IngressRoute
func (i *impl) QueryIngressRoute(ctx context.Context, in *ingroute.QueryIngressRouteRequest) (*ingroute.IngressRouteSet, error) {
	url := fmt.Sprintf(INGRESS_ROUTE_LIST_URL, in.Namespace)
	k8sIngRouteList := ingroute.NewK8sIngressRouteList()
	ingRouteSet := ingroute.NewIngressRouteSet()
	raw, err := i.clientSet.RESTClient().Get().AbsPath(url).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s] get ingress route resource failed, err: %s", in.Namespace, err.Error())
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
		return nil, fmt.Errorf("[namespace=%s name=%s] get ingress route resource failed, err: %s", in.Namespace, in.Name, err.Error())
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
	ingRouteReq.IngressRouteSpec = k8sIngRoute.Spec
	return ingRouteReq, nil
}

// 查询IngressRoute中间件列表
func (i *impl) QueryIngRouteMiddlewareList(ctx context.Context, in *ingroute.QueryIngRouteMwRequest) (*ingroute.MiddlewareList, error) {
	url := fmt.Sprintf(MIDDLEWARE_LIST_URL, in.Namespace)
	raw, err := i.clientSet.RESTClient().Get().AbsPath(url).DoRaw(ctx)
	if err != nil {
		return nil, fmt.Errorf("[namespace=%s] get ingress route middleware list resource failed, err: %s", in.Namespace, err.Error())
	}
	k8sMwList := ingroute.NewK8sMiddlewareList()
	err = json.Unmarshal(raw, k8sMwList)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed, err: %s", err.Error())
	}
	mwList := ingroute.NewMiddlewareList()
	for _, item := range k8sMwList.Items {
		mwList.AddItems(item.Metadata.Name)
	}
	return mwList, nil
}
