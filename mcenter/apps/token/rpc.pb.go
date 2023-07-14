// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/token/pb/rpc.proto

package token

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

// 颁发令牌请求
type IssueTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"grant_type" json:"grant_type"
	GrantType GRANT_TYPE `protobuf:"varint,1,opt,name=grant_type,json=grantType,proto3,enum=codehorse.mcenter.token.GRANT_TYPE" json:"grant_type" bson:"grant_type"`
	// @gotags: bson:"username" json:"username" validate:"required"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" bson:"username" validate:"required"`
	// @gotags: bson:"password" json:"password" validate:"required"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" bson:"password" validate:"required"`
}

func (x *IssueTokenRequest) Reset() {
	*x = IssueTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenRequest) ProtoMessage() {}

func (x *IssueTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenRequest.ProtoReflect.Descriptor instead.
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *IssueTokenRequest) GetGrantType() GRANT_TYPE {
	if x != nil {
		return x.GrantType
	}
	return GRANT_TYPE_PASSWORD
}

func (x *IssueTokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IssueTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 校验令牌请求
type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"access_token" json:"access_token" validate:"required"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"access_token" validate:"required"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_apps_token_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_token_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x64, 0x65, 0x68,
	0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70,
	0x62, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01,
	0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f,
	0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x09, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x39, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x65, 0x0a, 0x03, 0x52, 0x50,
	0x43, 0x12, 0x5e, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x64, 0x62, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_token_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_token_pb_rpc_proto_rawDescData = file_apps_token_pb_rpc_proto_rawDesc
)

func file_apps_token_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_token_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_token_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_token_pb_rpc_proto_rawDescData)
	})
	return file_apps_token_pb_rpc_proto_rawDescData
}

var file_apps_token_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_token_pb_rpc_proto_goTypes = []interface{}{
	(*IssueTokenRequest)(nil),    // 0: codehorse.mcenter.token.IssueTokenRequest
	(*ValidateTokenRequest)(nil), // 1: codehorse.mcenter.token.ValidateTokenRequest
	(GRANT_TYPE)(0),              // 2: codehorse.mcenter.token.GRANT_TYPE
	(*Token)(nil),                // 3: codehorse.mcenter.token.Token
}
var file_apps_token_pb_rpc_proto_depIdxs = []int32{
	2, // 0: codehorse.mcenter.token.IssueTokenRequest.grant_type:type_name -> codehorse.mcenter.token.GRANT_TYPE
	1, // 1: codehorse.mcenter.token.RPC.ValidateToken:input_type -> codehorse.mcenter.token.ValidateTokenRequest
	3, // 2: codehorse.mcenter.token.RPC.ValidateToken:output_type -> codehorse.mcenter.token.Token
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_token_pb_rpc_proto_init() }
func file_apps_token_pb_rpc_proto_init() {
	if File_apps_token_pb_rpc_proto != nil {
		return
	}
	file_apps_token_pb_token_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_token_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueTokenRequest); i {
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
		file_apps_token_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
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
			RawDescriptor: file_apps_token_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_token_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_token_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_token_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_token_pb_rpc_proto = out.File
	file_apps_token_pb_rpc_proto_rawDesc = nil
	file_apps_token_pb_rpc_proto_goTypes = nil
	file_apps_token_pb_rpc_proto_depIdxs = nil
}
