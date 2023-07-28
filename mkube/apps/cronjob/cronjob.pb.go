// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/cronjob/pb/cronjob.proto

package cronjob

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

// CreateCronJobRequest结构体
type CreateCronJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"base" json:"base"
	Base *CronJobBase `protobuf:"bytes,1,opt,name=Base,proto3" json:"base" bson:"base"`
	// @gotags: bson:"template" json:"template"
	Template *pod.Pod `protobuf:"bytes,2,opt,name=Template,proto3" json:"template" bson:"template"`
}

func (x *CreateCronJobRequest) Reset() {
	*x = CreateCronJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCronJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCronJobRequest) ProtoMessage() {}

func (x *CreateCronJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCronJobRequest.ProtoReflect.Descriptor instead.
func (*CreateCronJobRequest) Descriptor() ([]byte, []int) {
	return file_apps_cronjob_pb_cronjob_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCronJobRequest) GetBase() *CronJobBase {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CreateCronJobRequest) GetTemplate() *pod.Pod {
	if x != nil {
		return x.Template
	}
	return nil
}

// JobBase结构体
type JobBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"completions" json:"completions"
	Completions int32 `protobuf:"varint,1,opt,name=Completions,proto3" json:"completions" bson:"completions"`
	// @gotags: bson:"backoffLimit" json:"backoffLimit"
	BackoffLimit int32 `protobuf:"varint,2,opt,name=BackoffLimit,proto3" json:"backoffLimit" bson:"backoffLimit"`
}

func (x *JobBase) Reset() {
	*x = JobBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobBase) ProtoMessage() {}

func (x *JobBase) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[1]
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
	return file_apps_cronjob_pb_cronjob_proto_rawDescGZIP(), []int{1}
}

func (x *JobBase) GetCompletions() int32 {
	if x != nil {
		return x.Completions
	}
	return 0
}

func (x *JobBase) GetBackoffLimit() int32 {
	if x != nil {
		return x.BackoffLimit
	}
	return 0
}

// CronJobBase结构体
type CronJobBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"labels" json:"labels"
	Labels []*ListMapItem `protobuf:"bytes,3,rep,name=Labels,proto3" json:"labels" bson:"labels"`
	// @gotags: bson:"schedule" json:"schedule"
	Schedule string `protobuf:"bytes,4,opt,name=Schedule,proto3" json:"schedule" bson:"schedule"`
	// @gotags: bson:"suspend" json:"suspend"
	Suspend bool `protobuf:"varint,5,opt,name=Suspend,proto3" json:"suspend" bson:"suspend"`
	// @gotags: bson:"concurrencyPolicy" json:"concurrencyPolicy"
	ConcurrencyPolicy string `protobuf:"bytes,6,opt,name=ConcurrencyPolicy,proto3" json:"concurrencyPolicy" bson:"concurrencyPolicy"`
	// @gotags: bson:"successfulJobsHistoryLimit" json:"successfulJobsHistoryLimit"
	SuccessfulJobsHistoryLimit int32 `protobuf:"varint,7,opt,name=SuccessfulJobsHistoryLimit,proto3" json:"successfulJobsHistoryLimit" bson:"successfulJobsHistoryLimit"`
	// @gotags: bson:"failedJobsHistoryLimit" json:"failedJobsHistoryLimit"
	FailedJobsHistoryLimit int32 `protobuf:"varint,8,opt,name=FailedJobsHistoryLimit,proto3" json:"failedJobsHistoryLimit" bson:"failedJobsHistoryLimit"`
	// @gotags: bson:"jobBase" json:"jobBase"
	JobBase *JobBase `protobuf:"bytes,9,opt,name=JobBase,proto3" json:"jobBase" bson:"jobBase"`
}

func (x *CronJobBase) Reset() {
	*x = CronJobBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronJobBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronJobBase) ProtoMessage() {}

func (x *CronJobBase) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronJobBase.ProtoReflect.Descriptor instead.
func (*CronJobBase) Descriptor() ([]byte, []int) {
	return file_apps_cronjob_pb_cronjob_proto_rawDescGZIP(), []int{2}
}

func (x *CronJobBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CronJobBase) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CronJobBase) GetLabels() []*ListMapItem {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CronJobBase) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *CronJobBase) GetSuspend() bool {
	if x != nil {
		return x.Suspend
	}
	return false
}

func (x *CronJobBase) GetConcurrencyPolicy() string {
	if x != nil {
		return x.ConcurrencyPolicy
	}
	return ""
}

