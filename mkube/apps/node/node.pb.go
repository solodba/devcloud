// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/node/pb/node.proto

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

// NodeSet结构体
type NodeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"Items" json:"Items"
	Items []*NodeSetItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items" bson:"Items"`
}

func (x *NodeSet) Reset() {
	*x = NodeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSet) ProtoMessage() {}

func (x *NodeSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSet.ProtoReflect.Descriptor instead.
func (*NodeSet) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_node_proto_rawDescGZIP(), []int{0}
}

func (x *NodeSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *NodeSet) GetItems() []*NodeSetItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// NodeSetItem结构体
type NodeSetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"status" json:"status"
	Status string `protobuf:"bytes,2,opt,name=Status,proto3" json:"status" bson:"status"`
	// @gotags: bson:"age" json:"age"
	Age int64 `protobuf:"varint,3,opt,name=Age,proto3" json:"age" bson:"age"`
	// @gotags: bson:"internalIp" json:"internalIp"
	InternalIp string `protobuf:"bytes,4,opt,name=InternalIp,proto3" json:"internalIp" bson:"internalIp"`
	// @gotags: bson:"externalIp" json:"externalIp"
	ExternalIp string `protobuf:"bytes,5,opt,name=ExternalIp,proto3" json:"externalIp" bson:"externalIp"`
	// @gotags: bson:"version" json:"version"
	Version string `protobuf:"bytes,6,opt,name=Version,proto3" json:"version" bson:"version"`
	// @gotags: bson:"osImage" json:"osImage"
	OsImage string `protobuf:"bytes,7,opt,name=OsImage,proto3" json:"osImage" bson:"osImage"`
	// @gotags: bson:"kernelVersion" json:"kernelVersion"
	KernelVersion string `protobuf:"bytes,8,opt,name=KernelVersion,proto3" json:"kernelVersion" bson:"kernelVersion"`
	// @gotags: bson:"containerRuntime" json:"containerRuntime"
	ContainerRuntime string `protobuf:"bytes,9,opt,name=ContainerRuntime,proto3" json:"containerRuntime" bson:"containerRuntime"`
	// @gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,10,rep,name=Labels,proto3" json:"labels" bson:"labels"`
	// @gotags: bson:"taints" json:"taints"
	Taints []*Taint `protobuf:"bytes,11,rep,name=Taints,proto3" json:"taints" bson:"taints"`
}

func (x *NodeSetItem) Reset() {
	*x = NodeSetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSetItem) ProtoMessage() {}

func (x *NodeSetItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSetItem.ProtoReflect.Descriptor instead.
func (*NodeSetItem) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_node_proto_rawDescGZIP(), []int{1}
}

func (x *NodeSetItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeSetItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NodeSetItem) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *NodeSetItem) GetInternalIp() string {
	if x != nil {
		return x.InternalIp
	}
	return ""
}

func (x *NodeSetItem) GetExternalIp() string {
	if x != nil {
		return x.ExternalIp
	}
	return ""
}

func (x *NodeSetItem) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeSetItem) GetOsImage() string {
	if x != nil {
		return x.OsImage
	}
	return ""
}

func (x *NodeSetItem) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *NodeSetItem) GetContainerRuntime() string {
	if x != nil {
		return x.ContainerRuntime
	}
	return ""
}

func (x *NodeSetItem) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *NodeSetItem) GetTaints() []*Taint {
	if x != nil {
		return x.Taints
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
		mi := &file_apps_node_pb_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMapItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMapItem) ProtoMessage() {}

func (x *ListMapItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_node_proto_msgTypes[2]
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
	return file_apps_node_pb_node_proto_rawDescGZIP(), []int{2}
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

// Taint结构体
type Taint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"key" json:"key"
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"key" bson:"key"`
	// @gotags: bson:"value" json:"value"
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"value" bson:"value"`
	// @gotags: bson:"effect" json:"effect"
	Effect string `protobuf:"bytes,3,opt,name=Effect,proto3" json:"effect" bson:"effect"`
}

func (x *Taint) Reset() {
	*x = Taint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_node_pb_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Taint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Taint) ProtoMessage() {}

func (x *Taint) ProtoReflect() protoreflect.Message {
	mi := &file_apps_node_pb_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Taint.ProtoReflect.Descriptor instead.
func (*Taint) Descriptor() ([]byte, []int) {
	return file_apps_node_pb_node_proto_rawDescGZIP(), []int{3}
}

func (x *Taint) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Taint) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Taint) GetEffect() string {
	if x != nil {
		return x.Effect
	}
	return ""
}

var File_apps_node_pb_node_proto protoreflect.FileDescriptor

var file_apps_node_pb_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x22,
	0x58, 0x0a, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x37, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x81, 0x03, 0x0a, 0x0b, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4f, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x4b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x54, 0x61, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x54, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x35, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x47, 0x0a, 0x05, 0x54, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6b, 0x75,
	0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_node_pb_node_proto_rawDescOnce sync.Once
	file_apps_node_pb_node_proto_rawDescData = file_apps_node_pb_node_proto_rawDesc
)

func file_apps_node_pb_node_proto_rawDescGZIP() []byte {
	file_apps_node_pb_node_proto_rawDescOnce.Do(func() {
		file_apps_node_pb_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_node_pb_node_proto_rawDescData)
	})
	return file_apps_node_pb_node_proto_rawDescData
}

var file_apps_node_pb_node_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_node_pb_node_proto_goTypes = []interface{}{
	(*NodeSet)(nil),     // 0: codehorse.mkube.node.NodeSet
	(*NodeSetItem)(nil), // 1: codehorse.mkube.node.NodeSetItem
	(*ListMapItem)(nil), // 2: codehorse.mkube.node.ListMapItem
	(*Taint)(nil),       // 3: codehorse.mkube.node.Taint
}
var file_apps_node_pb_node_proto_depIdxs = []int32{
	1, // 0: codehorse.mkube.node.NodeSet.Items:type_name -> codehorse.mkube.node.NodeSetItem
	2, // 1: codehorse.mkube.node.NodeSetItem.Labels:type_name -> codehorse.mkube.node.ListMapItem
	3, // 2: codehorse.mkube.node.NodeSetItem.Taints:type_name -> codehorse.mkube.node.Taint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_node_pb_node_proto_init() }
func file_apps_node_pb_node_proto_init() {
	if File_apps_node_pb_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_node_pb_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSet); i {
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
		file_apps_node_pb_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSetItem); i {
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
		file_apps_node_pb_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_node_pb_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Taint); i {
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
			RawDescriptor: file_apps_node_pb_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_node_pb_node_proto_goTypes,
		DependencyIndexes: file_apps_node_pb_node_proto_depIdxs,
		MessageInfos:      file_apps_node_pb_node_proto_msgTypes,
	}.Build()
	File_apps_node_pb_node_proto = out.File
	file_apps_node_pb_node_proto_rawDesc = nil
	file_apps_node_pb_node_proto_goTypes = nil
	file_apps_node_pb_node_proto_depIdxs = nil
}
