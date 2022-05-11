// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: authors.proto

package authors

import (
	graphql "github.com/kitt-technology/protoc-gen-graphql/graphql"
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

type GetAuthorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetAuthorsRequest) Reset() {
	*x = GetAuthorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsRequest) ProtoMessage() {}

func (x *GetAuthorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorsRequest) Descriptor() ([]byte, []int) {
	return file_authors_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthorsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors  []*Author         `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	PageInfo *graphql.PageInfo `protobuf:"bytes,5,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *GetAuthorsResponse) Reset() {
	*x = GetAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsResponse) ProtoMessage() {}

func (x *GetAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_authors_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthorsResponse) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *GetAuthorsResponse) GetPageInfo() *graphql.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

type AuthorsBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results map[string]*Author `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AuthorsBatchResponse) Reset() {
	*x = AuthorsBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorsBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsBatchResponse) ProtoMessage() {}

func (x *AuthorsBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsBatchResponse.ProtoReflect.Descriptor instead.
func (*AuthorsBatchResponse) Descriptor() ([]byte, []int) {
	return file_authors_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorsBatchResponse) GetResults() map[string]*Author {
	if x != nil {
		return x.Results
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_authors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_authors_proto_rawDescGZIP(), []int{3}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_authors_proto protoreflect.FileDescriptor

var file_authors_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x74, 0x74, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x03, 0xc0, 0x44, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x6f, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x14, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x4b, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0xad, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x12, 0x47, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1a,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x1a, 0x12, 0x82, 0x44, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x35,
	0x30, 0x30, 0x35, 0x32, 0x42, 0x19, 0x5a, 0x17, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authors_proto_rawDescOnce sync.Once
	file_authors_proto_rawDescData = file_authors_proto_rawDesc
)

func file_authors_proto_rawDescGZIP() []byte {
	file_authors_proto_rawDescOnce.Do(func() {
		file_authors_proto_rawDescData = protoimpl.X.CompressGZIP(file_authors_proto_rawDescData)
	})
	return file_authors_proto_rawDescData
}

var file_authors_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_authors_proto_goTypes = []interface{}{
	(*GetAuthorsRequest)(nil),    // 0: authors.GetAuthorsRequest
	(*GetAuthorsResponse)(nil),   // 1: authors.GetAuthorsResponse
	(*AuthorsBatchResponse)(nil), // 2: authors.AuthorsBatchResponse
	(*Author)(nil),               // 3: authors.Author
	nil,                          // 4: authors.AuthorsBatchResponse.ResultsEntry
	(*graphql.PageInfo)(nil),     // 5: graphql.PageInfo
	(*graphql.BatchRequest)(nil), // 6: graphql.BatchRequest
}
var file_authors_proto_depIdxs = []int32{
	3, // 0: authors.GetAuthorsResponse.authors:type_name -> authors.Author
	5, // 1: authors.GetAuthorsResponse.page_info:type_name -> graphql.PageInfo
	4, // 2: authors.AuthorsBatchResponse.results:type_name -> authors.AuthorsBatchResponse.ResultsEntry
	3, // 3: authors.AuthorsBatchResponse.ResultsEntry.value:type_name -> authors.Author
	0, // 4: authors.Authors.getAuthors:input_type -> authors.GetAuthorsRequest
	6, // 5: authors.Authors.loadAuthors:input_type -> graphql.BatchRequest
	1, // 6: authors.Authors.getAuthors:output_type -> authors.GetAuthorsResponse
	2, // 7: authors.Authors.loadAuthors:output_type -> authors.AuthorsBatchResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_authors_proto_init() }
func file_authors_proto_init() {
	if File_authors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsRequest); i {
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
		file_authors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsResponse); i {
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
		file_authors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorsBatchResponse); i {
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
		file_authors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
			RawDescriptor: file_authors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authors_proto_goTypes,
		DependencyIndexes: file_authors_proto_depIdxs,
		MessageInfos:      file_authors_proto_msgTypes,
	}.Build()
	File_authors_proto = out.File
	file_authors_proto_rawDesc = nil
	file_authors_proto_goTypes = nil
	file_authors_proto_depIdxs = nil
}
