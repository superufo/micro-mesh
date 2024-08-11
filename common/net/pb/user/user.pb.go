// https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
// protoc --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative  user.proto
// protoc --go_out=. --go_opt=paths=source_relative user.proto

// protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. user.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: user.proto

package user

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

// 服务器登录请求信息
type UserLoginTos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	OrgCode  *string `protobuf:"bytes,3,opt,name=org_code,json=orgCode,proto3,oneof" json:"org_code,omitempty"`
	Type     *uint32 `protobuf:"varint,4,opt,name=type,proto3,oneof" json:"type,omitempty"`
}

func (x *UserLoginTos) Reset() {
	*x = UserLoginTos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginTos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginTos) ProtoMessage() {}

func (x *UserLoginTos) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginTos.ProtoReflect.Descriptor instead.
func (*UserLoginTos) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginTos) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserLoginTos) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserLoginTos) GetOrgCode() string {
	if x != nil && x.OrgCode != nil {
		return *x.OrgCode
	}
	return ""
}

func (x *UserLoginTos) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

// 回复客户端的登录信息
type DataToc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Time  string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *DataToc) Reset() {
	*x = DataToc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataToc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataToc) ProtoMessage() {}

func (x *DataToc) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataToc.ProtoReflect.Descriptor instead.
func (*DataToc) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *DataToc) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DataToc) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type ReplyToc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *DataToc `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReplyToc) Reset() {
	*x = ReplyToc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyToc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyToc) ProtoMessage() {}

func (x *ReplyToc) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyToc.ProtoReflect.Descriptor instead.
func (*ReplyToc) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyToc) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ReplyToc) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ReplyToc) GetData() *DataToc {
	if x != nil {
		return x.Data
	}
	return nil
}

type VerifyTokenTos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyTokenTos) Reset() {
	*x = VerifyTokenTos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenTos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenTos) ProtoMessage() {}

func (x *VerifyTokenTos) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenTos.ProtoReflect.Descriptor instead.
func (*VerifyTokenTos) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyTokenTos) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenToc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data bool   `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *VerifyTokenToc) Reset() {
	*x = VerifyTokenToc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenToc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenToc) ProtoMessage() {}

func (x *VerifyTokenToc) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenToc.ProtoReflect.Descriptor instead.
func (*VerifyTokenToc) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyTokenToc) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *VerifyTokenToc) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *VerifyTokenToc) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x74, 0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x6f,
	0x72, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x6f, 0x72, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x74, 0x6f, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x55, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x6f,
	0x63, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x4c, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x74, 0x6f, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x7c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x73, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x63, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x74, 0x6f, 0x73, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x6f, 0x63, 0x22, 0x00, 0x42, 0x24, 0x5a,
	0x22, 0x7a, 0x68, 0x75, 0x67, 0x65, 0x64, 0x61, 0x6f, 0x6a, 0x69, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_proto_goTypes = []interface{}{
	(*UserLoginTos)(nil),   // 0: user.user_login_tos
	(*DataToc)(nil),        // 1: user.data_toc
	(*ReplyToc)(nil),       // 2: user.reply_toc
	(*VerifyTokenTos)(nil), // 3: user.verify_token_tos
	(*VerifyTokenToc)(nil), // 4: user.verify_token_toc
}
var file_user_proto_depIdxs = []int32{
	1, // 0: user.reply_toc.data:type_name -> user.data_toc
	0, // 1: user.User.GetToken:input_type -> user.user_login_tos
	3, // 2: user.User.VerifyToken:input_type -> user.verify_token_tos
	2, // 3: user.User.GetToken:output_type -> user.reply_toc
	4, // 4: user.User.VerifyToken:output_type -> user.verify_token_toc
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginTos); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataToc); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyToc); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenTos); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenToc); i {
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
	file_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
