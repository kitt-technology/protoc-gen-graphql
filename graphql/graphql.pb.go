// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: graphql/graphql.proto

package graphql

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *BatchRequest) Reset() {
	*x = BatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_graphql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRequest) ProtoMessage() {}

func (x *BatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_graphql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRequest.ProtoReflect.Descriptor instead.
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return file_graphql_graphql_proto_rawDescGZIP(), []int{0}
}

func (x *BatchRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount  int32  `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	EndCursor   string `protobuf:"bytes,2,opt,name=end_cursor,json=endCursor,proto3" json:"end_cursor,omitempty"`
	HasNextPage bool   `protobuf:"varint,3,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_graphql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_graphql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_graphql_graphql_proto_rawDescGZIP(), []int{1}
}

func (x *PageInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *PageInfo) GetEndCursor() string {
	if x != nil {
		return x.EndCursor
	}
	return ""
}

func (x *PageInfo) GetHasNextPage() bool {
	if x != nil {
		return x.HasNextPage
	}
	return false
}

type FieldMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths    []string        `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	PathsMap map[string]bool `protobuf:"bytes,2,rep,name=paths_map,json=pathsMap,proto3" json:"paths_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *FieldMask) Reset() {
	*x = FieldMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphql_graphql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMask) ProtoMessage() {}

func (x *FieldMask) ProtoReflect() protoreflect.Message {
	mi := &file_graphql_graphql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMask.ProtoReflect.Descriptor instead.
func (*FieldMask) Descriptor() ([]byte, []int) {
	return file_graphql_graphql_proto_rawDescGZIP(), []int{2}
}

func (x *FieldMask) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *FieldMask) GetPathsMap() map[string]bool {
	if x != nil {
		return x.PathsMap
	}
	return nil
}

var file_graphql_graphql_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1085,
		Name:          "graphql.object_name",
		Tag:           "bytes,1085,opt,name=object_name",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1086,
		Name:          "graphql.mutation",
		Tag:           "varint,1086,opt,name=mutation",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1087,
		Name:          "graphql.skip_message",
		Tag:           "varint,1087,opt,name=skip_message",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1055,
		Name:          "graphql.disabled",
		Tag:           "varint,1055,opt,name=disabled",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1056,
		Name:          "graphql.package",
		Tag:           "bytes,1056,opt,name=package",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1088,
		Name:          "graphql.host",
		Tag:           "bytes,1088,opt,name=host",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1086,
		Name:          "graphql.enum_name",
		Tag:           "bytes,1086,opt,name=enum_name",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1096,
		Name:          "graphql.optional",
		Tag:           "varint,1096,opt,name=optional",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1097,
		Name:          "graphql.skip_field",
		Tag:           "varint,1097,opt,name=skip_field",
		Filename:      "graphql/graphql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1098,
		Name:          "graphql.batch_loader",
		Tag:           "varint,1098,opt,name=batch_loader",
		Filename:      "graphql/graphql.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string object_name = 1085;
	E_ObjectName = &file_graphql_graphql_proto_extTypes[0]
	// optional bool mutation = 1086;
	E_Mutation = &file_graphql_graphql_proto_extTypes[1]
	// optional bool skip_message = 1087;
	E_SkipMessage = &file_graphql_graphql_proto_extTypes[2]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional bool disabled = 1055;
	E_Disabled = &file_graphql_graphql_proto_extTypes[3]
	// optional string package = 1056;
	E_Package = &file_graphql_graphql_proto_extTypes[4]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional string host = 1088;
	E_Host = &file_graphql_graphql_proto_extTypes[5]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional string enum_name = 1086;
	E_EnumName = &file_graphql_graphql_proto_extTypes[6]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool optional = 1096;
	E_Optional = &file_graphql_graphql_proto_extTypes[7]
	// optional bool skip_field = 1097;
	E_SkipField = &file_graphql_graphql_proto_extTypes[8]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional bool batch_loader = 1098;
	E_BatchLoader = &file_graphql_graphql_proto_extTypes[9]
)

