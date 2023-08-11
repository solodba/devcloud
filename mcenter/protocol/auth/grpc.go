package auth

import (
	"context"

	"github.com/solodba/devcloud/mcenter/apps/service"
	"github.com/solodba/mcube/apps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// grpc中间件认证结构体
type grpcAuther struct {
	svc service.ServiceManager
}

// grpcAuther构造函数
func NewGrpcAuther() *grpcAuther {
	return &grpcAuther{
		svc: apps.GetInternalApp(service.AppName).(service.ServiceManager),
	}
}

// 通过解析metadata.MD获取ClientID和ClientSecret
func (a *grpcAuther) GetClientIDAndClientSecretFromMetadataMD(md metadata.MD) (clientID, clientSecret string) {
	clientIDList := md.Get(service.SERVICE_CLIENT_ID)
	clientSecretList := md.Get(service.SERVICE_CLIENT_SECRET)
	if len(clientIDList) > 0 {
		clientID = clientIDList[0]
	}
	if len(clientSecretList) > 0 {
		clientSecret = clientSecretList[0]
	}
	return
}

// grpc认证中间件方法
// type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
func (a *grpcAuther) AuthFunc(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	// 从context获取认证信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Aborted, "grpc访问需要认证!")
	}
	// 通过解析metadata.MD获取ClientID和ClientSecret
	clientID, clientSecret := a.GetClientIDAndClientSecretFromMetadataMD(md)
	if clientID == "" || clientSecret == "" {
		return nil, status.Error(codes.Aborted, "clientID或者clientSecret不能为空!")
	}
	// 通过内部服务调用查询是否存在clientID和clientSecret
	serviceReq := service.NewDescribeServiceRequest()
	serviceReq.DescribeType = service.DESCRIBE_BY_SERVICE_CREDENTIAL_ID
	serviceReq.DescribeValue = clientID
	serviceIns, err := a.svc.DescribeService(ctx, serviceReq)
	if err != nil {
		return nil, status.Error(codes.Aborted, "grpc认证失败!")
	}
	if serviceIns.Credential.ClientSecret != clientSecret {
		return nil, status.Error(codes.Aborted, "grpc认证失败!")
	}
	resp, err = handler(ctx, req)
	return
}
