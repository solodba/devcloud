// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/daemonset/pb/rpc.proto

package daemonset

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

// DeleteDaemonSetRequest结构体
type DeleteDaemonSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DeleteDaemonSetRequest) Reset() {
	*x = DeleteDaemonSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDaemonSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDaemonSetRequest) ProtoMessage() {}

func (x *DeleteDaemonSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDaemonSetRequest.ProtoReflect.Descriptor instead.
func (*DeleteDaemonSetRequest) Descriptor() ([]byte, []int) {
	return file_apps_daemonset_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteDaemonSetRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteDaemonSetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpdateDaemonSetRequest结构体
type UpdateDaemonSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"daemonset" json:"daemonset"
	DaemonSet *CreateDaemonSetRequest `protobuf:"bytes,1,opt,name=DaemonSet,proto3" json:"daemonset" bson:"daemonset"`
}

func (x *UpdateDaemonSetRequest) Reset() {
	*x = UpdateDaemonSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDaemonSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDaemonSetRequest) ProtoMessage() {}

func (x *UpdateDaemonSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDaemonSetRequest.ProtoReflect.Descriptor instead.
func (*UpdateDaemonSetRequest) Descriptor() ([]byte, []int) {
	return file_apps_daemonset_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDaemonSetRequest) GetDaemonSet() *CreateDaemonSetRequest {
	if x != nil {
		return x.DaemonSet
	}
	return nil
}

// QueryDaemonSetRequest结构体
type QueryDaemonSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"keyword" json:"keyword"
	Keyword string `protobuf:"bytes,2,opt,name=Keyword,proto3" json:"keyword" bson:"keyword"`
}

func (x *QueryDaemonSetRequest) Reset() {
	*x = QueryDaemonSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDaemonSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDaemonSetRequest) ProtoMessage() {}

func (x *QueryDaemonSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDaemonSetRequest.ProtoReflect.Descriptor instead.
func (*QueryDaemonSetRequest) Descriptor() ([]byte, []int) {
	return file_apps_daemonset_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryDaemonSetRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryDaemonSetRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// DescribeDaemonSetRequest结构体
type DescribeDaemonSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DescribeDaemonSetRequest) Reset() {
	*x = DescribeDaemonSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDaemonSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDaemonSetRequest) ProtoMessage() {}

func (x *DescribeDaemonSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_daemonset_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDaemonSetRequest.ProtoReflect.Descriptor instead.
func (*DescribeDaemonSetRequest) Descriptor() ([]byte, []int) {
	return file_apps_daemonset_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeDaemonSetRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DescribeDaemonSetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_apps_daemonset_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_daemonset_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x1a, 0x21, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4f, 0x0a, 0x09, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x22, 0x4f, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x4c, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0xc1, 0x04, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x6a, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x31, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x77, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x73, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6a,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74,
	0x2e, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x6c, 0x0a, 0x0e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x30, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x64,
	0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x7b, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x33, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_daemonset_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_daemonset_pb_rpc_proto_rawDescData = file_apps_daemonset_pb_rpc_proto_rawDesc
)

func file_apps_daemonset_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_daemonset_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_daemonset_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_daemonset_pb_rpc_proto_rawDescData)
	})
	return file_apps_daemonset_pb_rpc_proto_rawDescData
}

var file_apps_daemonset_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_daemonset_pb_rpc_proto_goTypes = []interface{}{
	(*DeleteDaemonSetRequest)(nil),   // 0: codehorse.mpaas.daemonset.DeleteDaemonSetRequest
	(*UpdateDaemonSetRequest)(nil),   // 1: codehorse.mpaas.daemonset.UpdateDaemonSetRequest
	(*QueryDaemonSetRequest)(nil),    // 2: codehorse.mpaas.daemonset.QueryDaemonSetRequest
	(*DescribeDaemonSetRequest)(nil), // 3: codehorse.mpaas.daemonset.DescribeDaemonSetRequest
	(*CreateDaemonSetRequest)(nil),   // 4: codehorse.mpaas.daemonset.CreateDaemonSetRequest
	(*DaemonSet)(nil),                // 5: codehorse.mpaas.daemonset.DaemonSet
	(*DaemonSetList)(nil),            // 6: codehorse.mpaas.daemonset.DaemonSetList
}
var file_apps_daemonset_pb_rpc_proto_depIdxs = []int32{
	4, // 0: codehorse.mpaas.daemonset.UpdateDaemonSetRequest.DaemonSet:type_name -> codehorse.mpaas.daemonset.CreateDaemonSetRequest
	4, // 1: codehorse.mpaas.daemonset.RPC.CreateDaemonSet:input_type -> codehorse.mpaas.daemonset.CreateDaemonSetRequest
	0, // 2: codehorse.mpaas.daemonset.RPC.DeleteDaemonSet:input_type -> codehorse.mpaas.daemonset.DeleteDaemonSetRequest
	1, // 3: codehorse.mpaas.daemonset.RPC.UpdateDaemonSet:input_type -> codehorse.mpaas.daemonset.UpdateDaemonSetRequest
	2, // 4: codehorse.mpaas.daemonset.RPC.QueryDaemonSet:input_type -> codehorse.mpaas.daemonset.QueryDaemonSetRequest
	3, // 5: codehorse.mpaas.daemonset.RPC.DescribeDaemonSet:input_type -> codehorse.mpaas.daemonset.DescribeDaemonSetRequest
	5, // 6: codehorse.mpaas.daemonset.RPC.CreateDaemonSet:output_type -> codehorse.mpaas.daemonset.DaemonSet
	4, // 7: codehorse.mpaas.daemonset.RPC.DeleteDaemonSet:output_type -> codehorse.mpaas.daemonset.CreateDaemonSetRequest
	5, // 8: codehorse.mpaas.daemonset.RPC.UpdateDaemonSet:output_type -> codehorse.mpaas.daemonset.DaemonSet
	6, // 9: codehorse.mpaas.daemonset.RPC.QueryDaemonSet:output_type -> codehorse.mpaas.daemonset.DaemonSetList
	4, // 10: codehorse.mpaas.daemonset.RPC.DescribeDaemonSet:output_type -> codehorse.mpaas.daemonset.CreateDaemonSetRequest
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_daemonset_pb_rpc_proto_init() }
func file_apps_daemonset_pb_rpc_proto_init() {
	if File_apps_daemonset_pb_rpc_proto != nil {
		return
	}
	file_apps_daemonset_pb_daemonset_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_daemonset_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDaemonSetRequest); i {
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
		file_apps_daemonset_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDaemonSetRequest); i {
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
		file_apps_daemonset_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDaemonSetRequest); i {
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
		file_apps_daemonset_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDaemonSetRequest); i {
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
			RawDescriptor: file_apps_daemonset_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_daemonset_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_daemonset_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_daemonset_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_daemonset_pb_rpc_proto = out.File
	file_apps_daemonset_pb_rpc_proto_rawDesc = nil
	file_apps_daemonset_pb_rpc_proto_goTypes = nil
	file_apps_daemonset_pb_rpc_proto_depIdxs = nil
}