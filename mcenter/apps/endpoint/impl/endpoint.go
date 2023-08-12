package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/endpoint"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 注册Endpoint服务
func (i *impl) RegistryEndpoint(ctx context.Context, in *endpoint.RegistryEndpointRequest) (*endpoint.EndpointSet, error) {
	es := endpoint.NewEndpointSet()
	opt := &options.UpdateOptions{}
	opt.SetUpsert(true)
	for idx := range in.Items {
		createEp := in.Items[idx]
		ep := endpoint.NewEndpoint(createEp)
		// 插入到mongodb
		_, err := i.col.UpdateOne(ctx, bson.M{"_id": ep.Meta.Id}, bson.M{"$set": ep}, opt)
		if err != nil {
			return nil, err
		}
		es.AddItems(ep)
	}
	return es, nil
}

// 查询Endpoint服务
func (i *impl) QueryEndpoint(ctx context.Context, in *endpoint.QueryEndpointRequest) (*endpoint.EndpointSet, error) {
	es := endpoint.NewEndpointSet()
	opts := &options.FindOptions{}
	in.Page.ComputeOffSet()
	opts.SetLimit(in.Page.PageSize)
	opts.SetSkip(in.Page.Offset)
	filter := bson.M{}
	cursor, err := i.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		ep := endpoint.NewDefaultEndpoint()
		err = cursor.Decode(ep)
		if err != nil {
			return nil, err
		}
		es.AddItems(ep)
	}
	es.Total, err = i.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return es, nil
}
