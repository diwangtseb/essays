// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pb.proto

package pb

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

type TransInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransInReq) Reset() {
	*x = TransInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransInReq) ProtoMessage() {}

func (x *TransInReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransInReq.ProtoReflect.Descriptor instead.
func (*TransInReq) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0}
}

func (x *TransInReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TransInReq) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransOutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransOutReq) Reset() {
	*x = TransOutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransOutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransOutReq) ProtoMessage() {}

func (x *TransOutReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransOutReq.ProtoReflect.Descriptor instead.
func (*TransOutReq) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{1}
}

func (x *TransOutReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *TransOutReq) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransInReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TransInReply) Reset() {
	*x = TransInReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransInReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransInReply) ProtoMessage() {}

func (x *TransInReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransInReply.ProtoReflect.Descriptor instead.
func (*TransInReply) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{2}
}

func (x *TransInReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type TransOutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TransOutReply) Reset() {
	*x = TransOutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransOutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransOutReply) ProtoMessage() {}

func (x *TransOutReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransOutReply.ProtoReflect.Descriptor instead.
func (*TransOutReply) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{3}
}

func (x *TransOutReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_pb_proto protoreflect.FileDescriptor

var file_pb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x74, 0x2e, 0x70,
	0x62, 0x22, 0x36, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x29, 0x0a, 0x0d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xb4, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x49, 0x6e, 0x12, 0x11, 0x2e, 0x64, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x64, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x52, 0x6f, 0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x64,
	0x74, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x64, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f,
	0x75, 0x74, 0x12, 0x12, 0x2e, 0x64, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x64, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_rawDescOnce sync.Once
	file_pb_proto_rawDescData = file_pb_proto_rawDesc
)

func file_pb_proto_rawDescGZIP() []byte {
	file_pb_proto_rawDescOnce.Do(func() {
		file_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_rawDescData)
	})
	return file_pb_proto_rawDescData
}

var file_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_proto_goTypes = []interface{}{
	(*TransInReq)(nil),    // 0: dt.pb.TransInReq
	(*TransOutReq)(nil),   // 1: dt.pb.TransOutReq
	(*TransInReply)(nil),  // 2: dt.pb.TransInReply
	(*TransOutReply)(nil), // 3: dt.pb.TransOutReply
}
var file_pb_proto_depIdxs = []int32{
	0, // 0: dt.pb.TransService.TransIn:input_type -> dt.pb.TransInReq
	0, // 1: dt.pb.TransService.TransInRoll:input_type -> dt.pb.TransInReq
	1, // 2: dt.pb.TransService.TransOut:input_type -> dt.pb.TransOutReq
	2, // 3: dt.pb.TransService.TransIn:output_type -> dt.pb.TransInReply
	2, // 4: dt.pb.TransService.TransInRoll:output_type -> dt.pb.TransInReply
	3, // 5: dt.pb.TransService.TransOut:output_type -> dt.pb.TransOutReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_proto_init() }
func file_pb_proto_init() {
	if File_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransInReq); i {
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
		file_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransOutReq); i {
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
		file_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransInReply); i {
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
		file_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransOutReply); i {
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
			RawDescriptor: file_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_proto_goTypes,
		DependencyIndexes: file_pb_proto_depIdxs,
		MessageInfos:      file_pb_proto_msgTypes,
	}.Build()
	File_pb_proto = out.File
	file_pb_proto_rawDesc = nil
	file_pb_proto_goTypes = nil
	file_pb_proto_depIdxs = nil
}
