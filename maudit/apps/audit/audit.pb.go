// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: apps/audit/pb/audit.proto

package audit

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

// AuditLog结构体
type AuditLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// @gotags: bson:"username" json:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" bson:"username"`
	// @gotags: bson:"time" json:"time"
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time" bson:"time"`
	// @gotags: bson:"service_id" json:"service_id"
	ServiceId string `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id"`
	// @gotags: bson:"operate" json:"operate"
	Operate string `protobuf:"bytes,5,opt,name=operate,proto3" json:"operate" bson:"operate"`
	// @gotags: bson:"request" json:"request"
	Request string `protobuf:"bytes,6,opt,name=request,proto3" json:"request" bson:"request"`
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_audit_pb_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_apps_audit_pb_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_apps_audit_pb_audit_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuditLog) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuditLog) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AuditLog) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *AuditLog) GetOperate() string {
	if x != nil {
		return x.Operate
	}
	return ""
}

func (x *AuditLog) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

// AuditLogSet结构体
type AuditLogSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=Total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*AuditLog `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *AuditLogSet) Reset() {
	*x = AuditLogSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_audit_pb_audit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLogSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogSet) ProtoMessage() {}

func (x *AuditLogSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_audit_pb_audit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogSet.ProtoReflect.Descriptor instead.
func (*AuditLogSet) Descriptor() ([]byte, []int) {
	return file_apps_audit_pb_audit_proto_rawDescGZIP(), []int{1}
}

func (x *AuditLogSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AuditLogSet) GetItems() []*AuditLog {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_audit_pb_audit_proto protoreflect.FileDescriptor

var file_apps_audit_pb_audit_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x64,
	0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x2e, 0x6d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_audit_pb_audit_proto_rawDescOnce sync.Once
	file_apps_audit_pb_audit_proto_rawDescData = file_apps_audit_pb_audit_proto_rawDesc
)

func file_apps_audit_pb_audit_proto_rawDescGZIP() []byte {
	file_apps_audit_pb_audit_proto_rawDescOnce.Do(func() {
		file_apps_audit_pb_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_audit_pb_audit_proto_rawDescData)
	})
	return file_apps_audit_pb_audit_proto_rawDescData
}

var file_apps_audit_pb_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_audit_pb_audit_proto_goTypes = []interface{}{
	(*AuditLog)(nil),    // 0: codehorse.maudit.audit.AuditLog
	(*AuditLogSet)(nil), // 1: codehorse.maudit.audit.AuditLogSet
}
var file_apps_audit_pb_audit_proto_depIdxs = []int32{
	0, // 0: codehorse.maudit.audit.AuditLogSet.items:type_name -> codehorse.maudit.audit.AuditLog
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_audit_pb_audit_proto_init() }
func file_apps_audit_pb_audit_proto_init() {
	if File_apps_audit_pb_audit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_audit_pb_audit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLog); i {
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
		file_apps_audit_pb_audit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLogSet); i {
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
			RawDescriptor: file_apps_audit_pb_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_audit_pb_audit_proto_goTypes,
		DependencyIndexes: file_apps_audit_pb_audit_proto_depIdxs,
		MessageInfos:      file_apps_audit_pb_audit_proto_msgTypes,
	}.Build()
	File_apps_audit_pb_audit_proto = out.File
	file_apps_audit_pb_audit_proto_rawDesc = nil
	file_apps_audit_pb_audit_proto_goTypes = nil
	file_apps_audit_pb_audit_proto_depIdxs = nil
}
