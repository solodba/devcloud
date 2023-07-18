// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: apps/svc/pb/rpc.proto

package svc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPC_CreateService_FullMethodName   = "/codehorse.mpaas.svc.RPC/CreateService"
	RPC_DeleteService_FullMethodName   = "/codehorse.mpaas.svc.RPC/DeleteService"
	RPC_UpdateService_FullMethodName   = "/codehorse.mpaas.svc.RPC/UpdateService"
	RPC_QueryService_FullMethodName    = "/codehorse.mpaas.svc.RPC/QueryService"
	RPC_DescribeService_FullMethodName = "/codehorse.mpaas.svc.RPC/DescribeService"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 创建Service
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Service, error)
	// 删除Service
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*Service, error)
	// 更新Service
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*Service, error)
	// 查询Service
	QueryService(ctx context.Context, in *QueryServiceRequest, opts ...grpc.CallOption) (*ServiceSet, error)
	// 查询Service详情
	DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, RPC_CreateService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, RPC_DeleteService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, RPC_UpdateService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryService(ctx context.Context, in *QueryServiceRequest, opts ...grpc.CallOption) (*ServiceSet, error) {
	out := new(ServiceSet)
	err := c.cc.Invoke(ctx, RPC_QueryService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, RPC_DescribeService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 创建Service
	CreateService(context.Context, *CreateServiceRequest) (*Service, error)
	// 删除Service
	DeleteService(context.Context, *DeleteServiceRequest) (*Service, error)
	// 更新Service
	UpdateService(context.Context, *UpdateServiceRequest) (*Service, error)
	// 查询Service
	QueryService(context.Context, *QueryServiceRequest) (*ServiceSet, error)
	// 查询Service详情
	DescribeService(context.Context, *DescribeServiceRequest) (*Service, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreateService(context.Context, *CreateServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedRPCServer) DeleteService(context.Context, *DeleteServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedRPCServer) UpdateService(context.Context, *UpdateServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedRPCServer) QueryService(context.Context, *QueryServiceRequest) (*ServiceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryService not implemented")
}
func (UnimplementedRPCServer) DescribeService(context.Context, *DescribeServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_UpdateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryService(ctx, req.(*QueryServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribeService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeService(ctx, req.(*DescribeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codehorse.mpaas.svc.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _RPC_CreateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _RPC_DeleteService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _RPC_UpdateService_Handler,
		},
		{
			MethodName: "QueryService",
			Handler:    _RPC_QueryService_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _RPC_DescribeService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/svc/pb/rpc.proto",
}
