// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: books.proto

package books

import (
	common_example "github.com/kitt-technology/protoc-gen-graphql/example/common-example"
	graphql "github.com/kitt-technology/protoc-gen-graphql/graphql"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Genre int32

const (
	Genre_Fiction   Genre = 0
	Genre_Biography Genre = 1
)

// Enum value maps for Genre.
var (
	Genre_name = map[int32]string{
		0: "Fiction",
		1: "Biography",
	}
	Genre_value = map[string]int32{
		"Fiction":   0,
		"Biography": 1,
	}
)

func (x Genre) Enum() *Genre {
	p := new(Genre)
	*p = x
	return p
}

func (x Genre) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Genre) Descriptor() protoreflect.EnumDescriptor {
	return file_books_proto_enumTypes[0].Descriptor()
}

func (Genre) Type() protoreflect.EnumType {
	return &file_books_proto_enumTypes[0]
}

func (x Genre) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Genre.Descriptor instead.
func (Genre) EnumDescriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{0}
}

type DoNothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DoNothing) Reset() {
	*x = DoNothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoNothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoNothing) ProtoMessage() {}

func (x *DoNothing) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoNothing.ProtoReflect.Descriptor instead.
func (*DoNothing) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{0}
}

type GetBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids           []string               `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	HardbackOnly  *wrapperspb.BoolValue  `protobuf:"bytes,2,opt,name=hardback_only,json=hardbackOnly,proto3" json:"hardback_only,omitempty"`
	Genres        []Genre                `protobuf:"varint,3,rep,packed,name=genres,proto3,enum=books.Genre" json:"genres,omitempty"`
	ReleasedAfter *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=released_after,json=releasedAfter,proto3" json:"released_after,omitempty"`
	IgnoreMe      string                 `protobuf:"bytes,5,opt,name=ignore_me,json=ignoreMe,proto3" json:"ignore_me,omitempty"`
}

func (x *GetBooksRequest) Reset() {
	*x = GetBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksRequest) ProtoMessage() {}

func (x *GetBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksRequest.ProtoReflect.Descriptor instead.
func (*GetBooksRequest) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{1}
}

func (x *GetBooksRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetBooksRequest) GetHardbackOnly() *wrapperspb.BoolValue {
	if x != nil {
		return x.HardbackOnly
	}
	return nil
}

func (x *GetBooksRequest) GetGenres() []Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *GetBooksRequest) GetReleasedAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleasedAfter
	}
	return nil
}

func (x *GetBooksRequest) GetIgnoreMe() string {
	if x != nil {
		return x.IgnoreMe
	}
	return ""
}

type GetBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBooksResponse) Reset() {
	*x = GetBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksResponse) ProtoMessage() {}

func (x *GetBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksResponse.ProtoReflect.Descriptor instead.
func (*GetBooksResponse) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{2}
}

func (x *GetBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type GetBooksByAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results map[string]*BooksByAuthor `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetBooksByAuthorResponse) Reset() {
	*x = GetBooksByAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksByAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksByAuthorResponse) ProtoMessage() {}

func (x *GetBooksByAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksByAuthorResponse.ProtoReflect.Descriptor instead.
func (*GetBooksByAuthorResponse) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{3}
}

func (x *GetBooksByAuthorResponse) GetResults() map[string]*BooksByAuthor {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetBooksBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reqs []*GetBooksRequest `protobuf:"bytes,1,rep,name=reqs,proto3" json:"reqs,omitempty"`
}

func (x *GetBooksBatchRequest) Reset() {
	*x = GetBooksBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksBatchRequest) ProtoMessage() {}

func (x *GetBooksBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksBatchRequest.ProtoReflect.Descriptor instead.
func (*GetBooksBatchRequest) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{4}
}

func (x *GetBooksBatchRequest) GetReqs() []*GetBooksRequest {
	if x != nil {
		return x.Reqs
	}
	return nil
}

type GetBooksBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results map[string]*GetBooksResponse `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetBooksBatchResponse) Reset() {
	*x = GetBooksBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksBatchResponse) ProtoMessage() {}

