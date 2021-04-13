// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: ooshandler/v1/ooshandler.proto

package ooshandler

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/nlnwa/veidemann-api/go/frontier/v1"
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

type SubmitUriRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri *v1.QueuedUri `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *SubmitUriRequest) Reset() {
	*x = SubmitUriRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ooshandler_v1_ooshandler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitUriRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitUriRequest) ProtoMessage() {}

func (x *SubmitUriRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ooshandler_v1_ooshandler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitUriRequest.ProtoReflect.Descriptor instead.
func (*SubmitUriRequest) Descriptor() ([]byte, []int) {
	return file_ooshandler_v1_ooshandler_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitUriRequest) GetUri() *v1.QueuedUri {
	if x != nil {
		return x.Uri
	}
	return nil
}

var File_ooshandler_v1_ooshandler_proto protoreflect.FileDescriptor

var file_ooshandler_v1_ooshandler_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x6f, 0x73, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x6f, 0x73, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x6f, 0x73, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x55, 0x72, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x55, 0x72, 0x69, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x32, 0x62, 0x0a, 0x0a, 0x4f, 0x6f, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x54, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x55, 0x72, 0x69, 0x12, 0x2d,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x6f, 0x73, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x55, 0x72, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x78, 0x0a, 0x25, 0x6e, 0x6f, 0x2e, 0x6e, 0x62,
	0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x6f, 0x73, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x4f, 0x6f, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x6f, 0x73, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x6f, 0x73, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ooshandler_v1_ooshandler_proto_rawDescOnce sync.Once
	file_ooshandler_v1_ooshandler_proto_rawDescData = file_ooshandler_v1_ooshandler_proto_rawDesc
)

func file_ooshandler_v1_ooshandler_proto_rawDescGZIP() []byte {
	file_ooshandler_v1_ooshandler_proto_rawDescOnce.Do(func() {
		file_ooshandler_v1_ooshandler_proto_rawDescData = protoimpl.X.CompressGZIP(file_ooshandler_v1_ooshandler_proto_rawDescData)
	})
	return file_ooshandler_v1_ooshandler_proto_rawDescData
}

var file_ooshandler_v1_ooshandler_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ooshandler_v1_ooshandler_proto_goTypes = []interface{}{
	(*SubmitUriRequest)(nil), // 0: veidemann.api.ooshandler.v1.SubmitUriRequest
	(*v1.QueuedUri)(nil),     // 1: veidemann.api.frontier.v1.QueuedUri
	(*empty.Empty)(nil),      // 2: google.protobuf.Empty
}
var file_ooshandler_v1_ooshandler_proto_depIdxs = []int32{
	1, // 0: veidemann.api.ooshandler.v1.SubmitUriRequest.uri:type_name -> veidemann.api.frontier.v1.QueuedUri
	0, // 1: veidemann.api.ooshandler.v1.OosHandler.SubmitUri:input_type -> veidemann.api.ooshandler.v1.SubmitUriRequest
	2, // 2: veidemann.api.ooshandler.v1.OosHandler.SubmitUri:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ooshandler_v1_ooshandler_proto_init() }
func file_ooshandler_v1_ooshandler_proto_init() {
	if File_ooshandler_v1_ooshandler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ooshandler_v1_ooshandler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitUriRequest); i {
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
			RawDescriptor: file_ooshandler_v1_ooshandler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ooshandler_v1_ooshandler_proto_goTypes,
		DependencyIndexes: file_ooshandler_v1_ooshandler_proto_depIdxs,
		MessageInfos:      file_ooshandler_v1_ooshandler_proto_msgTypes,
	}.Build()
	File_ooshandler_v1_ooshandler_proto = out.File
	file_ooshandler_v1_ooshandler_proto_rawDesc = nil
	file_ooshandler_v1_ooshandler_proto_goTypes = nil
	file_ooshandler_v1_ooshandler_proto_depIdxs = nil
}