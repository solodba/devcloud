// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/job/pb/job.proto

package job

import (
	pod "github.com/solodba/devcloud/mkube/apps/pod"
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

// JobBase结构体
type JobBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,3,rep,name=Labels,proto3" json:"labels" bson:"labels"`
	// @gotags: bson:"completions" json:"completions"
	Completions int32 `protobuf:"varint,4,opt,name=Completions,proto3" json:"completions" bson:"completions"`
}

func (x *JobBase) Reset() {
	*x = JobBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobBase) ProtoMessage() {}

func (x *JobBase) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobBase.ProtoReflect.Descriptor instead.
func (*JobBase) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobBase) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *JobBase) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *JobBase) GetCompletions() int32 {
	if x != nil {
		return x.Completions
	}
	return 0
}

// CreateJobRequest结构体
type CreateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"base" json:"base"
	Base *JobBase `protobuf:"bytes,1,opt,name=Base,proto3" json:"base" bson:"base"`
	// @gotags: bson:"template" json:"template"
	Template *pod.Pod `protobuf:"bytes,2,opt,name=Template,proto3" json:"template" bson:"template"`
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_job_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobRequest) GetBase() *JobBase {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CreateJobRequest) GetTemplate() *pod.Pod {
	if x != nil {
		return x.Template
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
		mi := &file_apps_job_pb_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMapItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMapItem) ProtoMessage() {}

func (x *ListMapItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_job_proto_msgTypes[2]
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
	return file_apps_job_pb_job_proto_rawDescGZIP(), []int{2}
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

// Job结构体
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"job"
	Job *CreateJobRequest `protobuf:"bytes,2,opt,name=Job,proto3" json:"job" bson:",inline"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_job_proto_rawDescGZIP(), []int{3}
}

func (x *Job) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Job) GetJob() *CreateJobRequest {
	if x != nil {
		return x.Job
	}
	return nil
}

// JobSetItem结构体
type JobSetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"completions" json:"completions"
	Completions int32 `protobuf:"varint,3,opt,name=Completions,proto3" json:"completions" bson:"completions"`
	// @gotags: bson:"succeeded" json:"succeeded"
	Succeeded int32 `protobuf:"varint,4,opt,name=Succeeded,proto3" json:"succeeded" bson:"succeeded"`
	// @gotags: bson:"duration" json:"duration"
	Duration int64 `protobuf:"varint,5,opt,name=Duration,proto3" json:"duration" bson:"duration"`
	// @gotags: bson:"age" json:"age"
	Age int64 `protobuf:"varint,6,opt,name=Age,proto3" json:"age" bson:"age"`
}

func (x *JobSetItem) Reset() {
	*x = JobSetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSetItem) ProtoMessage() {}

func (x *JobSetItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSetItem.ProtoReflect.Descriptor instead.
func (*JobSetItem) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_job_proto_rawDescGZIP(), []int{4}
}

func (x *JobSetItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobSetItem) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *JobSetItem) GetCompletions() int32 {
	if x != nil {
		return x.Completions
	}
	return 0
}

func (x *JobSetItem) GetSucceeded() int32 {
	if x != nil {
		return x.Succeeded
	}
	return 0
}

func (x *JobSetItem) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *JobSetItem) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

// JobSet结构体
type JobSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*JobSetItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"items" bson:"items"`
}

func (x *JobSet) Reset() {
	*x = JobSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_job_pb_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSet) ProtoMessage() {}

func (x *JobSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_job_pb_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSet.ProtoReflect.Descriptor instead.
func (*JobSet) Descriptor() ([]byte, []int) {
	return file_apps_job_pb_job_proto_rawDescGZIP(), []int{5}
}

func (x *JobSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *JobSet) GetItems() []*JobSetItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_job_pb_job_proto protoreflect.FileDescriptor

var file_apps_job_pb_job_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x1a, 0x15, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x70, 0x6f, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62,
	0x61, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x4a,
	0x6f, 0x62, 0x42, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7a, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72,
	0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x22, 0x35, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6e, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x2e,
	0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x37,
	0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x03, 0x4a, 0x6f, 0x62, 0x22, 0xac, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73,
	0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6b, 0x75,
	0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_job_pb_job_proto_rawDescOnce sync.Once
	file_apps_job_pb_job_proto_rawDescData = file_apps_job_pb_job_proto_rawDesc
)

func file_apps_job_pb_job_proto_rawDescGZIP() []byte {
	file_apps_job_pb_job_proto_rawDescOnce.Do(func() {
		file_apps_job_pb_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_job_pb_job_proto_rawDescData)
	})
	return file_apps_job_pb_job_proto_rawDescData
}

var file_apps_job_pb_job_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_job_pb_job_proto_goTypes = []interface{}{
	(*JobBase)(nil),          // 0: codehorse.mkube.job.JobBase
	(*CreateJobRequest)(nil), // 1: codehorse.mkube.job.CreateJobRequest
	(*ListMapItem)(nil),      // 2: codehorse.mkube.job.ListMapItem
	(*Job)(nil),              // 3: codehorse.mkube.job.Job
	(*JobSetItem)(nil),       // 4: codehorse.mkube.job.JobSetItem
	(*JobSet)(nil),           // 5: codehorse.mkube.job.JobSet
	(*pod.Pod)(nil),          // 6: codehorse.mkube.pod.Pod
	(*meta.Meta)(nil),        // 7: codehorse.mcube.meta.Meta
}
var file_apps_job_pb_job_proto_depIdxs = []int32{
	2, // 0: codehorse.mkube.job.JobBase.Labels:type_name -> codehorse.mkube.job.ListMapItem
	0, // 1: codehorse.mkube.job.CreateJobRequest.Base:type_name -> codehorse.mkube.job.JobBase
	6, // 2: codehorse.mkube.job.CreateJobRequest.Template:type_name -> codehorse.mkube.pod.Pod
	7, // 3: codehorse.mkube.job.Job.Meta:type_name -> codehorse.mcube.meta.Meta
	1, // 4: codehorse.mkube.job.Job.Job:type_name -> codehorse.mkube.job.CreateJobRequest
	4, // 5: codehorse.mkube.job.JobSet.Items:type_name -> codehorse.mkube.job.JobSetItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apps_job_pb_job_proto_init() }
func file_apps_job_pb_job_proto_init() {
	if File_apps_job_pb_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_job_pb_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobBase); i {
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
		file_apps_job_pb_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobRequest); i {
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
		file_apps_job_pb_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_job_pb_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_apps_job_pb_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSetItem); i {
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
		file_apps_job_pb_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSet); i {
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
			RawDescriptor: file_apps_job_pb_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_job_pb_job_proto_goTypes,
		DependencyIndexes: file_apps_job_pb_job_proto_depIdxs,
		MessageInfos:      file_apps_job_pb_job_proto_msgTypes,
	}.Build()
	File_apps_job_pb_job_proto = out.File
	file_apps_job_pb_job_proto_rawDesc = nil
	file_apps_job_pb_job_proto_goTypes = nil
	file_apps_job_pb_job_proto_depIdxs = nil
}
