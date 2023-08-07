// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/ingroute/pb/rpc.proto

package ingroute

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// UpdateIngressRouteRequest结构体
type UpdateIngressRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"ingressRoute" json:"ingressRoute"
	IngressRoute *CreateIngressRouteRequest `protobuf:"bytes,1,opt,name=IngressRoute,proto3" json:"ingressRoute" bson:"ingressRoute"`
}

func (x *UpdateIngressRouteRequest) Reset() {
	*x = UpdateIngressRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIngressRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIngressRouteRequest) ProtoMessage() {}

func (x *UpdateIngressRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIngressRouteRequest.ProtoReflect.Descriptor instead.
func (*UpdateIngressRouteRequest) Descriptor() ([]byte, []int) {
	return file_apps_ingroute_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateIngressRouteRequest) GetIngressRoute() *CreateIngressRouteRequest {
	if x != nil {
		return x.IngressRoute
	}
	return nil
}

// DeleteIngressRouteRequest结构体
type DeleteIngressRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DeleteIngressRouteRequest) Reset() {
	*x = DeleteIngressRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIngressRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIngressRouteRequest) ProtoMessage() {}

func (x *DeleteIngressRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIngressRouteRequest.ProtoReflect.Descriptor instead.
func (*DeleteIngressRouteRequest) Descriptor() ([]byte, []int) {
	return file_apps_ingroute_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteIngressRouteRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteIngressRouteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// QueryIngressRouteRequest结构体
type QueryIngressRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"keyword" json:"keyword"
	Keyword string `protobuf:"bytes,2,opt,name=Keyword,proto3" json:"keyword" bson:"keyword"`
}

func (x *QueryIngressRouteRequest) Reset() {
	*x = QueryIngressRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryIngressRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryIngressRouteRequest) ProtoMessage() {}

func (x *QueryIngressRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryIngressRouteRequest.ProtoReflect.Descriptor instead.
func (*QueryIngressRouteRequest) Descriptor() ([]byte, []int) {
	return file_apps_ingroute_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryIngressRouteRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryIngressRouteRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// DescribeIngressRouteRequest结构体
type DescribeIngressRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DescribeIngressRouteRequest) Reset() {
	*x = DescribeIngressRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeIngressRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeIngressRouteRequest) ProtoMessage() {}

func (x *DescribeIngressRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeIngressRouteRequest.ProtoReflect.Descriptor instead.
func (*DescribeIngressRouteRequest) Descriptor() ([]byte, []int) {
	return file_apps_ingroute_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeIngressRouteRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DescribeIngressRouteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// QueryIngRouteMwRequest结构体
type QueryIngRouteMwRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
}

func (x *QueryIngRouteMwRequest) Reset() {
	*x = QueryIngRouteMwRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryIngRouteMwRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryIngRouteMwRequest) ProtoMessage() {}

func (x *QueryIngRouteMwRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_ingroute_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryIngRouteMwRequest.ProtoReflect.Descriptor instead.
func (*QueryIngRouteMwRequest) Descriptor() ([]byte, []int) {
	return file_apps_ingroute_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *QueryIngRouteMwRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_apps_ingroute_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_ingroute_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e,
	0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x69, 0x6e, 0x67,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x0c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e, 0x67,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x4d, 0x0a,
	0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x18,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x4f, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x36, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x67, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0xdf, 0x05, 0x0a, 0x03, 0x52, 0x50,
	0x43, 0x12, 0x71, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69,
	0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x71, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e, 0x67,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x7e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x33, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x72, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69,
	0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x82, 0x01, 0x0a, 0x14,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e,
	0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x79, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x30, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x49, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b,
	0x75, 0x62, 0x65, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x4d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62,
	0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6b, 0x75, 0x62, 0x65,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_ingroute_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_ingroute_pb_rpc_proto_rawDescData = file_apps_ingroute_pb_rpc_proto_rawDesc
)

func file_apps_ingroute_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_ingroute_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_ingroute_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_ingroute_pb_rpc_proto_rawDescData)
	})
	return file_apps_ingroute_pb_rpc_proto_rawDescData
}

