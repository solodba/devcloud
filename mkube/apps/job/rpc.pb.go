// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/job/pb/rpc.proto

package job

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

// DeleteJobRequest结构体
type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteJobRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpdateJobRequest结构体
type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"job" json:"job"
	Job *CreateJobRequest `protobuf:"bytes,1,opt,name=Job,proto3" json:"job" bson:"job"`
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateJobRequest) GetJob() *CreateJobRequest {
	if x != nil {
		return x.Job
	}
	return nil
}

// QueryJobRequest结构体
type QueryJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"keyword" json:"keyword"
	Keyword string `protobuf:"bytes,2,opt,name=Keyword,proto3" json:"keyword" bson:"keyword"`
}

func (x *QueryJobRequest) Reset() {
	*x = QueryJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryJobRequest) ProtoMessage() {}

func (x *QueryJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryJobRequest.ProtoReflect.Descriptor instead.
func (*QueryJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryJobRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryJobRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// DescribeJobRequest结构体
type DescribeJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name" bson:"name"`
}

func (x *DescribeJobRequest) Reset() {
	*x = DescribeJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeJobRequest) ProtoMessage() {}

func (x *DescribeJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeJobRequest.ProtoReflect.Descriptor instead.
func (*DescribeJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeJobRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DescribeJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_apps_job_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_job_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x1a, 0x15, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x03, 0x4a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x03, 0x4a, 0x6f, 0x62, 0x22, 0x49, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x46, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xaa, 0x03, 0x0a, 0x03, 0x52, 0x50,
	0x43, 0x12, 0x4c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x25,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73,
	0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x12,
	0x59, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x25, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e,
	0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x12, 0x4d, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4a, 0x6f, 0x62, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65,
	0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74, 0x12, 0x5d, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_job_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_job_pb_rpc_proto_rawDescData = file_apps_job_pb_rpc_proto_rawDesc
)

func file_apps_job_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_job_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_job_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_job_pb_rpc_proto_rawDescData)
	})
	return file_apps_job_pb_rpc_proto_rawDescData
}

var file_apps_job_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_job_pb_rpc_proto_goTypes = []interface{}{
	(*DeleteJobRequest)(nil),   // 0: codehorse.mkube.job.DeleteJobRequest
	(*UpdateJobRequest)(nil),   // 1: codehorse.mkube.job.UpdateJobRequest
	(*QueryJobRequest)(nil),    // 2: codehorse.mkube.job.QueryJobRequest
	(*DescribeJobRequest)(nil), // 3: codehorse.mkube.job.DescribeJobRequest
	(*CreateJobRequest)(nil),   // 4: codehorse.mkube.job.CreateJobRequest
	(*Job)(nil),                // 5: codehorse.mkube.job.Job
	(*JobSet)(nil),             // 6: codehorse.mkube.job.JobSet
}
var file_apps_job_pb_rpc_proto_depIdxs = []int32{
	4, // 0: codehorse.mkube.job.UpdateJobRequest.Job:type_name -> codehorse.mkube.job.CreateJobRequest
	4, // 1: codehorse.mkube.job.RPC.CreateJob:input_type -> codehorse.mkube.job.CreateJobRequest
	0, // 2: codehorse.mkube.job.RPC.DeleteJob:input_type -> codehorse.mkube.job.DeleteJobRequest
	1, // 3: codehorse.mkube.job.RPC.UpdateJob:input_type -> codehorse.mkube.job.UpdateJobRequest
	2, // 4: codehorse.mkube.job.RPC.QueryJob:input_type -> codehorse.mkube.job.QueryJobRequest
	3, // 5: codehorse.mkube.job.RPC.DescribeJob:input_type -> codehorse.mkube.job.DescribeJobRequest
	5, // 6: codehorse.mkube.job.RPC.CreateJob:output_type -> codehorse.mkube.job.Job
	4, // 7: codehorse.mkube.job.RPC.DeleteJob:output_type -> codehorse.mkube.job.CreateJobRequest
	5, // 8: codehorse.mkube.job.RPC.UpdateJob:output_type -> codehorse.mkube.job.Job
	6, // 9: codehorse.mkube.job.RPC.QueryJob:output_type -> codehorse.mkube.job.JobSet
	4, // 10: codehorse.mkube.job.RPC.DescribeJob:output_type -> codehorse.mkube.job.CreateJobRequest
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_job_pb_rpc_proto_init() }
func file_apps_job_pb_rpc_proto_init() {
	if File_apps_job_pb_rpc_proto != nil {
		return
	}
	file_apps_job_pb_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_job_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobRequest); i {
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
		file_apps_job_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobRequest); i {
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
		file_apps_job_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryJobRequest); i {
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
		file_apps_job_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeJobRequest); i {
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
			RawDescriptor: file_apps_job_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_job_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_job_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_job_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_job_pb_rpc_proto = out.File
	file_apps_job_pb_rpc_proto_rawDesc = nil
	file_apps_job_pb_rpc_proto_goTypes = nil
	file_apps_job_pb_rpc_proto_depIdxs = nil
}
