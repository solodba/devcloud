// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: apps/secret/pb/rpc.proto

package secret

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

// DeleteSecretRequest结构体
type DeleteSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DeleteSecretRequest) Reset() {
	*x = DeleteSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretRequest) ProtoMessage() {}

func (x *DeleteSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretRequest.ProtoReflect.Descriptor instead.
func (*DeleteSecretRequest) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteSecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpdateSecretRequest结构体
type UpdateSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"secret" json:"secret"
	Secret *CreateSecretRequest `protobuf:"bytes,1,opt,name=Secret,proto3" json:"secret" bson:"secret"`
}

func (x *UpdateSecretRequest) Reset() {
	*x = UpdateSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSecretRequest) ProtoMessage() {}

func (x *UpdateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSecretRequest.ProtoReflect.Descriptor instead.
func (*UpdateSecretRequest) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateSecretRequest) GetSecret() *CreateSecretRequest {
	if x != nil {
		return x.Secret
	}
	return nil
}

// QuerySecretRequest结构体
type QuerySecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"keyword" json:"keyword"
	Keyword string `protobuf:"bytes,2,opt,name=Keyword,proto3" json:"keyword" bson:"keyword"`
}

func (x *QuerySecretRequest) Reset() {
	*x = QuerySecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySecretRequest) ProtoMessage() {}

func (x *QuerySecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySecretRequest.ProtoReflect.Descriptor instead.
func (*QuerySecretRequest) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QuerySecretRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// DescribeSecretRequest结构体
type DescribeSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DescribeSecretRequest) Reset() {
	*x = DescribeSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSecretRequest) ProtoMessage() {}

func (x *DescribeSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSecretRequest.ProtoReflect.Descriptor instead.
func (*DescribeSecretRequest) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeSecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DescribeSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_apps_secret_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_secret_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x64, 0x65,
	0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x1a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x43, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x4c, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xe9, 0x03,
	0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x5b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73,
	0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x62, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x5b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x5c, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65,
	0x74, 0x12, 0x66, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_secret_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_secret_pb_rpc_proto_rawDescData = file_apps_secret_pb_rpc_proto_rawDesc
)

func file_apps_secret_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_secret_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_secret_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_secret_pb_rpc_proto_rawDescData)
	})
	return file_apps_secret_pb_rpc_proto_rawDescData
}

var file_apps_secret_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_secret_pb_rpc_proto_goTypes = []interface{}{
	(*DeleteSecretRequest)(nil),   // 0: codehorse.mpaas.secret.DeleteSecretRequest
	(*UpdateSecretRequest)(nil),   // 1: codehorse.mpaas.secret.UpdateSecretRequest
	(*QuerySecretRequest)(nil),    // 2: codehorse.mpaas.secret.QuerySecretRequest
	(*DescribeSecretRequest)(nil), // 3: codehorse.mpaas.secret.DescribeSecretRequest
	(*CreateSecretRequest)(nil),   // 4: codehorse.mpaas.secret.CreateSecretRequest
	(*Secret)(nil),                // 5: codehorse.mpaas.secret.Secret
	(*SecretSetItem)(nil),         // 6: codehorse.mpaas.secret.SecretSetItem
	(*SecretSet)(nil),             // 7: codehorse.mpaas.secret.SecretSet
}
var file_apps_secret_pb_rpc_proto_depIdxs = []int32{
	4, // 0: codehorse.mpaas.secret.UpdateSecretRequest.Secret:type_name -> codehorse.mpaas.secret.CreateSecretRequest
	4, // 1: codehorse.mpaas.secret.RPC.CreateSecret:input_type -> codehorse.mpaas.secret.CreateSecretRequest
	0, // 2: codehorse.mpaas.secret.RPC.DeleteSecret:input_type -> codehorse.mpaas.secret.DeleteSecretRequest
	1, // 3: codehorse.mpaas.secret.RPC.UpdateSecret:input_type -> codehorse.mpaas.secret.UpdateSecretRequest
	2, // 4: codehorse.mpaas.secret.RPC.QuerySecret:input_type -> codehorse.mpaas.secret.QuerySecretRequest
	3, // 5: codehorse.mpaas.secret.RPC.DescribeSecret:input_type -> codehorse.mpaas.secret.DescribeSecretRequest
	5, // 6: codehorse.mpaas.secret.RPC.CreateSecret:output_type -> codehorse.mpaas.secret.Secret
	6, // 7: codehorse.mpaas.secret.RPC.DeleteSecret:output_type -> codehorse.mpaas.secret.SecretSetItem
	5, // 8: codehorse.mpaas.secret.RPC.UpdateSecret:output_type -> codehorse.mpaas.secret.Secret
	7, // 9: codehorse.mpaas.secret.RPC.QuerySecret:output_type -> codehorse.mpaas.secret.SecretSet
	6, // 10: codehorse.mpaas.secret.RPC.DescribeSecret:output_type -> codehorse.mpaas.secret.SecretSetItem
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_secret_pb_rpc_proto_init() }
func file_apps_secret_pb_rpc_proto_init() {
	if File_apps_secret_pb_rpc_proto != nil {
		return
	}
	file_apps_secret_pb_secret_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_secret_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSecretRequest); i {
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
		file_apps_secret_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSecretRequest); i {
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
		file_apps_secret_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySecretRequest); i {
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
		file_apps_secret_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSecretRequest); i {
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
			RawDescriptor: file_apps_secret_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_secret_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_secret_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_secret_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_secret_pb_rpc_proto = out.File
	file_apps_secret_pb_rpc_proto_rawDesc = nil
	file_apps_secret_pb_rpc_proto_goTypes = nil
	file_apps_secret_pb_rpc_proto_depIdxs = nil
}
