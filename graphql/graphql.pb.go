// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: graphql.proto

package graphql

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MutationOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Success string `protobuf:"bytes,2,opt,name=success,proto3" json:"success,omitempty"`
	Failure string `protobuf:"bytes,3,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *MutationOption) Reset() {
	*x = MutationOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutationOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutationOption) ProtoMessage() {}

func (x *MutationOption) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutationOption.ProtoReflect.Descriptor instead.
func (*MutationOption) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{0}
}

func (x *MutationOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MutationOption) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

func (x *MutationOption) GetFailure() string {
	if x != nil {
		return x.Failure
	}
	return ""
}

type FieldResolver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName      string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	DataloaderName string `protobuf:"bytes,2,opt,name=dataloader_name,json=dataloaderName,proto3" json:"dataloader_name,omitempty"`
}

func (x *FieldResolver) Reset() {
	*x = FieldResolver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldResolver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldResolver) ProtoMessage() {}

func (x *FieldResolver) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldResolver.ProtoReflect.Descriptor instead.
func (*FieldResolver) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{1}
}

func (x *FieldResolver) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *FieldResolver) GetDataloaderName() string {
	if x != nil {
		return x.DataloaderName
	}
	return ""
}

type BatchOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys    string `protobuf:"bytes,1,opt,name=keys,proto3" json:"keys,omitempty"`
	Results string `protobuf:"bytes,2,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *BatchOptions) Reset() {
	*x = BatchOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOptions) ProtoMessage() {}

func (x *BatchOptions) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOptions.ProtoReflect.Descriptor instead.
func (*BatchOptions) Descriptor() ([]byte, []int) {
	return file_graphql_proto_rawDescGZIP(), []int{2}
}

func (x *BatchOptions) GetKeys() string {
	if x != nil {
		return x.Keys
	}
	return ""
}

func (x *BatchOptions) GetResults() string {
	if x != nil {
		return x.Results
	}
	return ""
}

var file_graphql_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*MutationOption)(nil),
		Field:         1084,
		Name:          "graphql.mutation_options",
		Tag:           "bytes,1084,opt,name=mutation_options",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1086,
		Name:          "graphql.dataloader_ids",
		Tag:           "varint,1086,opt,name=dataloader_ids",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1087,
		Name:          "graphql.dataloader_object",
		Tag:           "varint,1087,opt,name=dataloader_object",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldResolver)(nil),
		Field:         1085,
		Name:          "graphql.field_resolver",
		Tag:           "bytes,1085,opt,name=field_resolver",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*BatchOptions)(nil),
		Field:         1087,
		Name:          "graphql.batch",
		Tag:           "bytes,1087,opt,name=batch",
		Filename:      "graphql.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1088,
		Name:          "graphql.host",
		Tag:           "bytes,1088,opt,name=host",
		Filename:      "graphql.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional graphql.MutationOption mutation_options = 1084;
	E_MutationOptions = &file_graphql_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional bool dataloader_ids = 1086;
	E_DataloaderIds = &file_graphql_proto_extTypes[1]
	// optional bool dataloader_object = 1087;
	E_DataloaderObject = &file_graphql_proto_extTypes[2]
	// optional graphql.FieldResolver field_resolver = 1085;
	E_FieldResolver = &file_graphql_proto_extTypes[3]
)

// Extension fields to descriptor.MethodOptions.
var (
	// optional graphql.BatchOptions batch = 1087;
	E_Batch = &file_graphql_proto_extTypes[4]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// optional string host = 1088;
	E_Host = &file_graphql_proto_extTypes[5]
)

var File_graphql_proto protoreflect.FileDescriptor

var file_graphql_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a,
	0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x3a, 0x64, 0x0a, 0x10, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xbc, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x3a, 0x45, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xbe, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x3a, 0x4b, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbf, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x5d, 0x0a, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x3a, 0x4c, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbf, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x3a, 0x34, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc0, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graphql_proto_rawDescOnce sync.Once
	file_graphql_proto_rawDescData = file_graphql_proto_rawDesc
)

func file_graphql_proto_rawDescGZIP() []byte {
	file_graphql_proto_rawDescOnce.Do(func() {
		file_graphql_proto_rawDescData = protoimpl.X.CompressGZIP(file_graphql_proto_rawDescData)
	})
	return file_graphql_proto_rawDescData
}

var file_graphql_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_graphql_proto_goTypes = []interface{}{
	(*MutationOption)(nil),            // 0: graphql.MutationOption
	(*FieldResolver)(nil),             // 1: graphql.FieldResolver
	(*BatchOptions)(nil),              // 2: graphql.BatchOptions
	(*descriptor.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
	(*descriptor.MethodOptions)(nil),  // 5: google.protobuf.MethodOptions
	(*descriptor.ServiceOptions)(nil), // 6: google.protobuf.ServiceOptions
}
var file_graphql_proto_depIdxs = []int32{
	3, // 0: graphql.mutation_options:extendee -> google.protobuf.MessageOptions
	4, // 1: graphql.dataloader_ids:extendee -> google.protobuf.FieldOptions
	4, // 2: graphql.dataloader_object:extendee -> google.protobuf.FieldOptions
	4, // 3: graphql.field_resolver:extendee -> google.protobuf.FieldOptions
	5, // 4: graphql.batch:extendee -> google.protobuf.MethodOptions
	6, // 5: graphql.host:extendee -> google.protobuf.ServiceOptions
	0, // 6: graphql.mutation_options:type_name -> graphql.MutationOption
	1, // 7: graphql.field_resolver:type_name -> graphql.FieldResolver
	2, // 8: graphql.batch:type_name -> graphql.BatchOptions
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	6, // [6:9] is the sub-list for extension type_name
	0, // [0:6] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_graphql_proto_init() }
func file_graphql_proto_init() {
	if File_graphql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graphql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutationOption); i {
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
		file_graphql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldResolver); i {
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
		file_graphql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOptions); i {
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
			RawDescriptor: file_graphql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_graphql_proto_goTypes,
		DependencyIndexes: file_graphql_proto_depIdxs,
		MessageInfos:      file_graphql_proto_msgTypes,
		ExtensionInfos:    file_graphql_proto_extTypes,
	}.Build()
	File_graphql_proto = out.File
	file_graphql_proto_rawDesc = nil
	file_graphql_proto_goTypes = nil
	file_graphql_proto_depIdxs = nil
}
