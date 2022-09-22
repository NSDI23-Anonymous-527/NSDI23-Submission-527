// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: ca.proto

package ca

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

type RhineSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Sig         []byte `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"`
	DataPostfix string `protobuf:"bytes,3,opt,name=DataPostfix,proto3" json:"DataPostfix,omitempty"`
}

func (x *RhineSig) Reset() {
	*x = RhineSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RhineSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RhineSig) ProtoMessage() {}

func (x *RhineSig) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RhineSig.ProtoReflect.Descriptor instead.
func (*RhineSig) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{0}
}

func (x *RhineSig) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RhineSig) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *RhineSig) GetDataPostfix() string {
	if x != nil {
		return x.DataPostfix
	}
	return ""
}

type SubmitNewDelegCARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acsr   *RhineSig `protobuf:"bytes,1,opt,name=Acsr,proto3" json:"Acsr,omitempty"`
	Rcertp []byte    `protobuf:"bytes,2,opt,name=Rcertp,proto3" json:"Rcertp,omitempty"`
	Rid    []byte    `protobuf:"bytes,3,opt,name=Rid,proto3" json:"Rid,omitempty"`
}

func (x *SubmitNewDelegCARequest) Reset() {
	*x = SubmitNewDelegCARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitNewDelegCARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitNewDelegCARequest) ProtoMessage() {}

func (x *SubmitNewDelegCARequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitNewDelegCARequest.ProtoReflect.Descriptor instead.
func (*SubmitNewDelegCARequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitNewDelegCARequest) GetAcsr() *RhineSig {
	if x != nil {
		return x.Acsr
	}
	return nil
}

func (x *SubmitNewDelegCARequest) GetRcertp() []byte {
	if x != nil {
		return x.Rcertp
	}
	return nil
}

func (x *SubmitNewDelegCARequest) GetRid() []byte {
	if x != nil {
		return x.Rid
	}
	return nil
}

type SubmitNewDelegCAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rcertc []byte   `protobuf:"bytes,1,opt,name=Rcertc,proto3" json:"Rcertc,omitempty"`
	Lcfms  [][]byte `protobuf:"bytes,2,rep,name=Lcfms,proto3" json:"Lcfms,omitempty"`
	Rid    []byte   `protobuf:"bytes,3,opt,name=Rid,proto3" json:"Rid,omitempty"`
}

func (x *SubmitNewDelegCAResponse) Reset() {
	*x = SubmitNewDelegCAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitNewDelegCAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitNewDelegCAResponse) ProtoMessage() {}

func (x *SubmitNewDelegCAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitNewDelegCAResponse.ProtoReflect.Descriptor instead.
func (*SubmitNewDelegCAResponse) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitNewDelegCAResponse) GetRcertc() []byte {
	if x != nil {
		return x.Rcertc
	}
	return nil
}

func (x *SubmitNewDelegCAResponse) GetLcfms() [][]byte {
	if x != nil {
		return x.Lcfms
	}
	return nil
}

func (x *SubmitNewDelegCAResponse) GetRid() []byte {
	if x != nil {
		return x.Rid
	}
	return nil
}

var File_ca_proto protoreflect.FileDescriptor

var file_ca_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x63, 0x61, 0x22, 0x52,
	0x0a, 0x08, 0x52, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x53, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x69, 0x67,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x66, 0x69, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x73, 0x74, 0x66,
	0x69, 0x78, 0x22, 0x65, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x43, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x04, 0x41, 0x63, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x61,
	0x2e, 0x52, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x69, 0x67, 0x52, 0x04, 0x41, 0x63, 0x73, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x63, 0x65, 0x72, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x52, 0x63, 0x65, 0x72, 0x74, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x52, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x18, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x43, 0x41, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x63, 0x65, 0x72, 0x74, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x52, 0x63, 0x65, 0x72, 0x74, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x63, 0x66, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x4c, 0x63,
	0x66, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x52, 0x69, 0x64, 0x32, 0x5c, 0x0a, 0x09, 0x43, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x65, 0x77, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x43, 0x41, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x43, 0x41, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e,
	0x65, 0x77, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x43, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x52, 0x48, 0x49,
	0x4e, 0x45, 0x2d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ca_proto_rawDescOnce sync.Once
	file_ca_proto_rawDescData = file_ca_proto_rawDesc
)

func file_ca_proto_rawDescGZIP() []byte {
	file_ca_proto_rawDescOnce.Do(func() {
		file_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_ca_proto_rawDescData)
	})
	return file_ca_proto_rawDescData
}

var file_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ca_proto_goTypes = []interface{}{
	(*RhineSig)(nil),                 // 0: ca.RhineSig
	(*SubmitNewDelegCARequest)(nil),  // 1: ca.SubmitNewDelegCARequest
	(*SubmitNewDelegCAResponse)(nil), // 2: ca.SubmitNewDelegCAResponse
}
var file_ca_proto_depIdxs = []int32{
	0, // 0: ca.SubmitNewDelegCARequest.Acsr:type_name -> ca.RhineSig
	1, // 1: ca.CAService.SubmitNewDelegCA:input_type -> ca.SubmitNewDelegCARequest
	2, // 2: ca.CAService.SubmitNewDelegCA:output_type -> ca.SubmitNewDelegCAResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ca_proto_init() }
func file_ca_proto_init() {
	if File_ca_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RhineSig); i {
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
		file_ca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitNewDelegCARequest); i {
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
		file_ca_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitNewDelegCAResponse); i {
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
			RawDescriptor: file_ca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ca_proto_goTypes,
		DependencyIndexes: file_ca_proto_depIdxs,
		MessageInfos:      file_ca_proto_msgTypes,
	}.Build()
	File_ca_proto = out.File
	file_ca_proto_rawDesc = nil
	file_ca_proto_goTypes = nil
	file_ca_proto_depIdxs = nil
}