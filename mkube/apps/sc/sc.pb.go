// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/sc/pb/sc.proto

package sc

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

// CreateSCRequest结构体
type CreateSCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	// gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,3,rep,name=Labels,proto3" json:"Labels,omitempty"`
	// gotags: bson:"provisioner" json:"provisioner"
	Provisioner string `protobuf:"bytes,4,opt,name=Provisioner,proto3" json:"Provisioner,omitempty"`
	// gotags: bson:"mountOptions" json:"mountOptions"
	MountOptions []string `protobuf:"bytes,5,rep,name=MountOptions,proto3" json:"MountOptions,omitempty"`
	// gotags: bson:"parameters" json:"parameters"
	Parameters []*ListMapItem `protobuf:"bytes,6,rep,name=Parameters,proto3" json:"Parameters,omitempty"`
	// gotags: bson:"reClaimPolicy" json:"reClaimPolicy"
	ReClaimPolicy string `protobuf:"bytes,7,opt,name=ReClaimPolicy,proto3" json:"ReClaimPolicy,omitempty"`
	// gotags: bson:"allowVolumeExpansion" json:"allowVolumeExpansion"
	AllowVolumeExpansion bool `protobuf:"varint,8,opt,name=AllowVolumeExpansion,proto3" json:"AllowVolumeExpansion,omitempty"`
	// gotags: bson:"volumeBindingMode" json:"volumeBindingMode"
	VolumeBindingMode string `protobuf:"bytes,9,opt,name=VolumeBindingMode,proto3" json:"VolumeBindingMode,omitempty"`
}

func (x *CreateSCRequest) Reset() {
	*x = CreateSCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_sc_pb_sc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSCRequest) ProtoMessage() {}

func (x *CreateSCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_sc_pb_sc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSCRequest.ProtoReflect.Descriptor instead.
func (*CreateSCRequest) Descriptor() ([]byte, []int) {
	return file_apps_sc_pb_sc_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSCRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSCRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateSCRequest) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateSCRequest) GetProvisioner() string {
	if x != nil {
		return x.Provisioner
	}
	return ""
}

func (x *CreateSCRequest) GetMountOptions() []string {
	if x != nil {
		return x.MountOptions
	}
	return nil
}

func (x *CreateSCRequest) GetParameters() []*ListMapItem {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *CreateSCRequest) GetReClaimPolicy() string {
	if x != nil {
		return x.ReClaimPolicy
	}
	return ""
}

func (x *CreateSCRequest) GetAllowVolumeExpansion() bool {
	if x != nil {
		return x.AllowVolumeExpansion
	}
	return false
}

func (x *CreateSCRequest) GetVolumeBindingMode() string {
	if x != nil {
		return x.VolumeBindingMode
	}
	return ""
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
		mi := &file_apps_sc_pb_sc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMapItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMapItem) ProtoMessage() {}

func (x *ListMapItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_sc_pb_sc_proto_msgTypes[1]
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
	return file_apps_sc_pb_sc_proto_rawDescGZIP(), []int{1}
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

// SC结构体
type SC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"sc"
	SC *CreateSCRequest `protobuf:"bytes,2,opt,name=SC,proto3" json:"sc" bson:",inline"`
}

func (x *SC) Reset() {
	*x = SC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_sc_pb_sc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC) ProtoMessage() {}

func (x *SC) ProtoReflect() protoreflect.Message {
	mi := &file_apps_sc_pb_sc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC.ProtoReflect.Descriptor instead.
func (*SC) Descriptor() ([]byte, []int) {
	return file_apps_sc_pb_sc_proto_rawDescGZIP(), []int{2}
}

func (x *SC) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SC) GetSC() *CreateSCRequest {
	if x != nil {
		return x.SC
	}
	return nil
}

