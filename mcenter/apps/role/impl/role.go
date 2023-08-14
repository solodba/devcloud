package impl

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 创建role
func (i *impl) CreateRole(ctx context.Context, in *role.CreateRoleRequest) (*role.Role, error) {
	roleIns := role.NewRole(in)
	_, err := i.col.InsertOne(ctx, roleIns)
	if err != nil {
		return nil, err
	}
	return roleIns, nil
}

// 查询role
func (i *impl) QueryRole(ctx context.Context, in *role.QueryRoleRequest) (*role.RoleSet, error) {
	filter := bson.M{}
	if len(in.RoleIds) > 0 {
		filter["_id"] = bson.M{"$in": in.RoleIds}
	}
	opts := &options.FindOptions{}
	opts.SetLimit(in.Page.PageSize)
	in.Page.ComputeOffSet()
	opts.SetSkip(in.Page.Offset)
	cursor, err := i.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	roleSetIns := role.NewRoleSet()
	for cursor.Next(ctx) {
		roleIns := role.NewDefaultRole()
		err = cursor.Decode(roleIns)
		if err != nil {
			return nil, err
		}
		roleSetIns.AddItems(roleIns)
	}
	roleSetIns.Total, err = i.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return roleSetIns, nil
}
