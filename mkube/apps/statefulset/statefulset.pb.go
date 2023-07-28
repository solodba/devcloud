// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/statefulset/pb/statefulset.proto

package statefulset

import (
	pod "github.com/solodba/devcloud/mkube/apps/pod"
	pvc "github.com/solodba/devcloud/mkube/apps/pvc"
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

// CreateStatefulSetRequest结构体
type CreateStatefulSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"base" json:"base"
	Base *Base `protobuf:"bytes,1,opt,name=Base,proto3" json:"base" bson:"base"`
	// @gotags: bson:"template" json:"template"
	Template *pod.Pod `protobuf:"bytes,2,opt,name=Template,proto3" json:"template" bson:"template"`
}

func (x *CreateStatefulSetRequest) Reset() {
	*x = CreateStatefulSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStatefulSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStatefulSetRequest) ProtoMessage() {}

func (x *CreateStatefulSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStatefulSetRequest.ProtoReflect.Descriptor instead.
func (*CreateStatefulSetRequest) Descriptor() ([]byte, []int) {
	return file_apps_statefulset_pb_statefulset_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStatefulSetRequest) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CreateStatefulSetRequest) GetTemplate() *pod.Pod {
	if x != nil {
		return x.Template
	}
	return nil
}

// Base结构体
type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,3,rep,name=Labels,proto3" json:"labels" bson:"labels"`
	// @gotags: bson:"replicas" json:"replicas"
	Replicas int32 `protobuf:"varint,4,opt,name=Replicas,proto3" json:"replicas" bson:"replicas"`
	// @gotags: bson:"selector" json:"selector"
	Selector []*ListMapItem `protobuf:"bytes,5,rep,name=Selector,proto3" json:"selector" bson:"selector"`
	// @gotags: bson:"serviceName" json:"serviceName"
	ServiceName string `protobuf:"bytes,6,opt,name=ServiceName,proto3" json:"serviceName" bson:"serviceName"`
	// @gotags: bson:"volumeClaimTemplates" json:"volumeClaimTemplates"
	VolumeClaimTemplates []*pvc.CreatePVCRequest `protobuf:"bytes,7,rep,name=VolumeClaimTemplates,proto3" json:"volumeClaimTemplates" bson:"volumeClaimTemplates"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_apps_statefulset_pb_statefulset_proto_rawDescGZIP(), []int{1}
}

func (x *Base) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Base) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Base) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Base) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Base) GetSelector() []*ListMapItem {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *Base) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Base) GetVolumeClaimTemplates() []*pvc.CreatePVCRequest {
	if x != nil {
		return x.VolumeClaimTemplates
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
		mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMapItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMapItem) ProtoMessage() {}

func (x *ListMapItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[2]
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
	return file_apps_statefulset_pb_statefulset_proto_rawDescGZIP(), []int{2}
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

// StatefulSet结构体
type StatefulSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"StatefulSet"
	StatefulSet *CreateStatefulSetRequest `protobuf:"bytes,2,opt,name=StatefulSet,proto3" json:"StatefulSet" bson:",inline"`
}

func (x *StatefulSet) Reset() {
	*x = StatefulSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatefulSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSet) ProtoMessage() {}

func (x *StatefulSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSet.ProtoReflect.Descriptor instead.
func (*StatefulSet) Descriptor() ([]byte, []int) {
	return file_apps_statefulset_pb_statefulset_proto_rawDescGZIP(), []int{3}
}

func (x *StatefulSet) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *StatefulSet) GetStatefulSet() *CreateStatefulSetRequest {
	if x != nil {
		return x.StatefulSet
	}
	return nil
}

// StatefulSetSetItem结构体
type StatefulSetSetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"ready" json:"ready"
	Ready int32 `protobuf:"varint,3,opt,name=Ready,proto3" json:"ready" bson:"ready"`
	// @gotags: bson:"replicas" json:"replicas"
	Replicas int32 `protobuf:"varint,4,opt,name=Replicas,proto3" json:"replicas" bson:"replicas"`
	// @gotags: bson:"age" json:"age"
	Age int64 `protobuf:"varint,5,opt,name=Age,proto3" json:"age" bson:"age"`
}

func (x *StatefulSetSetItem) Reset() {
	*x = StatefulSetSetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatefulSetSetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSetSetItem) ProtoMessage() {}

func (x *StatefulSetSetItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSetSetItem.ProtoReflect.Descriptor instead.
func (*StatefulSetSetItem) Descriptor() ([]byte, []int) {
	return file_apps_statefulset_pb_statefulset_proto_rawDescGZIP(), []int{4}
}

func (x *StatefulSetSetItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatefulSetSetItem) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *StatefulSetSetItem) GetReady() int32 {
	if x != nil {
		return x.Ready
	}
	return 0
}

func (x *StatefulSetSetItem) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *StatefulSetSetItem) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

// StatefulSetSet结构体
type StatefulSetSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*StatefulSetSetItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"items" bson:"items"`
}

func (x *StatefulSetSet) Reset() {
	*x = StatefulSetSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatefulSetSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSetSet) ProtoMessage() {}

