// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: transmitter.proto

package transmitter

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

type TransmitterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId         string  `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ExpectedValue     int32   `protobuf:"varint,2,opt,name=expectedValue,proto3" json:"expectedValue,omitempty"`
	StandartDiviation float64 `protobuf:"fixed64,3,opt,name=standartDiviation,proto3" json:"standartDiviation,omitempty"`
}

func (x *TransmitterRequest) Reset() {
	*x = TransmitterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transmitter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitterRequest) ProtoMessage() {}

func (x *TransmitterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitterRequest.ProtoReflect.Descriptor instead.
func (*TransmitterRequest) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{0}
}

func (x *TransmitterRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *TransmitterRequest) GetExpectedValue() int32 {
	if x != nil {
		return x.ExpectedValue
	}
	return 0
}

func (x *TransmitterRequest) GetStandartDiviation() float64 {
	if x != nil {
		return x.StandartDiviation
	}
	return 0
}

type TransmitterResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency float64                `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TransmitterResponce) Reset() {
	*x = TransmitterResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transmitter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitterResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitterResponce) ProtoMessage() {}

func (x *TransmitterResponce) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitterResponce.ProtoReflect.Descriptor instead.
func (*TransmitterResponce) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{1}
}

func (x *TransmitterResponce) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *TransmitterResponce) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *TransmitterResponce) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_transmitter_proto protoreflect.FileDescriptor

var file_transmitter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x44, 0x69, 0x76, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x74, 0x44, 0x69, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x13,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x63, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transmitter_proto_rawDescOnce sync.Once
	file_transmitter_proto_rawDescData = file_transmitter_proto_rawDesc
)

func file_transmitter_proto_rawDescGZIP() []byte {
	file_transmitter_proto_rawDescOnce.Do(func() {
		file_transmitter_proto_rawDescData = protoimpl.X.CompressGZIP(file_transmitter_proto_rawDescData)
	})
	return file_transmitter_proto_rawDescData
}

var file_transmitter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transmitter_proto_goTypes = []interface{}{
	(*TransmitterRequest)(nil),    // 0: transmitter.TransmitterRequest
	(*TransmitterResponce)(nil),   // 1: transmitter.TransmitterResponce
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_transmitter_proto_depIdxs = []int32{
	2, // 0: transmitter.TransmitterResponce.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: transmitter.Transmitter.Transmitter:input_type -> transmitter.TransmitterRequest
	1, // 2: transmitter.Transmitter.Transmitter:output_type -> transmitter.TransmitterResponce
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transmitter_proto_init() }
func file_transmitter_proto_init() {
	if File_transmitter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transmitter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitterRequest); i {
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
		file_transmitter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransmitterResponce); i {
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
			RawDescriptor: file_transmitter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transmitter_proto_goTypes,
		DependencyIndexes: file_transmitter_proto_depIdxs,
		MessageInfos:      file_transmitter_proto_msgTypes,
	}.Build()
	File_transmitter_proto = out.File
	file_transmitter_proto_rawDesc = nil
	file_transmitter_proto_goTypes = nil
	file_transmitter_proto_depIdxs = nil
}