func (x *CronJobBase) GetSuccessfulJobsHistoryLimit() int32 {
	if x != nil {
		return x.SuccessfulJobsHistoryLimit
	}
	return 0
}

func (x *CronJobBase) GetFailedJobsHistoryLimit() int32 {
	if x != nil {
		return x.FailedJobsHistoryLimit
	}
	return 0
}

func (x *CronJobBase) GetJobBase() *JobBase {
	if x != nil {
		return x.JobBase
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
		mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMapItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMapItem) ProtoMessage() {}

func (x *ListMapItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[3]
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
	return file_apps_cronjob_pb_cronjob_proto_rawDescGZIP(), []int{3}
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

// CronJob结构体
type CronJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=Meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"cronjob"
	CronJob *CreateCronJobRequest `protobuf:"bytes,2,opt,name=CronJob,proto3" json:"cronjob" bson:",inline"`
}

func (x *CronJob) Reset() {
	*x = CronJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronJob) ProtoMessage() {}

func (x *CronJob) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronJob.ProtoReflect.Descriptor instead.
func (*CronJob) Descriptor() ([]byte, []int) {
	return file_apps_cronjob_pb_cronjob_proto_rawDescGZIP(), []int{4}
}

func (x *CronJob) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CronJob) GetCronJob() *CreateCronJobRequest {
	if x != nil {
		return x.CronJob
	}
	return nil
}

// CronJobSetItem结构体
type CronJobSetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name" bson:"name"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=Namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"schedule" json:"schedule"
	Schedule string `protobuf:"bytes,3,opt,name=Schedule,proto3" json:"schedule" bson:"schedule"`
	// @gotags: bson:"suspend" json:"suspend"
	Suspend bool `protobuf:"varint,4,opt,name=Suspend,proto3" json:"suspend" bson:"suspend"`
	// @gotags: bson:"active" json:"active"
	Active int64 `protobuf:"varint,5,opt,name=Active,proto3" json:"active" bson:"active"`
	// @gotags: bson:"lastSchedule" json:"lastSchedule"
	LastSchedule int64 `protobuf:"varint,6,opt,name=LastSchedule,proto3" json:"lastSchedule" bson:"lastSchedule"`
	// @gotags: bson:"age" json:"age"
	Age int64 `protobuf:"varint,7,opt,name=Age,proto3" json:"age" bson:"age"`
}

func (x *CronJobSetItem) Reset() {
	*x = CronJobSetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronJobSetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronJobSetItem) ProtoMessage() {}

func (x *CronJobSetItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronJobSetItem.ProtoReflect.Descriptor instead.
func (*CronJobSetItem) Descriptor() ([]byte, []int) {
	return file_apps_cronjob_pb_cronjob_proto_rawDescGZIP(), []int{5}
}

func (x *CronJobSetItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CronJobSetItem) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CronJobSetItem) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *CronJobSetItem) GetSuspend() bool {
	if x != nil {
		return x.Suspend
	}
	return false
}

func (x *CronJobSetItem) GetActive() int64 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *CronJobSetItem) GetLastSchedule() int64 {
	if x != nil {
		return x.LastSchedule
	}
	return 0
}

func (x *CronJobSetItem) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

// CronJobSet结构体
type CronJobSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*CronJobSetItem `protobuf:"bytes,2,rep,name=Items,proto3" json:"items" bson:"items"`
}

func (x *CronJobSet) Reset() {
	*x = CronJobSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronJobSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronJobSet) ProtoMessage() {}

func (x *CronJobSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_cronjob_pb_cronjob_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronJobSet.ProtoReflect.Descriptor instead.
func (*CronJobSet) Descriptor() ([]byte, []int) {
	return file_apps_cronjob_pb_cronjob_proto_rawDescGZIP(), []int{6}
}

func (x *CronJobSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CronJobSet) GetItems() []*CronJobSetItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_cronjob_pb_cronjob_proto protoreflect.FileDescriptor

var file_apps_cronjob_pb_cronjob_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65,
	0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x1a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70,
	0x6f, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e,
	0x63, 0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x6f,
	0x64, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22,
	0x4f, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x42, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x95, 0x03, 0x0a, 0x0b, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x42, 0x61, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53,
	0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x3e, 0x0a, 0x1a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66,
	0x75, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4a, 0x6f,
	0x62, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x73,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3a, 0x0a, 0x07,
	0x4a, 0x6f, 0x62, 0x42, 0x61, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e,
	0x63, 0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x07, 0x4a, 0x6f, 0x62, 0x42, 0x61, 0x73, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x61, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x82, 0x01, 0x0a, 0x07, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x2e, 0x0a, 0x04, 0x4d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x07, 0x43,
	0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63,
	0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x43, 0x72, 0x6f,
	0x6e, 0x4a, 0x6f, 0x62, 0x22, 0xc6, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x53, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x4c,
	0x61, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x22, 0x61, 0x0a,
	0x0a, 0x43, 0x72, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x3d, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x53, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6d, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x6a,
	0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_cronjob_pb_cronjob_proto_rawDescOnce sync.Once
	file_apps_cronjob_pb_cronjob_proto_rawDescData = file_apps_cronjob_pb_cronjob_proto_rawDesc
)

