// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: apps/ingress/pb/rpc.proto

package ingress

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
	RPC_CreateIngress_FullMethodName   = "/codehorse.mpaas.ingress.RPC/CreateIngress"
	RPC_UpdateIngress_FullMethodName   = "/codehorse.mpaas.ingress.RPC/UpdateIngress"
	RPC_DeleteIngress_FullMethodName   = "/codehorse.mpaas.ingress.RPC/DeleteIngress"
	RPC_QueryIngress_FullMethodName    = "/codehorse.mpaas.ingress.RPC/QueryIngress"
	RPC_DescribeIngress_FullMethodName = "/codehorse.mpaas.ingress.RPC/DescribeIngress"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 创建Ingress
	CreateIngress(ctx context.Context, in *CreateIngressRequest, opts ...grpc.CallOption) (*Ingress, error)
	// 更新Ingress
	UpdateIngress(ctx context.Context, in *UpdateIngressRequest, opts ...grpc.CallOption) (*Ingress, error)
	// 删除Ingress
	DeleteIngress(ctx context.Context, in *DeleteIngressRequest, opts ...grpc.CallOption) (*Ingress, error)
	// 查询Ingress
	QueryIngress(ctx context.Context, in *QueryIngressRequest, opts ...grpc.CallOption) (*IngressSet, error)
	// 查询Ingress详情
	DescribeIngress(ctx context.Context, in *DescribeIngressRequest, opts ...grpc.CallOption) (*Ingress, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreateIngress(ctx context.Context, in *CreateIngressRequest, opts ...grpc.CallOption) (*Ingress, error) {
	out := new(Ingress)
	err := c.cc.Invoke(ctx, RPC_CreateIngress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) UpdateIngress(ctx context.Context, in *UpdateIngressRequest, opts ...grpc.CallOption) (*Ingress, error) {
	out := new(Ingress)
	err := c.cc.Invoke(ctx, RPC_UpdateIngress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteIngress(ctx context.Context, in *DeleteIngressRequest, opts ...grpc.CallOption) (*Ingress, error) {
	out := new(Ingress)
	err := c.cc.Invoke(ctx, RPC_DeleteIngress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryIngress(ctx context.Context, in *QueryIngressRequest, opts ...grpc.CallOption) (*IngressSet, error) {
	out := new(IngressSet)
	err := c.cc.Invoke(ctx, RPC_QueryIngress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeIngress(ctx context.Context, in *DescribeIngressRequest, opts ...grpc.CallOption) (*Ingress, error) {
	out := new(Ingress)
	err := c.cc.Invoke(ctx, RPC_DescribeIngress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 创建Ingress
	CreateIngress(context.Context, *CreateIngressRequest) (*Ingress, error)
	// 更新Ingress
	UpdateIngress(context.Context, *UpdateIngressRequest) (*Ingress, error)
	// 删除Ingress
	DeleteIngress(context.Context, *DeleteIngressRequest) (*Ingress, error)
	// 查询Ingress
	QueryIngress(context.Context, *QueryIngressRequest) (*IngressSet, error)
	// 查询Ingress详情
	DescribeIngress(context.Context, *DescribeIngressRequest) (*Ingress, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreateIngress(context.Context, *CreateIngressRequest) (*Ingress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIngress not implemented")
}
func (UnimplementedRPCServer) UpdateIngress(context.Context, *UpdateIngressRequest) (*Ingress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIngress not implemented")
}
func (UnimplementedRPCServer) DeleteIngress(context.Context, *DeleteIngressRequest) (*Ingress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIngress not implemented")
}
func (UnimplementedRPCServer) QueryIngress(context.Context, *QueryIngressRequest) (*IngressSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryIngress not implemented")
}
func (UnimplementedRPCServer) DescribeIngress(context.Context, *DescribeIngressRequest) (*Ingress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeIngress not implemented")
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

func _RPC_CreateIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIngressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateIngress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateIngress(ctx, req.(*CreateIngressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_UpdateIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIngressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).UpdateIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_UpdateIngress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).UpdateIngress(ctx, req.(*UpdateIngressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIngressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteIngress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteIngress(ctx, req.(*DeleteIngressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIngressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryIngress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryIngress(ctx, req.(*QueryIngressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeIngressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribeIngress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeIngress(ctx, req.(*DescribeIngressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codehorse.mpaas.ingress.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIngress",
			Handler:    _RPC_CreateIngress_Handler,
		},
		{
			MethodName: "UpdateIngress",
			Handler:    _RPC_UpdateIngress_Handler,
		},
		{
			MethodName: "DeleteIngress",
			Handler:    _RPC_DeleteIngress_Handler,
		},
		{
			MethodName: "QueryIngress",
			Handler:    _RPC_QueryIngress_Handler,
		},
		{
			MethodName: "DescribeIngress",
			Handler:    _RPC_DescribeIngress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/ingress/pb/rpc.proto",
}