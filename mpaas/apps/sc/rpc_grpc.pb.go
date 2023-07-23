// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: apps/sc/pb/rpc.proto

package sc

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
	RPC_CreateSC_FullMethodName = "/codehorse.mpaas.sc.RPC/CreateSC"
	RPC_DeleteSC_FullMethodName = "/codehorse.mpaas.sc.RPC/DeleteSC"
	RPC_QuerySC_FullMethodName  = "/codehorse.mpaas.sc.RPC/QuerySC"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 创建StorageClass
	CreateSC(ctx context.Context, in *CreateSCRequest, opts ...grpc.CallOption) (*SC, error)
	// 删除StorageClass
	DeleteSC(ctx context.Context, in *DeleteSCRequest, opts ...grpc.CallOption) (*SC, error)
	// 查询StorageClass集合
	QuerySC(ctx context.Context, in *QuerySCRequest, opts ...grpc.CallOption) (*SCSet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreateSC(ctx context.Context, in *CreateSCRequest, opts ...grpc.CallOption) (*SC, error) {
	out := new(SC)
	err := c.cc.Invoke(ctx, RPC_CreateSC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteSC(ctx context.Context, in *DeleteSCRequest, opts ...grpc.CallOption) (*SC, error) {
	out := new(SC)
	err := c.cc.Invoke(ctx, RPC_DeleteSC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QuerySC(ctx context.Context, in *QuerySCRequest, opts ...grpc.CallOption) (*SCSet, error) {
	out := new(SCSet)
	err := c.cc.Invoke(ctx, RPC_QuerySC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 创建StorageClass
	CreateSC(context.Context, *CreateSCRequest) (*SC, error)
	// 删除StorageClass
	DeleteSC(context.Context, *DeleteSCRequest) (*SC, error)
	// 查询StorageClass集合
	QuerySC(context.Context, *QuerySCRequest) (*SCSet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreateSC(context.Context, *CreateSCRequest) (*SC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSC not implemented")
}
func (UnimplementedRPCServer) DeleteSC(context.Context, *DeleteSCRequest) (*SC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSC not implemented")
}
func (UnimplementedRPCServer) QuerySC(context.Context, *QuerySCRequest) (*SCSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySC not implemented")
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

func _RPC_CreateSC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateSC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateSC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateSC(ctx, req.(*CreateSCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteSC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteSC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteSC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteSC(ctx, req.(*DeleteSCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QuerySC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QuerySC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QuerySC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QuerySC(ctx, req.(*QuerySCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codehorse.mpaas.sc.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSC",
			Handler:    _RPC_CreateSC_Handler,
		},
		{
			MethodName: "DeleteSC",
			Handler:    _RPC_DeleteSC_Handler,
		},
		{
			MethodName: "QuerySC",
			Handler:    _RPC_QuerySC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/sc/pb/rpc.proto",
}