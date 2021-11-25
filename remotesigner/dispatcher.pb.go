// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: dispatcher.proto

package remotesigner

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

type SignMilestoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeys   [][]byte `protobuf:"bytes,1,rep,name=pubKeys,proto3" json:"pubKeys,omitempty"`
	MsEssence []byte   `protobuf:"bytes,2,opt,name=msEssence,proto3" json:"msEssence,omitempty"`
}

func (x *SignMilestoneRequest) Reset() {
	*x = SignMilestoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignMilestoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignMilestoneRequest) ProtoMessage() {}

func (x *SignMilestoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignMilestoneRequest.ProtoReflect.Descriptor instead.
func (*SignMilestoneRequest) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{0}
}

func (x *SignMilestoneRequest) GetPubKeys() [][]byte {
	if x != nil {
		return x.PubKeys
	}
	return nil
}

func (x *SignMilestoneRequest) GetMsEssence() []byte {
	if x != nil {
		return x.MsEssence
	}
	return nil
}

type SignMilestoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signatures [][]byte `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *SignMilestoneResponse) Reset() {
	*x = SignMilestoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignMilestoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignMilestoneResponse) ProtoMessage() {}

func (x *SignMilestoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dispatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignMilestoneResponse.ProtoReflect.Descriptor instead.
func (*SignMilestoneResponse) Descriptor() ([]byte, []int) {
	return file_dispatcher_proto_rawDescGZIP(), []int{1}
}

func (x *SignMilestoneResponse) GetSignatures() [][]byte {
	if x != nil {
		return x.Signatures
	}
	return nil
}

var File_dispatcher_proto protoreflect.FileDescriptor

var file_dispatcher_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x76, 0x33,
	0x22, 0x4e, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x73, 0x45, 0x73, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x73, 0x45, 0x73, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x37, 0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x32, 0x6f, 0x0a, 0x13, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x58, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x12, 0x22, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x76, 0x33,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x76, 0x33, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x61, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x69, 0x6f, 0x74, 0x61, 0x2e, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x3b, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_dispatcher_proto_rawDescOnce sync.Once
	file_dispatcher_proto_rawDescData = file_dispatcher_proto_rawDesc
)

func file_dispatcher_proto_rawDescGZIP() []byte {
	file_dispatcher_proto_rawDescOnce.Do(func() {
		file_dispatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_dispatcher_proto_rawDescData)
	})
	return file_dispatcher_proto_rawDescData
}

var file_dispatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dispatcher_proto_goTypes = []interface{}{
	(*SignMilestoneRequest)(nil),  // 0: dispatcherv3.SignMilestoneRequest
	(*SignMilestoneResponse)(nil), // 1: dispatcherv3.SignMilestoneResponse
}
var file_dispatcher_proto_depIdxs = []int32{
	0, // 0: dispatcherv3.SignatureDispatcher.SignMilestone:input_type -> dispatcherv3.SignMilestoneRequest
	1, // 1: dispatcherv3.SignatureDispatcher.SignMilestone:output_type -> dispatcherv3.SignMilestoneResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dispatcher_proto_init() }
func file_dispatcher_proto_init() {
	if File_dispatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dispatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignMilestoneRequest); i {
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
		file_dispatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignMilestoneResponse); i {
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
			RawDescriptor: file_dispatcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dispatcher_proto_goTypes,
		DependencyIndexes: file_dispatcher_proto_depIdxs,
		MessageInfos:      file_dispatcher_proto_msgTypes,
	}.Build()
	File_dispatcher_proto = out.File
	file_dispatcher_proto_rawDesc = nil
	file_dispatcher_proto_goTypes = nil
	file_dispatcher_proto_depIdxs = nil
}
