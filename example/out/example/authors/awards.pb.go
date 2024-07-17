// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.27.0
// source: awards.proto

package authors

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

type AwardType int32

const (
	AwardType_UnknownAwardType  AwardType = 0
	AwardType_BookOfTheYear     AwardType = 1
	AwardType_BestModernFiction AwardType = 2
)

// Enum value maps for AwardType.
var (
	AwardType_name = map[int32]string{
		0: "UnknownAwardType",
		1: "BookOfTheYear",
		2: "BestModernFiction",
	}
	AwardType_value = map[string]int32{
		"UnknownAwardType":  0,
		"BookOfTheYear":     1,
		"BestModernFiction": 2,
	}
)

func (x AwardType) Enum() *AwardType {
	p := new(AwardType)
	*p = x
	return p
}

func (x AwardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AwardType) Descriptor() protoreflect.EnumDescriptor {
	return file_awards_proto_enumTypes[0].Descriptor()
}

func (AwardType) Type() protoreflect.EnumType {
	return &file_awards_proto_enumTypes[0]
}

func (x AwardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AwardType.Descriptor instead.
func (AwardType) EnumDescriptor() ([]byte, []int) {
	return file_awards_proto_rawDescGZIP(), []int{0}
}

type Award struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year       int64     `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Importance int64     `protobuf:"varint,3,opt,name=importance,proto3" json:"importance,omitempty"`
	Type       AwardType `protobuf:"varint,4,opt,name=type,proto3,enum=authors.AwardType" json:"type,omitempty"`
	Id         string    `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Award) Reset() {
	*x = Award{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Award) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Award) ProtoMessage() {}

func (x *Award) ProtoReflect() protoreflect.Message {
	mi := &file_awards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Award.ProtoReflect.Descriptor instead.
func (*Award) Descriptor() ([]byte, []int) {
	return file_awards_proto_rawDescGZIP(), []int{0}
}

func (x *Award) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Award) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Award) GetImportance() int64 {
	if x != nil {
		return x.Importance
	}
	return 0
}

func (x *Award) GetType() AwardType {
	if x != nil {
		return x.Type
	}
	return AwardType_UnknownAwardType
}

func (x *Award) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAwardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetAwardsRequest) Reset() {
	*x = GetAwardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwardsRequest) ProtoMessage() {}

func (x *GetAwardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_awards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwardsRequest.ProtoReflect.Descriptor instead.
func (*GetAwardsRequest) Descriptor() ([]byte, []int) {
	return file_awards_proto_rawDescGZIP(), []int{1}
}

func (x *GetAwardsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetAwardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Awards []*Award `protobuf:"bytes,1,rep,name=awards,proto3" json:"awards,omitempty"`
}

func (x *GetAwardsResponse) Reset() {
	*x = GetAwardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_awards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAwardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAwardsResponse) ProtoMessage() {}

func (x *GetAwardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_awards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAwardsResponse.ProtoReflect.Descriptor instead.
func (*GetAwardsResponse) Descriptor() ([]byte, []int) {
	return file_awards_proto_rawDescGZIP(), []int{2}
}

func (x *GetAwardsResponse) GetAwards() []*Award {
	if x != nil {
		return x.Awards
	}
	return nil
}

var File_awards_proto protoreflect.FileDescriptor

var file_awards_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x41, 0x77, 0x61, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x2e, 0x41, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x77, 0x61, 0x72, 0x64, 0x52, 0x06,
	0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2a, 0x4b, 0x0a, 0x09, 0x41, 0x77, 0x61, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x41, 0x77,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x6f, 0x6f,
	0x6b, 0x4f, 0x66, 0x54, 0x68, 0x65, 0x59, 0x65, 0x61, 0x72, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x42, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x6e, 0x46, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0x02, 0x42, 0x19, 0x5a, 0x17, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_awards_proto_rawDescOnce sync.Once
	file_awards_proto_rawDescData = file_awards_proto_rawDesc
)

func file_awards_proto_rawDescGZIP() []byte {
	file_awards_proto_rawDescOnce.Do(func() {
		file_awards_proto_rawDescData = protoimpl.X.CompressGZIP(file_awards_proto_rawDescData)
	})
	return file_awards_proto_rawDescData
}

var file_awards_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_awards_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_awards_proto_goTypes = []interface{}{
	(AwardType)(0),            // 0: authors.AwardType
	(*Award)(nil),             // 1: authors.Award
	(*GetAwardsRequest)(nil),  // 2: authors.GetAwardsRequest
	(*GetAwardsResponse)(nil), // 3: authors.GetAwardsResponse
}
var file_awards_proto_depIdxs = []int32{
	0, // 0: authors.Award.type:type_name -> authors.AwardType
	1, // 1: authors.GetAwardsResponse.awards:type_name -> authors.Award
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_awards_proto_init() }
func file_awards_proto_init() {
	if File_awards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_awards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Award); i {
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
		file_awards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwardsRequest); i {
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
		file_awards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAwardsResponse); i {
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
			RawDescriptor: file_awards_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_awards_proto_goTypes,
		DependencyIndexes: file_awards_proto_depIdxs,
		EnumInfos:         file_awards_proto_enumTypes,
		MessageInfos:      file_awards_proto_msgTypes,
	}.Build()
	File_awards_proto = out.File
	file_awards_proto_rawDesc = nil
	file_awards_proto_goTypes = nil
	file_awards_proto_depIdxs = nil
}