func (x *StatefulSetSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_statefulset_pb_statefulset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSetSet.ProtoReflect.Descriptor instead.
func (*StatefulSetSet) Descriptor() ([]byte, []int) {
	return file_apps_statefulset_pb_statefulset_proto_rawDescGZIP(), []int{5}
}

func (x *StatefulSetSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *StatefulSetSet) GetItems() []*StatefulSetSetItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_statefulset_pb_statefulset_proto protoreflect.FileDescriptor

var file_apps_statefulset_pb_statefulset_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73,
	0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75,
	0x6c, 0x73, 0x65, 0x74, 0x1a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x64, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x76, 0x63, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73,
	0x65, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0xd9, 0x02, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x40,
	0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x44, 0x0a, 0x08,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x76, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x56, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x14, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22,
	0x35, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x66, 0x75, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x57, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6c, 0x53, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x74, 0x22,
	0x8a, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x22, 0x6d, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x45, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e,
	0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62,
	0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6b, 0x75, 0x62, 0x65,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x73, 0x65,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_statefulset_pb_statefulset_proto_rawDescOnce sync.Once
	file_apps_statefulset_pb_statefulset_proto_rawDescData = file_apps_statefulset_pb_statefulset_proto_rawDesc
)

func file_apps_statefulset_pb_statefulset_proto_rawDescGZIP() []byte {
	file_apps_statefulset_pb_statefulset_proto_rawDescOnce.Do(func() {
		file_apps_statefulset_pb_statefulset_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_statefulset_pb_statefulset_proto_rawDescData)
	})
	return file_apps_statefulset_pb_statefulset_proto_rawDescData
}

var file_apps_statefulset_pb_statefulset_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_statefulset_pb_statefulset_proto_goTypes = []interface{}{
	(*CreateStatefulSetRequest)(nil), // 0: codehorse.mkube.statefulset.CreateStatefulSetRequest
	(*Base)(nil),                     // 1: codehorse.mkube.statefulset.Base
	(*ListMapItem)(nil),              // 2: codehorse.mkube.statefulset.ListMapItem
	(*StatefulSet)(nil),              // 3: codehorse.mkube.statefulset.StatefulSet
	(*StatefulSetSetItem)(nil),       // 4: codehorse.mkube.statefulset.StatefulSetSetItem
	(*StatefulSetSet)(nil),           // 5: codehorse.mkube.statefulset.StatefulSetSet
	(*pod.Pod)(nil),                  // 6: codehorse.mkube.pod.Pod
	(*pvc.CreatePVCRequest)(nil),     // 7: codehorse.mkube.pvc.CreatePVCRequest
	(*meta.Meta)(nil),                // 8: codehorse.mcube.meta.Meta
}
var file_apps_statefulset_pb_statefulset_proto_depIdxs = []int32{
	1, // 0: codehorse.mkube.statefulset.CreateStatefulSetRequest.Base:type_name -> codehorse.mkube.statefulset.Base
	6, // 1: codehorse.mkube.statefulset.CreateStatefulSetRequest.Template:type_name -> codehorse.mkube.pod.Pod
	2, // 2: codehorse.mkube.statefulset.Base.Labels:type_name -> codehorse.mkube.statefulset.ListMapItem
	2, // 3: codehorse.mkube.statefulset.Base.Selector:type_name -> codehorse.mkube.statefulset.ListMapItem
	7, // 4: codehorse.mkube.statefulset.Base.VolumeClaimTemplates:type_name -> codehorse.mkube.pvc.CreatePVCRequest
	8, // 5: codehorse.mkube.statefulset.StatefulSet.Meta:type_name -> codehorse.mcube.meta.Meta
	0, // 6: codehorse.mkube.statefulset.StatefulSet.StatefulSet:type_name -> codehorse.mkube.statefulset.CreateStatefulSetRequest
	4, // 7: codehorse.mkube.statefulset.StatefulSetSet.Items:type_name -> codehorse.mkube.statefulset.StatefulSetSetItem
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_apps_statefulset_pb_statefulset_proto_init() }
func file_apps_statefulset_pb_statefulset_proto_init() {
	if File_apps_statefulset_pb_statefulset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_statefulset_pb_statefulset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStatefulSetRequest); i {
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
		file_apps_statefulset_pb_statefulset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_apps_statefulset_pb_statefulset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_statefulset_pb_statefulset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatefulSet); i {
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
		file_apps_statefulset_pb_statefulset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatefulSetSetItem); i {
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
		file_apps_statefulset_pb_statefulset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatefulSetSet); i {
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
			RawDescriptor: file_apps_statefulset_pb_statefulset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_statefulset_pb_statefulset_proto_goTypes,
		DependencyIndexes: file_apps_statefulset_pb_statefulset_proto_depIdxs,
		MessageInfos:      file_apps_statefulset_pb_statefulset_proto_msgTypes,
	}.Build()
	File_apps_statefulset_pb_statefulset_proto = out.File
	file_apps_statefulset_pb_statefulset_proto_rawDesc = nil
	file_apps_statefulset_pb_statefulset_proto_goTypes = nil
	file_apps_statefulset_pb_statefulset_proto_depIdxs = nil
}
