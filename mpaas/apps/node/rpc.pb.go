// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/node/pb/rpc.proto

package node

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

// QueryNodeRequest结构体
type QueryNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"keyword" json:"keyword"
	Keyword string `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"keyword" bson:"keyword"`
}

func (x *QueryNodeRequest) Reset() {
	*x = QueryNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNodeRequest) ProtoMessage() {}

func (x *QueryNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNodeRequest.ProtoReflect.Descriptor instead.
func (*QueryNodeRequest) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryNodeRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// DescribeNodeRequest结构体
type DescribeNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DescribeNodeRequest) Reset() {
	*x = DescribeNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeNodeRequest) ProtoMessage() {}

func (x *DescribeNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeNodeRequest.ProtoReflect.Descriptor instead.
func (*DescribeNodeRequest) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeNodeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpdatedLabelRequest结构体
type UpdatedLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,2,rep,name=Labels,proto3" json:"labels" bson:"labels"`
}

func (x *UpdatedLabelRequest) Reset() {
	*x = UpdatedLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedLabelRequest) ProtoMessage() {}

func (x *UpdatedLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedLabelRequest.ProtoReflect.Descriptor instead.
func (*UpdatedLabelRequest) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatedLabelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatedLabelRequest) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

// UpdatedTaintRequest结构体
type UpdatedTaintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"taints" json:"taints"
	Taints []*Taint `protobuf:"bytes,2,rep,name=Taints,proto3" json:"taints" bson:"taints"`
}

func (x *UpdatedTaintRequest) Reset() {
	*x = UpdatedTaintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedTaintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedTaintRequest) ProtoMessage() {}

func (x *UpdatedTaintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedTaintRequest.ProtoReflect.Descriptor instead.
func (*UpdatedTaintRequest) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatedTaintRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatedTaintRequest) GetTaints() []*Taint {
	if x != nil {
		return x.Taints
	}
	return nil
}

// UpdatedLabelResponse结构体
type UpdatedLabelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatedLabelResponse) Reset() {
	*x = UpdatedLabelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedLabelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedLabelResponse) ProtoMessage() {}

func (x *UpdatedLabelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedLabelResponse.ProtoReflect.Descriptor instead.
func (*UpdatedLabelResponse) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_rpc_proto_rawDescGZIP(), []int{4}
}

// UpdatedTaintResponse结构体
type UpdatedTaintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatedTaintResponse) Reset() {
	*x = UpdatedTaintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedTaintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedTaintResponse) ProtoMessage() {}

func (x *UpdatedTaintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedTaintResponse.ProtoReflect.Descriptor instead.
func (*UpdatedTaintResponse) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_rpc_proto_rawDescGZIP(), []int{5}
}

var File_apps_node_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_node_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x17,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x64, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x5e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x54, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8b, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x52,
	0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x74, 0x12, 0x5c, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x68, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_node_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_node_pb_rpc_proto_rawDescData = file_apps_node_pb_rpc_proto_rawDesc
)

func file_apps_node_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_node_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_node_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_node_pb_rpc_proto_rawDescData)
	})
	return file_apps_node_pb_rpc_proto_rawDescData
}

var file_apps_node_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_node_pb_rpc_proto_goTypes = []interface{}{
	(*QueryNodeRequest)(nil),     // 0: codehorse.mpaas.node.QueryNodeRequest
	(*DescribeNodeRequest)(nil),  // 1: codehorse.mpaas.node.DescribeNodeRequest
	(*UpdatedLabelRequest)(nil),  // 2: codehorse.mpaas.node.UpdatedLabelRequest
	(*UpdatedTaintRequest)(nil),  // 3: codehorse.mpaas.node.UpdatedTaintRequest
	(*UpdatedLabelResponse)(nil), // 4: codehorse.mpaas.node.UpdatedLabelResponse
	(*UpdatedTaintResponse)(nil), // 5: codehorse.mpaas.node.UpdatedTaintResponse
	(*ListMapItem)(nil),          // 6: codehorse.mpaas.node.ListMapItem
	(*Taint)(nil),                // 7: codehorse.mpaas.node.Taint
	(*NodeSet)(nil),              // 8: codehorse.mpaas.node.NodeSet
	(*NodeSetItem)(nil),          // 9: codehorse.mpaas.node.NodeSetItem
}
var file_apps_node_pb_rpc_proto_depIdxs = []int32{
	6, // 0: codehorse.mpaas.node.UpdatedLabelRequest.Labels:type_name -> codehorse.mpaas.node.ListMapItem
	7, // 1: codehorse.mpaas.node.UpdatedTaintRequest.Taints:type_name -> codehorse.mpaas.node.Taint
	0, // 2: codehorse.mpaas.node.RPC.QueryNode:input_type -> codehorse.mpaas.node.QueryNodeRequest
	1, // 3: codehorse.mpaas.node.RPC.DescribeNode:input_type -> codehorse.mpaas.node.DescribeNodeRequest
	2, // 4: codehorse.mpaas.node.RPC.UpdateNodeLabel:input_type -> codehorse.mpaas.node.UpdatedLabelRequest
	3, // 5: codehorse.mpaas.node.RPC.UpdateNodeTaint:input_type -> codehorse.mpaas.node.UpdatedTaintRequest
	8, // 6: codehorse.mpaas.node.RPC.QueryNode:output_type -> codehorse.mpaas.node.NodeSet
	9, // 7: codehorse.mpaas.node.RPC.DescribeNode:output_type -> codehorse.mpaas.node.NodeSetItem
	4, // 8: codehorse.mpaas.node.RPC.UpdateNodeLabel:output_type -> codehorse.mpaas.node.UpdatedLabelResponse
	5, // 9: codehorse.mpaas.node.RPC.UpdateNodeTaint:output_type -> codehorse.mpaas.node.UpdatedTaintResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apps_node_pb_rpc_proto_init() }
func file_apps_node_pb_rpc_proto_init() {
	if File_apps_node_pb_rpc_proto != nil {
		return
	}
	file_apps_node_pb_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_node_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNodeRequest); i {
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
		file_apps_node_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeNodeRequest); i {
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
		file_apps_node_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedLabelRequest); i {
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
		file_apps_node_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedTaintRequest); i {
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
		file_apps_node_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedLabelResponse); i {
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
		file_apps_node_pb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedTaintResponse); i {
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
			RawDescriptor: file_apps_node_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_node_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_node_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_node_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_node_pb_rpc_proto = out.File
	file_apps_node_pb_rpc_proto_rawDesc = nil
	file_apps_node_pb_rpc_proto_goTypes = nil
	file_apps_node_pb_rpc_proto_depIdxs = nil
}