func file_apps_cronjob_pb_cronjob_proto_rawDescGZIP() []byte {
	file_apps_cronjob_pb_cronjob_proto_rawDescOnce.Do(func() {
		file_apps_cronjob_pb_cronjob_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_cronjob_pb_cronjob_proto_rawDescData)
	})
	return file_apps_cronjob_pb_cronjob_proto_rawDescData
}

var file_apps_cronjob_pb_cronjob_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_cronjob_pb_cronjob_proto_goTypes = []interface{}{
	(*CreateCronJobRequest)(nil), // 0: codehorse.mkube.cronjob.CreateCronJobRequest
	(*JobBase)(nil),              // 1: codehorse.mkube.cronjob.JobBase
	(*CronJobBase)(nil),          // 2: codehorse.mkube.cronjob.CronJobBase
	(*ListMapItem)(nil),          // 3: codehorse.mkube.cronjob.ListMapItem
	(*CronJob)(nil),              // 4: codehorse.mkube.cronjob.CronJob
	(*CronJobSetItem)(nil),       // 5: codehorse.mkube.cronjob.CronJobSetItem
	(*CronJobSet)(nil),           // 6: codehorse.mkube.cronjob.CronJobSet
	(*pod.Pod)(nil),              // 7: codehorse.mkube.pod.Pod
	(*meta.Meta)(nil),            // 8: codehorse.mcube.meta.Meta
}
var file_apps_cronjob_pb_cronjob_proto_depIdxs = []int32{
	2, // 0: codehorse.mkube.cronjob.CreateCronJobRequest.Base:type_name -> codehorse.mkube.cronjob.CronJobBase
	7, // 1: codehorse.mkube.cronjob.CreateCronJobRequest.Template:type_name -> codehorse.mkube.pod.Pod
	3, // 2: codehorse.mkube.cronjob.CronJobBase.Labels:type_name -> codehorse.mkube.cronjob.ListMapItem
	1, // 3: codehorse.mkube.cronjob.CronJobBase.JobBase:type_name -> codehorse.mkube.cronjob.JobBase
	8, // 4: codehorse.mkube.cronjob.CronJob.Meta:type_name -> codehorse.mcube.meta.Meta
	0, // 5: codehorse.mkube.cronjob.CronJob.CronJob:type_name -> codehorse.mkube.cronjob.CreateCronJobRequest
	5, // 6: codehorse.mkube.cronjob.CronJobSet.Items:type_name -> codehorse.mkube.cronjob.CronJobSetItem
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_cronjob_pb_cronjob_proto_init() }
func file_apps_cronjob_pb_cronjob_proto_init() {
	if File_apps_cronjob_pb_cronjob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_cronjob_pb_cronjob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCronJobRequest); i {
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
		file_apps_cronjob_pb_cronjob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_cronjob_pb_cronjob_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronJobBase); i {
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
		file_apps_cronjob_pb_cronjob_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_apps_cronjob_pb_cronjob_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronJob); i {
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
		file_apps_cronjob_pb_cronjob_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronJobSetItem); i {
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
		file_apps_cronjob_pb_cronjob_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronJobSet); i {
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
			RawDescriptor: file_apps_cronjob_pb_cronjob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_cronjob_pb_cronjob_proto_goTypes,
		DependencyIndexes: file_apps_cronjob_pb_cronjob_proto_depIdxs,
		MessageInfos:      file_apps_cronjob_pb_cronjob_proto_msgTypes,
	}.Build()
	File_apps_cronjob_pb_cronjob_proto = out.File
	file_apps_cronjob_pb_cronjob_proto_rawDesc = nil
	file_apps_cronjob_pb_cronjob_proto_goTypes = nil
	file_apps_cronjob_pb_cronjob_proto_depIdxs = nil
}
