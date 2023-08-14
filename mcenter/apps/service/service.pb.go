// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: apps/service/pb/service.proto

package service

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

// CreateServiceRequest结构体
type CreateServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" bson:"name"`
}

func (x *CreateServiceRequest) Reset() {
	*x = CreateServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceRequest) ProtoMessage() {}

func (x *CreateServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServiceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateServiceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Service结构体
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateServiceRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// @gotags: bson:",inline" json:"credential"
	Credential *Credential `protobuf:"bytes,3,opt,name=credential,proto3" json:"credential" bson:",inline"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *Service) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Service) GetSpec() *CreateServiceRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Service) GetCredential() *Credential {
	if x != nil {
		return x.Credential
	}
	return nil
}

// Credential结构体
type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,1,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// @gotags: bson:"client_id" json:"client_id"
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id" bson:"client_id"`
	// @gotags: bson:"client_secret" json:"client_secret"
	ClientSecret string `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" bson:"client_secret"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{2}
}

func (x *Credential) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Credential) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Credential) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// ServiceSet结构体
type ServiceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Service `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ServiceSet) Reset() {
	*x = ServiceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSet) ProtoMessage() {}

func (x *ServiceSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSet.ProtoReflect.Descriptor instead.
func (*ServiceSet) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ServiceSet) GetItems() []*Service {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_service_pb_service_proto protoreflect.FileDescriptor

var file_apps_service_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x60, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x43, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x45, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x6b, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x5c, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_service_pb_service_proto_rawDescOnce sync.Once
	file_apps_service_pb_service_proto_rawDescData = file_apps_service_pb_service_proto_rawDesc
)

func file_apps_service_pb_service_proto_rawDescGZIP() []byte {
	file_apps_service_pb_service_proto_rawDescOnce.Do(func() {
		file_apps_service_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_service_pb_service_proto_rawDescData)
	})
	return file_apps_service_pb_service_proto_rawDescData
}

var file_apps_service_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_service_pb_service_proto_goTypes = []interface{}{
	(*CreateServiceRequest)(nil), // 0: codehorse.mcenter.service.CreateServiceRequest
	(*Service)(nil),              // 1: codehorse.mcenter.service.Service
	(*Credential)(nil),           // 2: codehorse.mcenter.service.Credential
	(*ServiceSet)(nil),           // 3: codehorse.mcenter.service.ServiceSet
	(*meta.Meta)(nil),            // 4: codehorse.mcube.meta.Meta
}
var file_apps_service_pb_service_proto_depIdxs = []int32{
	4, // 0: codehorse.mcenter.service.Service.meta:type_name -> codehorse.mcube.meta.Meta
	0, // 1: codehorse.mcenter.service.Service.spec:type_name -> codehorse.mcenter.service.CreateServiceRequest
	2, // 2: codehorse.mcenter.service.Service.credential:type_name -> codehorse.mcenter.service.Credential
	1, // 3: codehorse.mcenter.service.ServiceSet.items:type_name -> codehorse.mcenter.service.Service
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_service_pb_service_proto_init() }
func file_apps_service_pb_service_proto_init() {
	if File_apps_service_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_service_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceRequest); i {
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
		file_apps_service_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_apps_service_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_apps_service_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSet); i {
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
			RawDescriptor: file_apps_service_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_service_pb_service_proto_goTypes,
		DependencyIndexes: file_apps_service_pb_service_proto_depIdxs,
		MessageInfos:      file_apps_service_pb_service_proto_msgTypes,
	}.Build()
	File_apps_service_pb_service_proto = out.File
	file_apps_service_pb_service_proto_rawDesc = nil
	file_apps_service_pb_service_proto_goTypes = nil
	file_apps_service_pb_service_proto_depIdxs = nil
}