func (x *GetBooksBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksBatchResponse.ProtoReflect.Descriptor instead.
func (*GetBooksBatchResponse) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{5}
}

func (x *GetBooksBatchResponse) GetResults() map[string]*GetBooksResponse {
	if x != nil {
		return x.Results
	}
	return nil
}

type BooksByAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Book `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *BooksByAuthor) Reset() {
	*x = BooksByAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksByAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksByAuthor) ProtoMessage() {}

func (x *BooksByAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksByAuthor.ProtoReflect.Descriptor instead.
func (*BooksByAuthor) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{6}
}

func (x *BooksByAuthor) GetResults() []*Book {
	if x != nil {
		return x.Results
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AuthorId    string                 `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Genre       Genre                  `protobuf:"varint,4,opt,name=genre,proto3,enum=books.Genre" json:"genre,omitempty"`
	ReleaseDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Price       float32                `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Copies      *int64                 `protobuf:"varint,7,opt,name=copies,proto3,oneof" json:"copies,omitempty"`
	PriceTwo    *common_example.Money  `protobuf:"bytes,8,opt,name=price_two,json=priceTwo,proto3" json:"price_two,omitempty"`
	IsSigned    *wrapperspb.BoolValue  `protobuf:"bytes,9,opt,name=is_signed,json=isSigned,proto3" json:"is_signed,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{7}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Book) GetGenre() Genre {
	if x != nil {
		return x.Genre
	}
	return Genre_Fiction
}

func (x *Book) GetReleaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseDate
	}
	return nil
}

func (x *Book) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Book) GetCopies() int64 {
	if x != nil && x.Copies != nil {
		return *x.Copies
	}
	return 0
}

func (x *Book) GetPriceTwo() *common_example.Money {
	if x != nil {
		return x.PriceTwo
	}
	return nil
}

func (x *Book) GetIsSigned() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsSigned
	}
	return nil
}

type SkipMe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OhNo string `protobuf:"bytes,1,opt,name=oh_no,json=ohNo,proto3" json:"oh_no,omitempty"`
}

func (x *SkipMe) Reset() {
	*x = SkipMe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkipMe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkipMe) ProtoMessage() {}

func (x *SkipMe) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkipMe.ProtoReflect.Descriptor instead.
func (*SkipMe) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{8}
}

func (x *SkipMe) GetOhNo() string {
	if x != nil {
		return x.OhNo
	}
	return ""
}

var File_books_proto protoreflect.FileDescriptor

var file_books_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x69, 0x74, 0x74, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x74, 0x74, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x22, 0x85, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x44, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x3f, 0x0a,
	0x0d, 0x68, 0x61, 0x72, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x68, 0x61, 0x72, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x24,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64,
	0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x5f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc8, 0x44, 0x01, 0x52,
	0x08, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x4d, 0x65, 0x3a, 0x0f, 0xea, 0x43, 0x0c, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x22, 0xb4, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x42, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x04, 0x72, 0x65, 0x71, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x72, 0x65, 0x71, 0x73, 0x22, 0xb1, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x53, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x36, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xd5, 0x02, 0x0a, 0x04, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x63,
	0x6f, 0x70, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x63,
	0x6f, 0x70, 0x69, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x77, 0x6f, 0x12, 0x37, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x73, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x70, 0x69, 0x65, 0x73,
	0x22, 0x22, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x70, 0x4d, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x6f, 0x68,
	0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x68, 0x4e, 0x6f, 0x3a,
	0x03, 0xf8, 0x43, 0x01, 0x2a, 0x23, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x69,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x10, 0x01, 0x32, 0xf9, 0x01, 0x0a, 0x05, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12,
	0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0xd0, 0x44,
	0x01, 0x1a, 0x12, 0x82, 0x44, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a,
	0x35, 0x30, 0x30, 0x35, 0x31, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_proto_rawDescOnce sync.Once
	file_books_proto_rawDescData = file_books_proto_rawDesc
)

func file_books_proto_rawDescGZIP() []byte {
	file_books_proto_rawDescOnce.Do(func() {
		file_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_proto_rawDescData)
	})
	return file_books_proto_rawDescData
}

