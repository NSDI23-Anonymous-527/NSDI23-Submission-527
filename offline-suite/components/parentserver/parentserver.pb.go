// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: parentserver.proto

package parentserver

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

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Sig  []byte `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"`
}

func (x *RhineSig) Reset() {
	*x = RhineSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parentserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RhineSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RhineSig) ProtoMessage() {}

func (x *RhineSig) ProtoReflect() protoreflect.Message {
	mi := &file_parentserver_proto_msgTypes[0]
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
	return file_parentserver_proto_rawDescGZIP(), []int{0}
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

// Message sent from child to parent to initialize delegation
type InitDelegationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid []byte `protobuf:"bytes,1,opt,name=Rid,proto3" json:"Rid,omitempty"`
	Csr []byte `protobuf:"bytes,2,opt,name=Csr,proto3" json:"Csr,omitempty"`
}

func (x *InitDelegationRequest) Reset() {
	*x = InitDelegationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parentserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitDelegationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDelegationRequest) ProtoMessage() {}

func (x *InitDelegationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parentserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDelegationRequest.ProtoReflect.Descriptor instead.
func (*InitDelegationRequest) Descriptor() ([]byte, []int) {
	return file_parentserver_proto_rawDescGZIP(), []int{1}
}

func (x *InitDelegationRequest) GetRid() []byte {
	if x != nil {
		return x.Rid
	}
	return nil
}

func (x *InitDelegationRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

// Message received by child as response to delegation
type InitDelegationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Approvalcommit *RhineSig `protobuf:"bytes,1,opt,name=Approvalcommit,proto3" json:"Approvalcommit,omitempty"`
	Rcertp         []byte    `protobuf:"bytes,2,opt,name=Rcertp,proto3" json:"Rcertp,omitempty"`
}

func (x *InitDelegationResponse) Reset() {
	*x = InitDelegationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parentserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitDelegationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDelegationResponse) ProtoMessage() {}

func (x *InitDelegationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parentserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDelegationResponse.ProtoReflect.Descriptor instead.
func (*InitDelegationResponse) Descriptor() ([]byte, []int) {
	return file_parentserver_proto_rawDescGZIP(), []int{2}
}

func (x *InitDelegationResponse) GetApprovalcommit() *RhineSig {
	if x != nil {
		return x.Approvalcommit
	}
	return nil
}

func (x *InitDelegationResponse) GetRcertp() []byte {
	if x != nil {
		return x.Rcertp
	}
	return nil
}

var File_parentserver_proto protoreflect.FileDescriptor

var file_parentserver_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x53, 0x69, 0x67, 0x22, 0x3b, 0x0a, 0x15, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x52, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x52, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x43, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x43, 0x73,
	0x72, 0x22, 0x70, 0x0a, 0x16, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x69, 0x67, 0x52, 0x0e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x63, 0x65, 0x72, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x52, 0x63, 0x65,
	0x72, 0x74, 0x70, 0x32, 0x6e, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x68, 0x69, 0x6e, 0x65, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x52, 0x48, 0x49,
	0x4e, 0x45, 0x2d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x41, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parentserver_proto_rawDescOnce sync.Once
	file_parentserver_proto_rawDescData = file_parentserver_proto_rawDesc
)

func file_parentserver_proto_rawDescGZIP() []byte {
	file_parentserver_proto_rawDescOnce.Do(func() {
		file_parentserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_parentserver_proto_rawDescData)
	})
	return file_parentserver_proto_rawDescData
}

var file_parentserver_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_parentserver_proto_goTypes = []interface{}{
	(*RhineSig)(nil),               // 0: parentserver.RhineSig
	(*InitDelegationRequest)(nil),  // 1: parentserver.InitDelegationRequest
	(*InitDelegationResponse)(nil), // 2: parentserver.InitDelegationResponse
}
var file_parentserver_proto_depIdxs = []int32{
	0, // 0: parentserver.InitDelegationResponse.Approvalcommit:type_name -> parentserver.RhineSig
	1, // 1: parentserver.ParentService.InitDelegation:input_type -> parentserver.InitDelegationRequest
	2, // 2: parentserver.ParentService.InitDelegation:output_type -> parentserver.InitDelegationResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_parentserver_proto_init() }
func file_parentserver_proto_init() {
	if File_parentserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parentserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_parentserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitDelegationRequest); i {
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
		file_parentserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitDelegationResponse); i {
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
			RawDescriptor: file_parentserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parentserver_proto_goTypes,
		DependencyIndexes: file_parentserver_proto_depIdxs,
		MessageInfos:      file_parentserver_proto_msgTypes,
	}.Build()
	File_parentserver_proto = out.File
	file_parentserver_proto_rawDesc = nil
	file_parentserver_proto_goTypes = nil
	file_parentserver_proto_depIdxs = nil
}
