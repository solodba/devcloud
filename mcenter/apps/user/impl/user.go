package impl

import (
	"context"

	"github.com/solodba/devcloud/tree/main/mcenter/apps/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 创建用户
func (i *impl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	userIns, err := user.NewUser(in)
	if err != nil {
		return nil, err
	}
	_, err = i.col.InsertOne(ctx, userIns)
	if err != nil {
		return nil, err
	}
	return userIns, nil
}

// 删除用户
func (i *impl) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	req := user.NewDescribeUserRequest()
	req.DescribeValue = in.Username
	userIns, err := i.DescribeUser(ctx, req)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"username": userIns.Spec.Username}
	_, err = i.col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return userIns, nil
}

// 更新用户
func (i *impl) UpdateUser(ctx context.Context, in *user.UpdateUserRequest) (*user.User, error) {
	return nil, nil
}

// 查询用户
func (i *impl) Queryuser(ctx context.Context, in *user.QueryUserRequest) (*user.UserSet, error) {
	filter := bson.M{}
	if in.Keywords != "" {
		filter["username"] = bson.M{"$regex": in.Keywords, "$options": "im"}
	}
	opts := &options.FindOptions{}
	in.Page.ComputeOffSet()
	opts.SetSkip(in.Page.Offset)
	opts.SetLimit(in.Page.PageSize)
	cursor, err := i.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	userSet := user.NewUserSet()
	for cursor.Next(ctx) {
		userIns := user.NewDefaultUser()
		err = cursor.Decode(userIns)
		if err != nil {
			return nil, err
		}
		userIns.Desense()
		userSet.AddItems(userIns)
	}
	userSet.Total, err = i.col.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return userSet, nil
}

// 查询用户详情
func (i *impl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
	filter := bson.M{}
	switch in.DescribeType {
	case user.DESCRIBE_BY_USERNAME:
		filter["username"] = in.DescribeValue
	case user.DESCRIBE_BY_USER_ID:
		filter["_id"] = in.DescribeValue
	}
	userIns := user.NewDefaultUser()
	err := i.col.FindOne(ctx, filter).Decode(userIns)
	if err != nil {
		return nil, err
	}
	return userIns, nil
}