var file_apps_ingroute_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_ingroute_pb_rpc_proto_goTypes = []interface{}{
	(*UpdateIngressRouteRequest)(nil),   // 0: codehorse.mkube.ingroute.UpdateIngressRouteRequest
	(*DeleteIngressRouteRequest)(nil),   // 1: codehorse.mkube.ingroute.DeleteIngressRouteRequest
	(*QueryIngressRouteRequest)(nil),    // 2: codehorse.mkube.ingroute.QueryIngressRouteRequest
	(*DescribeIngressRouteRequest)(nil), // 3: codehorse.mkube.ingroute.DescribeIngressRouteRequest
	(*QueryIngRouteMwRequest)(nil),      // 4: codehorse.mkube.ingroute.QueryIngRouteMwRequest
	(*CreateIngressRouteRequest)(nil),   // 5: codehorse.mkube.ingroute.CreateIngressRouteRequest
	(*IngressRoute)(nil),                // 6: codehorse.mkube.ingroute.IngressRoute
	(*IngressRouteSet)(nil),             // 7: codehorse.mkube.ingroute.IngressRouteSet
	(*MiddlewareList)(nil),              // 8: codehorse.mkube.ingroute.MiddlewareList
}
var file_apps_ingroute_pb_rpc_proto_depIdxs = []int32{
	5, // 0: codehorse.mkube.ingroute.UpdateIngressRouteRequest.IngressRoute:type_name -> codehorse.mkube.ingroute.CreateIngressRouteRequest
	5, // 1: codehorse.mkube.ingroute.RPC.CreateIngressRoute:input_type -> codehorse.mkube.ingroute.CreateIngressRouteRequest
	0, // 2: codehorse.mkube.ingroute.RPC.UpdateIngressRoute:input_type -> codehorse.mkube.ingroute.UpdateIngressRouteRequest
	1, // 3: codehorse.mkube.ingroute.RPC.DeleteIngressRoute:input_type -> codehorse.mkube.ingroute.DeleteIngressRouteRequest
	2, // 4: codehorse.mkube.ingroute.RPC.QueryIngressRoute:input_type -> codehorse.mkube.ingroute.QueryIngressRouteRequest
	3, // 5: codehorse.mkube.ingroute.RPC.DescribeIngressRoute:input_type -> codehorse.mkube.ingroute.DescribeIngressRouteRequest
	4, // 6: codehorse.mkube.ingroute.RPC.QueryIngRouteMiddlewareList:input_type -> codehorse.mkube.ingroute.QueryIngRouteMwRequest
	6, // 7: codehorse.mkube.ingroute.RPC.CreateIngressRoute:output_type -> codehorse.mkube.ingroute.IngressRoute
	6, // 8: codehorse.mkube.ingroute.RPC.UpdateIngressRoute:output_type -> codehorse.mkube.ingroute.IngressRoute
	5, // 9: codehorse.mkube.ingroute.RPC.DeleteIngressRoute:output_type -> codehorse.mkube.ingroute.CreateIngressRouteRequest
	7, // 10: codehorse.mkube.ingroute.RPC.QueryIngressRoute:output_type -> codehorse.mkube.ingroute.IngressRouteSet
	5, // 11: codehorse.mkube.ingroute.RPC.DescribeIngressRoute:output_type -> codehorse.mkube.ingroute.CreateIngressRouteRequest
	8, // 12: codehorse.mkube.ingroute.RPC.QueryIngRouteMiddlewareList:output_type -> codehorse.mkube.ingroute.MiddlewareList
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_ingroute_pb_rpc_proto_init() }
func file_apps_ingroute_pb_rpc_proto_init() {
	if File_apps_ingroute_pb_rpc_proto != nil {
		return
	}
	file_apps_ingroute_pb_ingroute_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_ingroute_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIngressRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_ingroute_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIngressRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_ingroute_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryIngressRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_ingroute_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeIngressRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_ingroute_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryIngRouteMwRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_ingroute_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_ingroute_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_ingroute_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_ingroute_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_ingroute_pb_rpc_proto = out.File
	file_apps_ingroute_pb_rpc_proto_rawDesc = nil
	file_apps_ingroute_pb_rpc_proto_goTypes = nil
	file_apps_ingroute_pb_rpc_proto_depIdxs = nil
}
