// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: pkg/proto/admin.proto

package admin

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

type CheckIsAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *CheckIsAdminRequest) Reset() {
	*x = CheckIsAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIsAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIsAdminRequest) ProtoMessage() {}

func (x *CheckIsAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIsAdminRequest.ProtoReflect.Descriptor instead.
func (*CheckIsAdminRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_admin_proto_rawDescGZIP(), []int{0}
}

func (x *CheckIsAdminRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckIsAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAuth  bool `protobuf:"varint,1,opt,name=isAuth,proto3" json:"isAuth,omitempty"`
	IsAdmin bool `protobuf:"varint,2,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
}

func (x *CheckIsAdminResponse) Reset() {
	*x = CheckIsAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIsAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIsAdminResponse) ProtoMessage() {}

func (x *CheckIsAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIsAdminResponse.ProtoReflect.Descriptor instead.
func (*CheckIsAdminResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CheckIsAdminResponse) GetIsAuth() bool {
	if x != nil {
		return x.IsAuth
	}
	return false
}

func (x *CheckIsAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

var File_pkg_proto_admin_proto protoreflect.FileDescriptor

var file_pkg_proto_admin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x48, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x32, 0x4b,
	0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_admin_proto_rawDescOnce sync.Once
	file_pkg_proto_admin_proto_rawDescData = file_pkg_proto_admin_proto_rawDesc
)

func file_pkg_proto_admin_proto_rawDescGZIP() []byte {
	file_pkg_proto_admin_proto_rawDescOnce.Do(func() {
		file_pkg_proto_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_admin_proto_rawDescData)
	})
	return file_pkg_proto_admin_proto_rawDescData
}

var file_pkg_proto_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_admin_proto_goTypes = []interface{}{
	(*CheckIsAdminRequest)(nil),  // 0: CheckIsAdminRequest
	(*CheckIsAdminResponse)(nil), // 1: CheckIsAdminResponse
}
var file_pkg_proto_admin_proto_depIdxs = []int32{
	0, // 0: AdminService.CheckIsAdmin:input_type -> CheckIsAdminRequest
	1, // 1: AdminService.CheckIsAdmin:output_type -> CheckIsAdminResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_admin_proto_init() }
func file_pkg_proto_admin_proto_init() {
	if File_pkg_proto_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIsAdminRequest); i {
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
		file_pkg_proto_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIsAdminResponse); i {
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
			RawDescriptor: file_pkg_proto_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_admin_proto_goTypes,
		DependencyIndexes: file_pkg_proto_admin_proto_depIdxs,
		MessageInfos:      file_pkg_proto_admin_proto_msgTypes,
	}.Build()
	File_pkg_proto_admin_proto = out.File
	file_pkg_proto_admin_proto_rawDesc = nil
	file_pkg_proto_admin_proto_goTypes = nil
	file_pkg_proto_admin_proto_depIdxs = nil
}
