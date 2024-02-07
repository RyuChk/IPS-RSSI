// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ips/map/v1/presence.proto

package mapv1

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

type FetchOnlineUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchOnlineUserRequest) Reset() {
	*x = FetchOnlineUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_presence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchOnlineUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchOnlineUserRequest) ProtoMessage() {}

func (x *FetchOnlineUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_presence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchOnlineUserRequest.ProtoReflect.Descriptor instead.
func (*FetchOnlineUserRequest) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_presence_proto_rawDescGZIP(), []int{0}
}

type FetchOnlineUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchOnlineUserResponse) Reset() {
	*x = FetchOnlineUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_map_v1_presence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchOnlineUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchOnlineUserResponse) ProtoMessage() {}

func (x *FetchOnlineUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_map_v1_presence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchOnlineUserResponse.ProtoReflect.Descriptor instead.
func (*FetchOnlineUserResponse) Descriptor() ([]byte, []int) {
	return file_ips_map_v1_presence_proto_rawDescGZIP(), []int{1}
}

var File_ips_map_v1_presence_proto protoreflect.FileDescriptor

var file_ips_map_v1_presence_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x70, 0x73,
	0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x18, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x19, 0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6d, 0x0a, 0x0f,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5a, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x22, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb5, 0x01, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x42, 0x6f, 0x6e, 0x65, 0x2f, 0x69, 0x70, 0x73, 0x2d, 0x72, 0x73, 0x73, 0x69, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x73, 0x2f, 0x6d,
	0x61, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x4d,
	0x58, 0xaa, 0x02, 0x0a, 0x49, 0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0a, 0x49, 0x70, 0x73, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x49, 0x70,
	0x73, 0x5c, 0x4d, 0x61, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x49, 0x70, 0x73, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ips_map_v1_presence_proto_rawDescOnce sync.Once
	file_ips_map_v1_presence_proto_rawDescData = file_ips_map_v1_presence_proto_rawDesc
)

func file_ips_map_v1_presence_proto_rawDescGZIP() []byte {
	file_ips_map_v1_presence_proto_rawDescOnce.Do(func() {
		file_ips_map_v1_presence_proto_rawDescData = protoimpl.X.CompressGZIP(file_ips_map_v1_presence_proto_rawDescData)
	})
	return file_ips_map_v1_presence_proto_rawDescData
}

var file_ips_map_v1_presence_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ips_map_v1_presence_proto_goTypes = []interface{}{
	(*FetchOnlineUserRequest)(nil),  // 0: ips.map.v1.FetchOnlineUserRequest
	(*FetchOnlineUserResponse)(nil), // 1: ips.map.v1.FetchOnlineUserResponse
}
var file_ips_map_v1_presence_proto_depIdxs = []int32{
	0, // 0: ips.map.v1.PresenceService.FetchOnlineUser:input_type -> ips.map.v1.FetchOnlineUserRequest
	1, // 1: ips.map.v1.PresenceService.FetchOnlineUser:output_type -> ips.map.v1.FetchOnlineUserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ips_map_v1_presence_proto_init() }
func file_ips_map_v1_presence_proto_init() {
	if File_ips_map_v1_presence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ips_map_v1_presence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchOnlineUserRequest); i {
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
		file_ips_map_v1_presence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchOnlineUserResponse); i {
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
			RawDescriptor: file_ips_map_v1_presence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ips_map_v1_presence_proto_goTypes,
		DependencyIndexes: file_ips_map_v1_presence_proto_depIdxs,
		MessageInfos:      file_ips_map_v1_presence_proto_msgTypes,
	}.Build()
	File_ips_map_v1_presence_proto = out.File
	file_ips_map_v1_presence_proto_rawDesc = nil
	file_ips_map_v1_presence_proto_goTypes = nil
	file_ips_map_v1_presence_proto_depIdxs = nil
}