// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: proto/streaming.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamServerTimeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CurrentTime   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamServerTimeResponse) Reset() {
	*x = StreamServerTimeResponse{}
	mi := &file_proto_streaming_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamServerTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamServerTimeResponse) ProtoMessage() {}

func (x *StreamServerTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_streaming_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamServerTimeResponse.ProtoReflect.Descriptor instead.
func (*StreamServerTimeResponse) Descriptor() ([]byte, []int) {
	return file_proto_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *StreamServerTimeResponse) GetCurrentTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CurrentTime
	}
	return nil
}

type StreamServerTimeRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	IntervalSeconds int32                  `protobuf:"varint,1,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StreamServerTimeRequest) Reset() {
	*x = StreamServerTimeRequest{}
	mi := &file_proto_streaming_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamServerTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamServerTimeRequest) ProtoMessage() {}

func (x *StreamServerTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_streaming_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamServerTimeRequest.ProtoReflect.Descriptor instead.
func (*StreamServerTimeRequest) Descriptor() ([]byte, []int) {
	return file_proto_streaming_proto_rawDescGZIP(), []int{1}
}

func (x *StreamServerTimeRequest) GetIntervalSeconds() int32 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

var File_proto_streaming_proto protoreflect.FileDescriptor

var file_proto_streaming_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x44,
	0x0a, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x32, 0x71, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x6d, 0x69, 0x74, 0x30, 0x30, 0x37, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_streaming_proto_rawDescOnce sync.Once
	file_proto_streaming_proto_rawDescData = file_proto_streaming_proto_rawDesc
)

func file_proto_streaming_proto_rawDescGZIP() []byte {
	file_proto_streaming_proto_rawDescOnce.Do(func() {
		file_proto_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_streaming_proto_rawDescData)
	})
	return file_proto_streaming_proto_rawDescData
}

var file_proto_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_streaming_proto_goTypes = []any{
	(*StreamServerTimeResponse)(nil), // 0: streaming.StreamServerTimeResponse
	(*StreamServerTimeRequest)(nil),  // 1: streaming.StreamServerTimeRequest
	(*timestamppb.Timestamp)(nil),    // 2: google.protobuf.Timestamp
}
var file_proto_streaming_proto_depIdxs = []int32{
	2, // 0: streaming.StreamServerTimeResponse.current_time:type_name -> google.protobuf.Timestamp
	1, // 1: streaming.StreamingService.StreamServerTime:input_type -> streaming.StreamServerTimeRequest
	0, // 2: streaming.StreamingService.StreamServerTime:output_type -> streaming.StreamServerTimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_streaming_proto_init() }
func file_proto_streaming_proto_init() {
	if File_proto_streaming_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_streaming_proto_goTypes,
		DependencyIndexes: file_proto_streaming_proto_depIdxs,
		MessageInfos:      file_proto_streaming_proto_msgTypes,
	}.Build()
	File_proto_streaming_proto = out.File
	file_proto_streaming_proto_rawDesc = nil
	file_proto_streaming_proto_goTypes = nil
	file_proto_streaming_proto_depIdxs = nil
}
