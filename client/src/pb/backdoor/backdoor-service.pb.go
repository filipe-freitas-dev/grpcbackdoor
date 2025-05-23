// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.28.3
// source: src/proto/backdoor-service.proto

package backdoor

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

type Out struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Output        string                 `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Out) Reset() {
	*x = Out{}
	mi := &file_src_proto_backdoor_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Out) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Out) ProtoMessage() {}

func (x *Out) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_backdoor_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Out.ProtoReflect.Descriptor instead.
func (*Out) Descriptor() ([]byte, []int) {
	return file_src_proto_backdoor_service_proto_rawDescGZIP(), []int{0}
}

func (x *Out) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_src_proto_backdoor_service_proto protoreflect.FileDescriptor

var file_src_proto_backdoor_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x64, 0x6f, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x6f, 0x6f, 0x72, 0x22, 0x1d, 0x0a, 0x03,
	0x4f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x41, 0x0a, 0x0f, 0x42,
	0x61, 0x63, 0x6b, 0x64, 0x6f, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x78, 0x65, 0x63, 0x12, 0x0d, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x64, 0x6f, 0x6f, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x64, 0x6f, 0x6f, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13,
	0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x64,
	0x6f, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_backdoor_service_proto_rawDescOnce sync.Once
	file_src_proto_backdoor_service_proto_rawDescData = file_src_proto_backdoor_service_proto_rawDesc
)

func file_src_proto_backdoor_service_proto_rawDescGZIP() []byte {
	file_src_proto_backdoor_service_proto_rawDescOnce.Do(func() {
		file_src_proto_backdoor_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_backdoor_service_proto_rawDescData)
	})
	return file_src_proto_backdoor_service_proto_rawDescData
}

var file_src_proto_backdoor_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_src_proto_backdoor_service_proto_goTypes = []any{
	(*Out)(nil), // 0: backdoor.Out
}
var file_src_proto_backdoor_service_proto_depIdxs = []int32{
	0, // 0: backdoor.BackdoorService.RemoteExec:input_type -> backdoor.Out
	0, // 1: backdoor.BackdoorService.RemoteExec:output_type -> backdoor.Out
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_proto_backdoor_service_proto_init() }
func file_src_proto_backdoor_service_proto_init() {
	if File_src_proto_backdoor_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_proto_backdoor_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_backdoor_service_proto_goTypes,
		DependencyIndexes: file_src_proto_backdoor_service_proto_depIdxs,
		MessageInfos:      file_src_proto_backdoor_service_proto_msgTypes,
	}.Build()
	File_src_proto_backdoor_service_proto = out.File
	file_src_proto_backdoor_service_proto_rawDesc = nil
	file_src_proto_backdoor_service_proto_goTypes = nil
	file_src_proto_backdoor_service_proto_depIdxs = nil
}