var file_books_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_books_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_books_proto_goTypes = []interface{}{
	(Genre)(0),                       // 0: books.Genre
	(*DoNothing)(nil),                // 1: books.DoNothing
	(*GetBooksRequest)(nil),          // 2: books.GetBooksRequest
	(*GetBooksResponse)(nil),         // 3: books.GetBooksResponse
	(*GetBooksByAuthorResponse)(nil), // 4: books.GetBooksByAuthorResponse
	(*GetBooksBatchRequest)(nil),     // 5: books.GetBooksBatchRequest
	(*GetBooksBatchResponse)(nil),    // 6: books.GetBooksBatchResponse
	(*BooksByAuthor)(nil),            // 7: books.BooksByAuthor
	(*Book)(nil),                     // 8: books.Book
	(*SkipMe)(nil),                   // 9: books.SkipMe
	nil,                              // 10: books.GetBooksByAuthorResponse.ResultsEntry
	nil,                              // 11: books.GetBooksBatchResponse.ResultsEntry
	(*wrapperspb.BoolValue)(nil),     // 12: google.protobuf.BoolValue
	(*timestamppb.Timestamp)(nil),    // 13: google.protobuf.Timestamp
	(*common_example.Money)(nil),     // 14: common_example.Money
	(*graphql.BatchRequest)(nil),     // 15: graphql.BatchRequest
}
var file_books_proto_depIdxs = []int32{
	12, // 0: books.GetBooksRequest.hardback_only:type_name -> google.protobuf.BoolValue
	0,  // 1: books.GetBooksRequest.genres:type_name -> books.Genre
	13, // 2: books.GetBooksRequest.released_after:type_name -> google.protobuf.Timestamp
	8,  // 3: books.GetBooksResponse.books:type_name -> books.Book
	10, // 4: books.GetBooksByAuthorResponse.results:type_name -> books.GetBooksByAuthorResponse.ResultsEntry
	2,  // 5: books.GetBooksBatchRequest.reqs:type_name -> books.GetBooksRequest
	11, // 6: books.GetBooksBatchResponse.results:type_name -> books.GetBooksBatchResponse.ResultsEntry
	8,  // 7: books.BooksByAuthor.results:type_name -> books.Book
	0,  // 8: books.Book.genre:type_name -> books.Genre
	13, // 9: books.Book.release_date:type_name -> google.protobuf.Timestamp
	14, // 10: books.Book.price_two:type_name -> common_example.Money
	12, // 11: books.Book.is_signed:type_name -> google.protobuf.BoolValue
	7,  // 12: books.GetBooksByAuthorResponse.ResultsEntry.value:type_name -> books.BooksByAuthor
	3,  // 13: books.GetBooksBatchResponse.ResultsEntry.value:type_name -> books.GetBooksResponse
	2,  // 14: books.Books.getBooks:input_type -> books.GetBooksRequest
	15, // 15: books.Books.getBooksByAuthor:input_type -> graphql.BatchRequest
	5,  // 16: books.Books.getBooksBatch:input_type -> books.GetBooksBatchRequest
	3,  // 17: books.Books.getBooks:output_type -> books.GetBooksResponse
	4,  // 18: books.Books.getBooksByAuthor:output_type -> books.GetBooksByAuthorResponse
	6,  // 19: books.Books.getBooksBatch:output_type -> books.GetBooksBatchResponse
	17, // [17:20] is the sub-list for method output_type
	14, // [14:17] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_books_proto_init() }
func file_books_proto_init() {
	if File_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoNothing); i {
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
		file_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksRequest); i {
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
		file_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksResponse); i {
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
		file_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksByAuthorResponse); i {
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
		file_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksBatchRequest); i {
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
		file_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksBatchResponse); i {
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
		file_books_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksByAuthor); i {
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
		file_books_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_books_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkipMe); i {
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
	file_books_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_books_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_proto_goTypes,
		DependencyIndexes: file_books_proto_depIdxs,
		EnumInfos:         file_books_proto_enumTypes,
		MessageInfos:      file_books_proto_msgTypes,
	}.Build()
	File_books_proto = out.File
	file_books_proto_rawDesc = nil
	file_books_proto_goTypes = nil
	file_books_proto_depIdxs = nil
}