// SCSetItem结构体
type SCSetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	// gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,3,rep,name=Labels,proto3" json:"Labels,omitempty"`
	// gotags: bson:"provisioner" json:"provisioner"
	Provisioner string `protobuf:"bytes,4,opt,name=Provisioner,proto3" json:"Provisioner,omitempty"`
	// gotags: bson:"mountOptions" json:"mountOptions"
	MountOptions []string `protobuf:"bytes,5,rep,name=MountOptions,proto3" json:"MountOptions,omitempty"`
	// gotags: bson:"parameters" json:"parameters"
	Parameters []*ListMapItem `protobuf:"bytes,6,rep,name=Parameters,proto3" json:"Parameters,omitempty"`
	// gotags: bson:"reClaimPolicy" json:"reClaimPolicy"
	ReClaimPolicy string `protobuf:"bytes,7,opt,name=ReClaimPolicy,proto3" json:"ReClaimPolicy,omitempty"`
	// gotags: bson:"allowVolumeExpansion" json:"allowVolumeExpansion"
	AllowVolumeExpansion bool `protobuf:"varint,8,opt,name=AllowVolumeExpansion,proto3" json:"AllowVolumeExpansion,omitempty"`
	// gotags: bson:"volumeBindingMode" json:"volumeBindingMode"
	VolumeBindingMode string `protobuf:"bytes,9,opt,name=VolumeBindingMode,proto3" json:"VolumeBindingMode,omitempty"`
	// gotags: bson:"age" json:"age"
	Age int64 `protobuf:"varint,10,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (x *SCSetItem) Reset() {
	*x = SCSetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_sc_pb_sc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCSetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCSetItem) ProtoMessage() {}

func (x *SCSetItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_sc_pb_sc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCSetItem.ProtoReflect.Descriptor instead.
func (*SCSetItem) Descriptor() ([]byte, []int) {
	return file_apps_sc_pb_sc_proto_rawDescGZIP(), []int{3}
}

func (x *SCSetItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SCSetItem) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SCSetItem) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *SCSetItem) GetProvisioner() string {
	if x != nil {
		return x.Provisioner
	}
	return ""
}

func (x *SCSetItem) GetMountOptions() []string {
	if x != nil {
		return x.MountOptions
	}
	return nil
}

func (x *SCSetItem) GetParameters() []*ListMapItem {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *SCSetItem) GetReClaimPolicy() string {
	if x != nil {
		return x.ReClaimPolicy
	}
	return ""
}

func (x *SCSetItem) GetAllowVolumeExpansion() bool {
	if x != nil {
		return x.AllowVolumeExpansion
	}
	return false
}

func (x *SCSetItem) GetVolumeBindingMode() string {
	if x != nil {
		return x.VolumeBindingMode
	}
	return ""
}

func (x *SCSetItem) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

// SCSet结构体
type SCSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	// gotags: bson:"items" json:"items"
	Items []*SCSetItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *SCSet) Reset() {
	*x = SCSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_sc_pb_sc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCSet) ProtoMessage() {}

func (x *SCSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_sc_pb_sc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCSet.ProtoReflect.Descriptor instead.
func (*SCSet) Descriptor() ([]byte, []int) {
	return file_apps_sc_pb_sc_proto_rawDescGZIP(), []int{4}
}

func (x *SCSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SCSet) GetItems() []*SCSetItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_sc_pb_sc_proto protoreflect.FileDescriptor

var file_apps_sc_pb_sc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x63, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8b, 0x03, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x43, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x63, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52,
	0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x32, 0x0a, 0x14,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x61, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x11, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x35,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x69, 0x0a, 0x02, 0x53, 0x43, 0x12, 0x2e, 0x0a, 0x04, 0x4d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x02, 0x53,
	0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x02, 0x53, 0x43,
	0x22, 0x97, 0x03, 0x0a, 0x09, 0x53, 0x43, 0x53, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x73, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3f, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e,
	0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x32, 0x0a, 0x14, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x05, 0x53, 0x43,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x63, 0x2e, 0x53, 0x43,
	0x53, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6b,
	0x75, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_sc_pb_sc_proto_rawDescOnce sync.Once
	file_apps_sc_pb_sc_proto_rawDescData = file_apps_sc_pb_sc_proto_rawDesc
)

func file_apps_sc_pb_sc_proto_rawDescGZIP() []byte {
	file_apps_sc_pb_sc_proto_rawDescOnce.Do(func() {
		file_apps_sc_pb_sc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_sc_pb_sc_proto_rawDescData)
	})
	return file_apps_sc_pb_sc_proto_rawDescData
}

var file_apps_sc_pb_sc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_sc_pb_sc_proto_goTypes = []interface{}{
	(*CreateSCRequest)(nil), // 0: codehorse.mkube.sc.CreateSCRequest
	(*ListMapItem)(nil),     // 1: codehorse.mkube.sc.ListMapItem
	(*SC)(nil),              // 2: codehorse.mkube.sc.SC
	(*SCSetItem)(nil),       // 3: codehorse.mkube.sc.SCSetItem
	(*SCSet)(nil),           // 4: codehorse.mkube.sc.SCSet
	(*meta.Meta)(nil),       // 5: codehorse.mcube.meta.Meta
}
var file_apps_sc_pb_sc_proto_depIdxs = []int32{
	1, // 0: codehorse.mkube.sc.CreateSCRequest.Labels:type_name -> codehorse.mkube.sc.ListMapItem
	1, // 1: codehorse.mkube.sc.CreateSCRequest.Parameters:type_name -> codehorse.mkube.sc.ListMapItem
	5, // 2: codehorse.mkube.sc.SC.Meta:type_name -> codehorse.mcube.meta.Meta
	0, // 3: codehorse.mkube.sc.SC.SC:type_name -> codehorse.mkube.sc.CreateSCRequest
	1, // 4: codehorse.mkube.sc.SCSetItem.Labels:type_name -> codehorse.mkube.sc.ListMapItem
	1, // 5: codehorse.mkube.sc.SCSetItem.Parameters:type_name -> codehorse.mkube.sc.ListMapItem
	3, // 6: codehorse.mkube.sc.SCSet.Items:type_name -> codehorse.mkube.sc.SCSetItem
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_sc_pb_sc_proto_init() }
func file_apps_sc_pb_sc_proto_init() {
	if File_apps_sc_pb_sc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_sc_pb_sc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSCRequest); i {
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
		file_apps_sc_pb_sc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_sc_pb_sc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC); i {
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
		file_apps_sc_pb_sc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCSetItem); i {
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
		file_apps_sc_pb_sc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCSet); i {
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
			RawDescriptor: file_apps_sc_pb_sc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_sc_pb_sc_proto_goTypes,
		DependencyIndexes: file_apps_sc_pb_sc_proto_depIdxs,
		MessageInfos:      file_apps_sc_pb_sc_proto_msgTypes,
	}.Build()
	File_apps_sc_pb_sc_proto = out.File
	file_apps_sc_pb_sc_proto_rawDesc = nil
	file_apps_sc_pb_sc_proto_goTypes = nil
	file_apps_sc_pb_sc_proto_depIdxs = nil
}