var File_graphql_graphql_proto protoreflect.FileDescriptor

var file_graphql_graphql_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x6e, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x3d, 0x0a, 0x09, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x70, 0x61, 0x74, 0x68, 0x73, 0x4d, 0x61, 0x70, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x61, 0x74,
	0x68, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x41, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x3c, 0x0a, 0x08, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbe, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x43, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbf, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x73, 0x6b, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x39, 0x0a, 0x08,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9f, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x3a, 0x37, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xa0, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x3a, 0x34, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc0, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x3a, 0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xbe, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x75, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x3a, 0x3a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc8, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x3a, 0x3d,
	0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc9, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x42, 0x0a,
	0x0c, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xca, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x69, 0x74, 0x74, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_graphql_graphql_proto_rawDescOnce sync.Once
	file_graphql_graphql_proto_rawDescData = file_graphql_graphql_proto_rawDesc
)

func file_graphql_graphql_proto_rawDescGZIP() []byte {
	file_graphql_graphql_proto_rawDescOnce.Do(func() {
		file_graphql_graphql_proto_rawDescData = protoimpl.X.CompressGZIP(file_graphql_graphql_proto_rawDescData)
	})
	return file_graphql_graphql_proto_rawDescData
}

var file_graphql_graphql_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_graphql_graphql_proto_goTypes = []interface{}{
	(*BatchRequest)(nil),                // 0: graphql.BatchRequest
	(*PageInfo)(nil),                    // 1: graphql.PageInfo
	(*FieldMask)(nil),                   // 2: graphql.FieldMask
	nil,                                 // 3: graphql.FieldMask.PathsMapEntry
	(*descriptorpb.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
	(*descriptorpb.FileOptions)(nil),    // 5: google.protobuf.FileOptions
	(*descriptorpb.ServiceOptions)(nil), // 6: google.protobuf.ServiceOptions
	(*descriptorpb.EnumOptions)(nil),    // 7: google.protobuf.EnumOptions
	(*descriptorpb.FieldOptions)(nil),   // 8: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil),  // 9: google.protobuf.MethodOptions
}
var file_graphql_graphql_proto_depIdxs = []int32{
	3,  // 0: graphql.FieldMask.paths_map:type_name -> graphql.FieldMask.PathsMapEntry
	4,  // 1: graphql.object_name:extendee -> google.protobuf.MessageOptions
	4,  // 2: graphql.mutation:extendee -> google.protobuf.MessageOptions
	4,  // 3: graphql.skip_message:extendee -> google.protobuf.MessageOptions
	5,  // 4: graphql.disabled:extendee -> google.protobuf.FileOptions
	5,  // 5: graphql.package:extendee -> google.protobuf.FileOptions
	6,  // 6: graphql.host:extendee -> google.protobuf.ServiceOptions
	7,  // 7: graphql.enum_name:extendee -> google.protobuf.EnumOptions
	8,  // 8: graphql.optional:extendee -> google.protobuf.FieldOptions
	8,  // 9: graphql.skip_field:extendee -> google.protobuf.FieldOptions
	9,  // 10: graphql.batch_loader:extendee -> google.protobuf.MethodOptions
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	1,  // [1:11] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_graphql_graphql_proto_init() }
func file_graphql_graphql_proto_init() {
	if File_graphql_graphql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graphql_graphql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchRequest); i {
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
		file_graphql_graphql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_graphql_graphql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMask); i {
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
			RawDescriptor: file_graphql_graphql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 10,
			NumServices:   0,
		},
		GoTypes:           file_graphql_graphql_proto_goTypes,
		DependencyIndexes: file_graphql_graphql_proto_depIdxs,
		MessageInfos:      file_graphql_graphql_proto_msgTypes,
		ExtensionInfos:    file_graphql_graphql_proto_extTypes,
	}.Build()
	File_graphql_graphql_proto = out.File
	file_graphql_graphql_proto_rawDesc = nil
	file_graphql_graphql_proto_goTypes = nil
	file_graphql_graphql_proto_depIdxs = nil
}
