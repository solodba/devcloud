// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: apps/pvc/pb/rpc.proto

package pvc

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
	RPC_CreatePVC_FullMethodName = "/codehorse.mkube.pvc.RPC/CreatePVC"
	RPC_DeletePVC_FullMethodName = "/codehorse.mkube.pvc.RPC/DeletePVC"
	RPC_QueryPVC_FullMethodName  = "/codehorse.mkube.pvc.RPC/QueryPVC"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 创建PersistentVolumeClaim
	CreatePVC(ctx context.Context, in *CreatePVCRequest, opts ...grpc.CallOption) (*PVC, error)
	// 删除PersistentVolumeClaim
	DeletePVC(ctx context.Context, in *DeletePVCRequest, opts ...grpc.CallOption) (*PVC, error)
	// 查询PersistentVolumeClaim集合
	QueryPVC(ctx context.Context, in *QueryPVCRequest, opts ...grpc.CallOption) (*PVCSet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreatePVC(ctx context.Context, in *CreatePVCRequest, opts ...grpc.CallOption) (*PVC, error) {
	out := new(PVC)
	err := c.cc.Invoke(ctx, RPC_CreatePVC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeletePVC(ctx context.Context, in *DeletePVCRequest, opts ...grpc.CallOption) (*PVC, error) {
	out := new(PVC)
	err := c.cc.Invoke(ctx, RPC_DeletePVC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryPVC(ctx context.Context, in *QueryPVCRequest, opts ...grpc.CallOption) (*PVCSet, error) {
	out := new(PVCSet)
	err := c.cc.Invoke(ctx, RPC_QueryPVC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 创建PersistentVolumeClaim
	CreatePVC(context.Context, *CreatePVCRequest) (*PVC, error)
	// 删除PersistentVolumeClaim
	DeletePVC(context.Context, *DeletePVCRequest) (*PVC, error)
	// 查询PersistentVolumeClaim集合
	QueryPVC(context.Context, *QueryPVCRequest) (*PVCSet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreatePVC(context.Context, *CreatePVCRequest) (*PVC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePVC not implemented")
}
func (UnimplementedRPCServer) DeletePVC(context.Context, *DeletePVCRequest) (*PVC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePVC not implemented")
}
func (UnimplementedRPCServer) QueryPVC(context.Context, *QueryPVCRequest) (*PVCSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPVC not implemented")
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

func _RPC_CreatePVC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePVCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreatePVC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreatePVC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreatePVC(ctx, req.(*CreatePVCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeletePVC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePVCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeletePVC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeletePVC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeletePVC(ctx, req.(*DeletePVCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryPVC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPVCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryPVC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryPVC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryPVC(ctx, req.(*QueryPVCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codehorse.mkube.pvc.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePVC",
			Handler:    _RPC_CreatePVC_Handler,
		},
		{
			MethodName: "DeletePVC",
			Handler:    _RPC_DeletePVC_Handler,
		},
		{
			MethodName: "QueryPVC",
			Handler:    _RPC_QueryPVC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/pvc/pb/rpc.proto",
}
