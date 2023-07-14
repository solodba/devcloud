// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/pod/pb/rpc.proto

package pod

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

// 删除Pod请求
type DeletePodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
}

func (x *DeletePodRequest) Reset() {
	*x = DeletePodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pod_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePodRequest) ProtoMessage() {}

func (x *DeletePodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pod_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePodRequest.ProtoReflect.Descriptor instead.
func (*DeletePodRequest) Descriptor() ([]byte, []int) {
	return file_apps_pod_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DeletePodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeletePodRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// 更新Pod请求
type UpdatePodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"pod" json:"pod"
	Pod *Pod `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod" bson:"pod"`
}

func (x *UpdatePodRequest) Reset() {
	*x = UpdatePodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pod_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePodRequest) ProtoMessage() {}

func (x *UpdatePodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pod_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePodRequest.ProtoReflect.Descriptor instead.
func (*UpdatePodRequest) Descriptor() ([]byte, []int) {
	return file_apps_pod_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePodRequest) GetPod() *Pod {
	if x != nil {
		return x.Pod
	}
	return nil
}

// 查询Pod请求
type QueryPodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"nodeName" json:"nodeName"
	NodeName string `protobuf:"bytes,1,opt,name=NodeName,proto3" json:"nodeName" bson:"nodeName"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"keyword" json:"keyword"
	Keyword string `protobuf:"bytes,3,opt,name=Keyword,proto3" json:"keyword" bson:"keyword"`
}

func (x *QueryPodRequest) Reset() {
	*x = QueryPodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pod_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPodRequest) ProtoMessage() {}

func (x *QueryPodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pod_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPodRequest.ProtoReflect.Descriptor instead.
func (*QueryPodRequest) Descriptor() ([]byte, []int) {
	return file_apps_pod_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryPodRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *QueryPodRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryPodRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// 查询Pod详情请求
type DescribePodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
}

func (x *DescribePodRequest) Reset() {
	*x = DescribePodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_pod_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePodRequest) ProtoMessage() {}

func (x *DescribePodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_pod_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePodRequest.ProtoReflect.Descriptor instead.
func (*DescribePodRequest) Descriptor() ([]byte, []int) {
	return file_apps_pod_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribePodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribePodRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_apps_pod_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_pod_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x1a, 0x15, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x70, 0x6f, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64,
	0x2e, 0x50, 0x6f, 0x64, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x22, 0x65, 0x0a, 0x0f, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x46, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0x90, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43,
	0x12, 0x4c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x25, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x70, 0x6f, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x12, 0x4c,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f,
	0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x12, 0x4c, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x12, 0x4d, 0x0a, 0x08, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x6f, 0x64, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70,
	0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x65, 0x74, 0x12, 0x50, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62,
	0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x70, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_pod_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_pod_pb_rpc_proto_rawDescData = file_apps_pod_pb_rpc_proto_rawDesc
)

func file_apps_pod_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_pod_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_pod_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_pod_pb_rpc_proto_rawDescData)
	})
	return file_apps_pod_pb_rpc_proto_rawDescData
}

var file_apps_pod_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_pod_pb_rpc_proto_goTypes = []interface{}{
	(*DeletePodRequest)(nil),   // 0: codehorse.mpaas.pod.DeletePodRequest
	(*UpdatePodRequest)(nil),   // 1: codehorse.mpaas.pod.UpdatePodRequest
	(*QueryPodRequest)(nil),    // 2: codehorse.mpaas.pod.QueryPodRequest
	(*DescribePodRequest)(nil), // 3: codehorse.mpaas.pod.DescribePodRequest
	(*Pod)(nil),                // 4: codehorse.mpaas.pod.Pod
	(*CreatePodRequest)(nil),   // 5: codehorse.mpaas.pod.CreatePodRequest
	(*PodSet)(nil),             // 6: codehorse.mpaas.pod.PodSet
}
var file_apps_pod_pb_rpc_proto_depIdxs = []int32{
	4, // 0: codehorse.mpaas.pod.UpdatePodRequest.pod:type_name -> codehorse.mpaas.pod.Pod
	5, // 1: codehorse.mpaas.pod.RPC.CreatePod:input_type -> codehorse.mpaas.pod.CreatePodRequest
	0, // 2: codehorse.mpaas.pod.RPC.DeletePod:input_type -> codehorse.mpaas.pod.DeletePodRequest
	1, // 3: codehorse.mpaas.pod.RPC.UpdatePod:input_type -> codehorse.mpaas.pod.UpdatePodRequest
	2, // 4: codehorse.mpaas.pod.RPC.QueryPod:input_type -> codehorse.mpaas.pod.QueryPodRequest
	3, // 5: codehorse.mpaas.pod.RPC.DescribePod:input_type -> codehorse.mpaas.pod.DescribePodRequest
	4, // 6: codehorse.mpaas.pod.RPC.CreatePod:output_type -> codehorse.mpaas.pod.Pod
	4, // 7: codehorse.mpaas.pod.RPC.DeletePod:output_type -> codehorse.mpaas.pod.Pod
	4, // 8: codehorse.mpaas.pod.RPC.UpdatePod:output_type -> codehorse.mpaas.pod.Pod
	6, // 9: codehorse.mpaas.pod.RPC.QueryPod:output_type -> codehorse.mpaas.pod.PodSet
	4, // 10: codehorse.mpaas.pod.RPC.DescribePod:output_type -> codehorse.mpaas.pod.Pod
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_pod_pb_rpc_proto_init() }
func file_apps_pod_pb_rpc_proto_init() {
	if File_apps_pod_pb_rpc_proto != nil {
		return
	}
	file_apps_pod_pb_pod_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_pod_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePodRequest); i {
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
		file_apps_pod_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePodRequest); i {
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
		file_apps_pod_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPodRequest); i {
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
		file_apps_pod_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePodRequest); i {
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
			RawDescriptor: file_apps_pod_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_pod_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_pod_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_pod_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_pod_pb_rpc_proto = out.File
	file_apps_pod_pb_rpc_proto_rawDesc = nil
	file_apps_pod_pb_rpc_proto_goTypes = nil
	file_apps_pod_pb_rpc_proto_depIdxs = nil
}
