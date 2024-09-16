// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: deep.proto

package proto

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

type NumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *NumReq) Reset() {
	*x = NumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumReq) ProtoMessage() {}

func (x *NumReq) ProtoReflect() protoreflect.Message {
	mi := &file_deep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumReq.ProtoReflect.Descriptor instead.
func (*NumReq) Descriptor() ([]byte, []int) {
	return file_deep_proto_rawDescGZIP(), []int{0}
}

func (x *NumReq) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *NumReq) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type Num struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Num) Reset() {
	*x = Num{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deep_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Num) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Num) ProtoMessage() {}

func (x *Num) ProtoReflect() protoreflect.Message {
	mi := &file_deep_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Num.ProtoReflect.Descriptor instead.
func (*Num) Descriptor() ([]byte, []int) {
	return file_deep_proto_rawDescGZIP(), []int{1}
}

func (x *Num) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

type NumResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C int32 `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *NumResp) Reset() {
	*x = NumResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deep_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumResp) ProtoMessage() {}

func (x *NumResp) ProtoReflect() protoreflect.Message {
	mi := &file_deep_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumResp.ProtoReflect.Descriptor instead.
func (*NumResp) Descriptor() ([]byte, []int) {
	return file_deep_proto_rawDescGZIP(), []int{2}
}

func (x *NumResp) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

var File_deep_proto protoreflect.FileDescriptor

var file_deep_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65,
	0x65, 0x70, 0x22, 0x24, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x13, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x12,
	0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x22, 0x17, 0x0a,
	0x07, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x63, 0x32, 0x61, 0x0a, 0x07, 0x44, 0x65, 0x65, 0x70, 0x53, 0x72,
	0x76, 0x12, 0x28, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x12, 0x0c,
	0x2e, 0x64, 0x65, 0x65, 0x70, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x64,
	0x65, 0x65, 0x70, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x41, 0x6c, 0x6c, 0x54, 0x68, 0x65, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x09, 0x2e,
	0x64, 0x65, 0x65, 0x70, 0x2e, 0x4e, 0x75, 0x6d, 0x1a, 0x0d, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x2e,
	0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x28, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_deep_proto_rawDescOnce sync.Once
	file_deep_proto_rawDescData = file_deep_proto_rawDesc
)

func file_deep_proto_rawDescGZIP() []byte {
	file_deep_proto_rawDescOnce.Do(func() {
		file_deep_proto_rawDescData = protoimpl.X.CompressGZIP(file_deep_proto_rawDescData)
	})
	return file_deep_proto_rawDescData
}

var file_deep_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_deep_proto_goTypes = []any{
	(*NumReq)(nil),  // 0: deep.NumReq
	(*Num)(nil),     // 1: deep.Num
	(*NumResp)(nil), // 2: deep.NumResp
}
var file_deep_proto_depIdxs = []int32{
	0, // 0: deep.DeepSrv.AddTwoNum:input_type -> deep.NumReq
	1, // 1: deep.DeepSrv.AddAllTheseNum:input_type -> deep.Num
	2, // 2: deep.DeepSrv.AddTwoNum:output_type -> deep.NumResp
	2, // 3: deep.DeepSrv.AddAllTheseNum:output_type -> deep.NumResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_deep_proto_init() }
func file_deep_proto_init() {
	if File_deep_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deep_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NumReq); i {
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
		file_deep_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Num); i {
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
		file_deep_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NumResp); i {
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
			RawDescriptor: file_deep_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deep_proto_goTypes,
		DependencyIndexes: file_deep_proto_depIdxs,
		MessageInfos:      file_deep_proto_msgTypes,
	}.Build()
	File_deep_proto = out.File
	file_deep_proto_rawDesc = nil
	file_deep_proto_goTypes = nil
	file_deep_proto_depIdxs = nil
}
