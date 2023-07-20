// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: apps/secret/pb/secret.proto

package secret

import (
	meta "github.com/solodba/mcube/pb/meta"
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

// Secret结构体
type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"secret"
	Secret *CreateSecretRequest `protobuf:"bytes,2,opt,name=Secret,proto3" json:"secret" bson:",inline"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_secret_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Secret) GetSecret() *CreateSecretRequest {
	if x != nil {
		return x.Secret
	}
	return nil
}

// SecretSet结构体
type SecretSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*SecretSetItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"items" bson:"items"`
}

func (x *SecretSet) Reset() {
	*x = SecretSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretSet) ProtoMessage() {}

func (x *SecretSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretSet.ProtoReflect.Descriptor instead.
func (*SecretSet) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_secret_proto_rawDescGZIP(), []int{1}
}

func (x *SecretSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SecretSet) GetItems() []*SecretSetItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// CreateSecretRequest结构体
type CreateSecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,3,rep,name=Labels,proto3" json:"labels" bson:"labels"`
	// @gotags: bson:"data" json:"data"
	Data []*ListMapItem `protobuf:"bytes,4,rep,name=Data,proto3" json:"data" bson:"data"`
}

func (x *CreateSecretRequest) Reset() {
	*x = CreateSecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretRequest) ProtoMessage() {}

func (x *CreateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretRequest.ProtoReflect.Descriptor instead.
func (*CreateSecretRequest) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_secret_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSecretRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateSecretRequest) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateSecretRequest) GetData() []*ListMapItem {
	if x != nil {
		return x.Data
	}
	return nil
}

// 通用键值对结构体
type ListMapItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"key" json:"key"
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"key" bson:"key"`
	// @gotags: bson:"value" json:"value"
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"value" bson:"value"`
}

func (x *ListMapItem) Reset() {
	*x = ListMapItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMapItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMapItem) ProtoMessage() {}

func (x *ListMapItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_secret_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMapItem.ProtoReflect.Descriptor instead.
func (*ListMapItem) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_secret_proto_rawDescGZIP(), []int{3}
}

func (x *ListMapItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListMapItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// SecretSetItem结构体
type SecretSetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"dataNum" json:"dataNum"
	DataNum int64 `protobuf:"varint,3,opt,name=DataNum,proto3" json:"dataNum" bson:"dataNum"`
	// @gotags: bson:"type" json:"type"
	Type string `protobuf:"bytes,4,opt,name=Type,proto3" json:"type" bson:"type"`
	// @gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,5,rep,name=Labels,proto3" json:"labels" bson:"labels"`
	// @gotags: bson:"data" json:"data"
	Data []*ListMapItem `protobuf:"bytes,6,rep,name=Data,proto3" json:"data" bson:"data"`
}

func (x *SecretSetItem) Reset() {
	*x = SecretSetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_secret_pb_secret_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretSetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretSetItem) ProtoMessage() {}

func (x *SecretSetItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_secret_pb_secret_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretSetItem.ProtoReflect.Descriptor instead.
func (*SecretSetItem) Descriptor() ([]byte, []int) {
	return file_apps_secret_pb_secret_proto_rawDescGZIP(), []int{4}
}

func (x *SecretSetItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecretSetItem) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SecretSetItem) GetDataNum() int64 {
	if x != nil {
		return x.DataNum
	}
	return 0
}

func (x *SecretSetItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SecretSetItem) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *SecretSetItem) GetData() []*ListMapItem {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_apps_secret_pb_secret_proto protoreflect.FileDescriptor

var file_apps_secret_pb_secret_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x62,
	0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x64, 0x62, 0x61, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x06,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x5e, 0x0a, 0x09, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3b,
	0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73,
	0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61,
	0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x75, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x37, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61,
	0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_secret_pb_secret_proto_rawDescOnce sync.Once
	file_apps_secret_pb_secret_proto_rawDescData = file_apps_secret_pb_secret_proto_rawDesc
)

func file_apps_secret_pb_secret_proto_rawDescGZIP() []byte {
	file_apps_secret_pb_secret_proto_rawDescOnce.Do(func() {
		file_apps_secret_pb_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_secret_pb_secret_proto_rawDescData)
	})
	return file_apps_secret_pb_secret_proto_rawDescData
}

var file_apps_secret_pb_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_secret_pb_secret_proto_goTypes = []interface{}{
	(*Secret)(nil),              // 0: codehorse.mpaas.secret.Secret
	(*SecretSet)(nil),           // 1: codehorse.mpaas.secret.SecretSet
	(*CreateSecretRequest)(nil), // 2: codehorse.mpaas.secret.CreateSecretRequest
	(*ListMapItem)(nil),         // 3: codehorse.mpaas.secret.ListMapItem
	(*SecretSetItem)(nil),       // 4: codehorse.mpaas.secret.SecretSetItem
	(*meta.Meta)(nil),           // 5: codehorse.mcube.meta.Meta
}
var file_apps_secret_pb_secret_proto_depIdxs = []int32{
	5, // 0: codehorse.mpaas.secret.Secret.Meta:type_name -> codehorse.mcube.meta.Meta
	2, // 1: codehorse.mpaas.secret.Secret.Secret:type_name -> codehorse.mpaas.secret.CreateSecretRequest
	4, // 2: codehorse.mpaas.secret.SecretSet.Items:type_name -> codehorse.mpaas.secret.SecretSetItem
	3, // 3: codehorse.mpaas.secret.CreateSecretRequest.Labels:type_name -> codehorse.mpaas.secret.ListMapItem
	3, // 4: codehorse.mpaas.secret.CreateSecretRequest.Data:type_name -> codehorse.mpaas.secret.ListMapItem
	3, // 5: codehorse.mpaas.secret.SecretSetItem.Labels:type_name -> codehorse.mpaas.secret.ListMapItem
	3, // 6: codehorse.mpaas.secret.SecretSetItem.Data:type_name -> codehorse.mpaas.secret.ListMapItem
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_secret_pb_secret_proto_init() }
func file_apps_secret_pb_secret_proto_init() {
	if File_apps_secret_pb_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_secret_pb_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_apps_secret_pb_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretSet); i {
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
		file_apps_secret_pb_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSecretRequest); i {
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
		file_apps_secret_pb_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMapItem); i {
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
		file_apps_secret_pb_secret_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretSetItem); i {
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
			RawDescriptor: file_apps_secret_pb_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_secret_pb_secret_proto_goTypes,
		DependencyIndexes: file_apps_secret_pb_secret_proto_depIdxs,
		MessageInfos:      file_apps_secret_pb_secret_proto_msgTypes,
	}.Build()
	File_apps_secret_pb_secret_proto = out.File
	file_apps_secret_pb_secret_proto_rawDesc = nil
	file_apps_secret_pb_secret_proto_goTypes = nil
	file_apps_secret_pb_secret_proto_depIdxs = nil
}
