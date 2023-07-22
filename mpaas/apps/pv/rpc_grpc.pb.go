// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: apps/pv/pb/rpc.proto

package pv

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
	RPC_CreatePV_FullMethodName = "/codehorse.mpaas.pv.RPC/CreatePV"
	RPC_DeletePV_FullMethodName = "/codehorse.mpaas.pv.RPC/DeletePV"
	RPC_QueryPV_FullMethodName  = "/codehorse.mpaas.pv.RPC/QueryPV"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 创建PersistentVolume
	CreatePV(ctx context.Context, in *CreatePVRequest, opts ...grpc.CallOption) (*PV, error)
	// 删除PersistentVolume
	DeletePV(ctx context.Context, in *DeletePVRequest, opts ...grpc.CallOption) (*PV, error)
	// 查询PersistentVolume集合
	QueryPV(ctx context.Context, in *QueryPVRequest, opts ...grpc.CallOption) (*PVSet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreatePV(ctx context.Context, in *CreatePVRequest, opts ...grpc.CallOption) (*PV, error) {
	out := new(PV)
	err := c.cc.Invoke(ctx, RPC_CreatePV_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeletePV(ctx context.Context, in *DeletePVRequest, opts ...grpc.CallOption) (*PV, error) {
	out := new(PV)
	err := c.cc.Invoke(ctx, RPC_DeletePV_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryPV(ctx context.Context, in *QueryPVRequest, opts ...grpc.CallOption) (*PVSet, error) {
	out := new(PVSet)
	err := c.cc.Invoke(ctx, RPC_QueryPV_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 创建PersistentVolume
	CreatePV(context.Context, *CreatePVRequest) (*PV, error)
	// 删除PersistentVolume
	DeletePV(context.Context, *DeletePVRequest) (*PV, error)
	// 查询PersistentVolume集合
	QueryPV(context.Context, *QueryPVRequest) (*PVSet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreatePV(context.Context, *CreatePVRequest) (*PV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePV not implemented")
}
func (UnimplementedRPCServer) DeletePV(context.Context, *DeletePVRequest) (*PV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePV not implemented")
}
func (UnimplementedRPCServer) QueryPV(context.Context, *QueryPVRequest) (*PVSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPV not implemented")
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

func _RPC_CreatePV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreatePV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreatePV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreatePV(ctx, req.(*CreatePVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeletePV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeletePV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeletePV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeletePV(ctx, req.(*DeletePVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryPV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryPV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryPV_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryPV(ctx, req.(*QueryPVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codehorse.mpaas.pv.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePV",
			Handler:    _RPC_CreatePV_Handler,
		},
		{
			MethodName: "DeletePV",
			Handler:    _RPC_DeletePV_Handler,
		},
		{
			MethodName: "QueryPV",
			Handler:    _RPC_QueryPV_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/pv/pb/rpc.proto",
}
