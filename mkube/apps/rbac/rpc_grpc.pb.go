// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: apps/rbac/pb/rpc.proto

package rbac

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
	RPC_CreateOrUpdateServiceAccount_FullMethodName = "/codehorse.mkube.rbac.RPC/CreateOrUpdateServiceAccount"
	RPC_DeleteServiceAccount_FullMethodName         = "/codehorse.mkube.rbac.RPC/DeleteServiceAccount"
	RPC_GetServiceAccountList_FullMethodName        = "/codehorse.mkube.rbac.RPC/GetServiceAccountList"
	RPC_CreateOrUpdateClusterRole_FullMethodName    = "/codehorse.mkube.rbac.RPC/CreateOrUpdateClusterRole"
	RPC_DeleteClusterRole_FullMethodName            = "/codehorse.mkube.rbac.RPC/DeleteClusterRole"
	RPC_GetClusterRoleList_FullMethodName           = "/codehorse.mkube.rbac.RPC/GetClusterRoleList"
	RPC_GetClusterRoleDetail_FullMethodName         = "/codehorse.mkube.rbac.RPC/GetClusterRoleDetail"
	RPC_CreateOrUpdateRole_FullMethodName           = "/codehorse.mkube.rbac.RPC/CreateOrUpdateRole"
	RPC_DeleteRole_FullMethodName                   = "/codehorse.mkube.rbac.RPC/DeleteRole"
	RPC_GetRoleList_FullMethodName                  = "/codehorse.mkube.rbac.RPC/GetRoleList"
	RPC_GetRoleDetail_FullMethodName                = "/codehorse.mkube.rbac.RPC/GetRoleDetail"
	RPC_CreateOrUpdateRb_FullMethodName             = "/codehorse.mkube.rbac.RPC/CreateOrUpdateRb"
	RPC_DeleteRb_FullMethodName                     = "/codehorse.mkube.rbac.RPC/DeleteRb"
	RPC_GetRbList_FullMethodName                    = "/codehorse.mkube.rbac.RPC/GetRbList"
	RPC_GetRbDetail_FullMethodName                  = "/codehorse.mkube.rbac.RPC/GetRbDetail"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 创建或者更新ServcieAccount
	CreateOrUpdateServiceAccount(ctx context.Context, in *ServiceAccount, opts ...grpc.CallOption) (*Response, error)
	// 删除ServiceAccount
	DeleteServiceAccount(ctx context.Context, in *DeleteServiceAccountRequest, opts ...grpc.CallOption) (*Response, error)
	// 查询ServiceAccount集合
	GetServiceAccountList(ctx context.Context, in *GetServiceAccountListRequest, opts ...grpc.CallOption) (*ServiceAccountSet, error)
	// 创建或者更新ClusterRole
	CreateOrUpdateClusterRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Response, error)
	// 删除ClusterRole
	DeleteClusterRole(ctx context.Context, in *DeleteClusterRoleRequest, opts ...grpc.CallOption) (*Response, error)
	// 查询ClusterRole集合
	GetClusterRoleList(ctx context.Context, in *GetClusterRoleListRequest, opts ...grpc.CallOption) (*ClusterRoleSet, error)
	// 查询ClusterRole详情
	GetClusterRoleDetail(ctx context.Context, in *GetClusterRoleDetailRequest, opts ...grpc.CallOption) (*Role, error)
	// 创建Role
	CreateOrUpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Response, error)
	// 删除Role
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*Response, error)
	// 查询Role集合
	GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*RoleSet, error)
	// 查询Role详情
	GetRoleDetail(ctx context.Context, in *GetRoleDetailRequest, opts ...grpc.CallOption) (*Role, error)
	// 创建RoleBinding
	CreateOrUpdateRb(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*Response, error)
	// 删除RoleBinding
	DeleteRb(ctx context.Context, in *DeleteRbRequest, opts ...grpc.CallOption) (*Response, error)
	// 查看RoleBinding列表
	GetRbList(ctx context.Context, in *GetRbListRequest, opts ...grpc.CallOption) (*RoleBindingSet, error)
	// 查询RoleBinding详情
	GetRbDetail(ctx context.Context, in *GetRbDetailRequest, opts ...grpc.CallOption) (*RoleBinding, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreateOrUpdateServiceAccount(ctx context.Context, in *ServiceAccount, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_CreateOrUpdateServiceAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteServiceAccount(ctx context.Context, in *DeleteServiceAccountRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_DeleteServiceAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetServiceAccountList(ctx context.Context, in *GetServiceAccountListRequest, opts ...grpc.CallOption) (*ServiceAccountSet, error) {
	out := new(ServiceAccountSet)
	err := c.cc.Invoke(ctx, RPC_GetServiceAccountList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CreateOrUpdateClusterRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_CreateOrUpdateClusterRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteClusterRole(ctx context.Context, in *DeleteClusterRoleRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_DeleteClusterRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetClusterRoleList(ctx context.Context, in *GetClusterRoleListRequest, opts ...grpc.CallOption) (*ClusterRoleSet, error) {
	out := new(ClusterRoleSet)
	err := c.cc.Invoke(ctx, RPC_GetClusterRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetClusterRoleDetail(ctx context.Context, in *GetClusterRoleDetailRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, RPC_GetClusterRoleDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CreateOrUpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_CreateOrUpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*RoleSet, error) {
	out := new(RoleSet)
	err := c.cc.Invoke(ctx, RPC_GetRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetRoleDetail(ctx context.Context, in *GetRoleDetailRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, RPC_GetRoleDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CreateOrUpdateRb(ctx context.Context, in *RoleBinding, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_CreateOrUpdateRb_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteRb(ctx context.Context, in *DeleteRbRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, RPC_DeleteRb_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetRbList(ctx context.Context, in *GetRbListRequest, opts ...grpc.CallOption) (*RoleBindingSet, error) {
	out := new(RoleBindingSet)
	err := c.cc.Invoke(ctx, RPC_GetRbList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetRbDetail(ctx context.Context, in *GetRbDetailRequest, opts ...grpc.CallOption) (*RoleBinding, error) {
	out := new(RoleBinding)
	err := c.cc.Invoke(ctx, RPC_GetRbDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 创建或者更新ServcieAccount
	CreateOrUpdateServiceAccount(context.Context, *ServiceAccount) (*Response, error)
	// 删除ServiceAccount
	DeleteServiceAccount(context.Context, *DeleteServiceAccountRequest) (*Response, error)
	// 查询ServiceAccount集合
	GetServiceAccountList(context.Context, *GetServiceAccountListRequest) (*ServiceAccountSet, error)
	// 创建或者更新ClusterRole
	CreateOrUpdateClusterRole(context.Context, *Role) (*Response, error)
	// 删除ClusterRole
	DeleteClusterRole(context.Context, *DeleteClusterRoleRequest) (*Response, error)
	// 查询ClusterRole集合
	GetClusterRoleList(context.Context, *GetClusterRoleListRequest) (*ClusterRoleSet, error)
	// 查询ClusterRole详情
	GetClusterRoleDetail(context.Context, *GetClusterRoleDetailRequest) (*Role, error)
	// 创建Role
	CreateOrUpdateRole(context.Context, *Role) (*Response, error)
	// 删除Role
	DeleteRole(context.Context, *DeleteRoleRequest) (*Response, error)
	// 查询Role集合
	GetRoleList(context.Context, *GetRoleListRequest) (*RoleSet, error)
	// 查询Role详情
	GetRoleDetail(context.Context, *GetRoleDetailRequest) (*Role, error)
	// 创建RoleBinding
	CreateOrUpdateRb(context.Context, *RoleBinding) (*Response, error)
	// 删除RoleBinding
	DeleteRb(context.Context, *DeleteRbRequest) (*Response, error)
	// 查看RoleBinding列表
	GetRbList(context.Context, *GetRbListRequest) (*RoleBindingSet, error)
	// 查询RoleBinding详情
	GetRbDetail(context.Context, *GetRbDetailRequest) (*RoleBinding, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreateOrUpdateServiceAccount(context.Context, *ServiceAccount) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateServiceAccount not implemented")
}
func (UnimplementedRPCServer) DeleteServiceAccount(context.Context, *DeleteServiceAccountRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceAccount not implemented")
}
func (UnimplementedRPCServer) GetServiceAccountList(context.Context, *GetServiceAccountListRequest) (*ServiceAccountSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceAccountList not implemented")
}
func (UnimplementedRPCServer) CreateOrUpdateClusterRole(context.Context, *Role) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateClusterRole not implemented")
}
func (UnimplementedRPCServer) DeleteClusterRole(context.Context, *DeleteClusterRoleRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClusterRole not implemented")
}
func (UnimplementedRPCServer) GetClusterRoleList(context.Context, *GetClusterRoleListRequest) (*ClusterRoleSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterRoleList not implemented")
}
func (UnimplementedRPCServer) GetClusterRoleDetail(context.Context, *GetClusterRoleDetailRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterRoleDetail not implemented")
}
func (UnimplementedRPCServer) CreateOrUpdateRole(context.Context, *Role) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateRole not implemented")
}
func (UnimplementedRPCServer) DeleteRole(context.Context, *DeleteRoleRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRPCServer) GetRoleList(context.Context, *GetRoleListRequest) (*RoleSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleList not implemented")
}
func (UnimplementedRPCServer) GetRoleDetail(context.Context, *GetRoleDetailRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleDetail not implemented")
}
func (UnimplementedRPCServer) CreateOrUpdateRb(context.Context, *RoleBinding) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateRb not implemented")
}
func (UnimplementedRPCServer) DeleteRb(context.Context, *DeleteRbRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRb not implemented")
}
func (UnimplementedRPCServer) GetRbList(context.Context, *GetRbListRequest) (*RoleBindingSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRbList not implemented")
}
func (UnimplementedRPCServer) GetRbDetail(context.Context, *GetRbDetailRequest) (*RoleBinding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRbDetail not implemented")
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

func _RPC_CreateOrUpdateServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateOrUpdateServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateOrUpdateServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateOrUpdateServiceAccount(ctx, req.(*ServiceAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteServiceAccount(ctx, req.(*DeleteServiceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_GetServiceAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceAccountListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).GetServiceAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_GetServiceAccountList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).GetServiceAccountList(ctx, req.(*GetServiceAccountListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_CreateOrUpdateClusterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateOrUpdateClusterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateOrUpdateClusterRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateOrUpdateClusterRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteClusterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteClusterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteClusterRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteClusterRole(ctx, req.(*DeleteClusterRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_GetClusterRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).GetClusterRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_GetClusterRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).GetClusterRoleList(ctx, req.(*GetClusterRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_GetClusterRoleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRoleDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).GetClusterRoleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_GetClusterRoleDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).GetClusterRoleDetail(ctx, req.(*GetClusterRoleDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_CreateOrUpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateOrUpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateOrUpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateOrUpdateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_GetRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).GetRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_GetRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).GetRoleList(ctx, req.(*GetRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_GetRoleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).GetRoleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_GetRoleDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).GetRoleDetail(ctx, req.(*GetRoleDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_CreateOrUpdateRb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateOrUpdateRb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateOrUpdateRb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateOrUpdateRb(ctx, req.(*RoleBinding))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DeleteRb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DeleteRb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DeleteRb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DeleteRb(ctx, req.(*DeleteRbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_GetRbList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRbListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).GetRbList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_GetRbList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).GetRbList(ctx, req.(*GetRbListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_GetRbDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRbDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).GetRbDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_GetRbDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).GetRbDetail(ctx, req.(*GetRbDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codehorse.mkube.rbac.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateServiceAccount",
			Handler:    _RPC_CreateOrUpdateServiceAccount_Handler,
		},
		{
			MethodName: "DeleteServiceAccount",
			Handler:    _RPC_DeleteServiceAccount_Handler,
		},
		{
			MethodName: "GetServiceAccountList",
			Handler:    _RPC_GetServiceAccountList_Handler,
		},
		{
			MethodName: "CreateOrUpdateClusterRole",
			Handler:    _RPC_CreateOrUpdateClusterRole_Handler,
		},
		{
			MethodName: "DeleteClusterRole",
			Handler:    _RPC_DeleteClusterRole_Handler,
		},
		{
			MethodName: "GetClusterRoleList",
			Handler:    _RPC_GetClusterRoleList_Handler,
		},
		{
			MethodName: "GetClusterRoleDetail",
			Handler:    _RPC_GetClusterRoleDetail_Handler,
		},
		{
			MethodName: "CreateOrUpdateRole",
			Handler:    _RPC_CreateOrUpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RPC_DeleteRole_Handler,
		},
		{
			MethodName: "GetRoleList",
			Handler:    _RPC_GetRoleList_Handler,
		},
		{
			MethodName: "GetRoleDetail",
			Handler:    _RPC_GetRoleDetail_Handler,
		},
		{
			MethodName: "CreateOrUpdateRb",
			Handler:    _RPC_CreateOrUpdateRb_Handler,
		},
		{
			MethodName: "DeleteRb",
			Handler:    _RPC_DeleteRb_Handler,
		},
		{
			MethodName: "GetRbList",
			Handler:    _RPC_GetRbList_Handler,
		},
		{
			MethodName: "GetRbDetail",
			Handler:    _RPC_GetRbDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/rbac/pb/rpc.proto",
}