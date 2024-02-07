// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ips/model/v1/model.proto

package modelv1

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

type PredictCoordinateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strength []float64 `protobuf:"fixed64,1,rep,packed,name=strength,proto3" json:"strength,omitempty"`
}

func (x *PredictCoordinateRequest) Reset() {
	*x = PredictCoordinateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_model_v1_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictCoordinateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictCoordinateRequest) ProtoMessage() {}

func (x *PredictCoordinateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_model_v1_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictCoordinateRequest.ProtoReflect.Descriptor instead.
func (*PredictCoordinateRequest) Descriptor() ([]byte, []int) {
	return file_ips_model_v1_model_proto_rawDescGZIP(), []int{0}
}

func (x *PredictCoordinateRequest) GetStrength() []float64 {
	if x != nil {
		return x.Strength
	}
	return nil
}

type PredictCoordinateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *PredictCoordinateResponse) Reset() {
	*x = PredictCoordinateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_model_v1_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictCoordinateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictCoordinateResponse) ProtoMessage() {}

func (x *PredictCoordinateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_model_v1_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictCoordinateResponse.ProtoReflect.Descriptor instead.
func (*PredictCoordinateResponse) Descriptor() ([]byte, []int) {
	return file_ips_model_v1_model_proto_rawDescGZIP(), []int{1}
}

func (x *PredictCoordinateResponse) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PredictCoordinateResponse) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PredictCoordinateResponse) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_ips_model_v1_model_proto protoreflect.FileDescriptor

var file_ips_model_v1_model_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x70, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x70, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x36, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x22, 0x45, 0x0a, 0x19, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x32, 0x74, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x69,
	0x70, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc0, 0x01,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x42, 0x6f, 0x6e, 0x65, 0x2f, 0x69, 0x70, 0x73, 0x2d, 0x72, 0x73, 0x73, 0x69,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x49, 0x4d, 0x58, 0xaa, 0x02, 0x0c, 0x49, 0x70, 0x73, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x49, 0x70, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x49, 0x70, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0e, 0x49, 0x70, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ips_model_v1_model_proto_rawDescOnce sync.Once
	file_ips_model_v1_model_proto_rawDescData = file_ips_model_v1_model_proto_rawDesc
)

func file_ips_model_v1_model_proto_rawDescGZIP() []byte {
	file_ips_model_v1_model_proto_rawDescOnce.Do(func() {
		file_ips_model_v1_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_ips_model_v1_model_proto_rawDescData)
	})
	return file_ips_model_v1_model_proto_rawDescData
}

var file_ips_model_v1_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ips_model_v1_model_proto_goTypes = []interface{}{
	(*PredictCoordinateRequest)(nil),  // 0: ips.model.v1.PredictCoordinateRequest
	(*PredictCoordinateResponse)(nil), // 1: ips.model.v1.PredictCoordinateResponse
}
var file_ips_model_v1_model_proto_depIdxs = []int32{
	0, // 0: ips.model.v1.ModelService.PredictCoordinate:input_type -> ips.model.v1.PredictCoordinateRequest
	1, // 1: ips.model.v1.ModelService.PredictCoordinate:output_type -> ips.model.v1.PredictCoordinateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ips_model_v1_model_proto_init() }
func file_ips_model_v1_model_proto_init() {
	if File_ips_model_v1_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ips_model_v1_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictCoordinateRequest); i {
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
		file_ips_model_v1_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictCoordinateResponse); i {
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
			RawDescriptor: file_ips_model_v1_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ips_model_v1_model_proto_goTypes,
		DependencyIndexes: file_ips_model_v1_model_proto_depIdxs,
		MessageInfos:      file_ips_model_v1_model_proto_msgTypes,
	}.Build()
	File_ips_model_v1_model_proto = out.File
	file_ips_model_v1_model_proto_rawDesc = nil
	file_ips_model_v1_model_proto_goTypes = nil
	file_ips_model_v1_model_proto_depIdxs = nil
}