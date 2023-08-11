package auth

import (
	"context"

	"google.golang.org/grpc"
)

// grpc中间件认证结构体
type grpcAuther struct {
}

// grpcAuther构造函数
func NewGrpcAuther() *grpcAuther {
	return &grpcAuther{}
}

// grpc认证中间件方法
// type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
func (a *grpcAuther) AuthFunc(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	// 从context获取认证信息
	// md, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	return nil, nil
	// }
	return nil, nil
}
