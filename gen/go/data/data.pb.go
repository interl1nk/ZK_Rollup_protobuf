// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: data/data.proto

package protobufv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gen           string                 `protobuf:"bytes,1,opt,name=gen,proto3" json:"gen,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenRequest) Reset() {
	*x = GenRequest{}
	mi := &file_data_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenRequest) ProtoMessage() {}

func (x *GenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenRequest.ProtoReflect.Descriptor instead.
func (*GenRequest) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *GenRequest) GetGen() string {
	if x != nil {
		return x.Gen
	}
	return ""
}

type GenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gen           string                 `protobuf:"bytes,1,opt,name=gen,proto3" json:"gen,omitempty"`
	Valid         bool                   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenResponse) Reset() {
	*x = GenResponse{}
	mi := &file_data_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenResponse) ProtoMessage() {}

func (x *GenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenResponse.ProtoReflect.Descriptor instead.
func (*GenResponse) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{1}
}

func (x *GenResponse) GetGen() string {
	if x != nil {
		return x.Gen
	}
	return ""
}

func (x *GenResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_data_data_proto protoreflect.FileDescriptor

var file_data_data_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x22, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0x39,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x65, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6c, 0x31, 0x6e, 0x6b, 0x2e, 0x7a, 0x6b, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x75, 0x70,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_data_data_proto_rawDescOnce sync.Once
	file_data_data_proto_rawDescData []byte
)

func file_data_data_proto_rawDescGZIP() []byte {
	file_data_data_proto_rawDescOnce.Do(func() {
		file_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_data_data_proto_rawDesc), len(file_data_data_proto_rawDesc)))
	})
	return file_data_data_proto_rawDescData
}

var file_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_data_data_proto_goTypes = []any{
	(*GenRequest)(nil),  // 0: data.GenRequest
	(*GenResponse)(nil), // 1: data.GenResponse
}
var file_data_data_proto_depIdxs = []int32{
	0, // 0: data.Data.GetGenData:input_type -> data.GenRequest
	1, // 1: data.Data.GetGenData:output_type -> data.GenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_data_proto_init() }
func file_data_data_proto_init() {
	if File_data_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_data_data_proto_rawDesc), len(file_data_data_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_data_proto_goTypes,
		DependencyIndexes: file_data_data_proto_depIdxs,
		MessageInfos:      file_data_data_proto_msgTypes,
	}.Build()
	File_data_data_proto = out.File
	file_data_data_proto_goTypes = nil
	file_data_data_proto_depIdxs = nil
}
