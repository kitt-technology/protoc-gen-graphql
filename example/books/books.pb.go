// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: books.proto

package books

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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

	Ids               []string                `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	HardbackOnly      *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=hardback_only,json=hardbackOnly,proto3" json:"hardback_only,omitempty"`
	Price             *wrapperspb.FloatValue  `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Genres            []Genre                 `protobuf:"varint,4,rep,packed,name=genres,proto3,enum=books.Genre" json:"genres,omitempty"`
	ReleasedAfter     *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=released_after,json=releasedAfter,proto3" json:"released_after,omitempty"`
	PriceGreaterThan  float32                 `protobuf:"fixed32,6,opt,name=price_greater_than,json=priceGreaterThan,proto3" json:"price_greater_than,omitempty"`
	CopiesGreaterThan int64                   `protobuf:"varint,7,opt,name=copies_greater_than,json=copiesGreaterThan,proto3" json:"copies_greater_than,omitempty"`
	CopiesLessThan    int32                   `protobuf:"varint,8,opt,name=copies_less_than,json=copiesLessThan,proto3" json:"copies_less_than,omitempty"`
	PriceLessThan     float64                 `protobuf:"fixed64,9,opt,name=price_less_than,json=priceLessThan,proto3" json:"price_less_than,omitempty"`
	FooBar            *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=foo_bar,json=fooBar,proto3" json:"foo_bar,omitempty"`
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

func (x *GetBooksRequest) GetPrice() *wrapperspb.FloatValue {
	if x != nil {
		return x.Price
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

func (x *GetBooksRequest) GetPriceGreaterThan() float32 {
	if x != nil {
		return x.PriceGreaterThan
	}
	return 0
}

func (x *GetBooksRequest) GetCopiesGreaterThan() int64 {
	if x != nil {
		return x.CopiesGreaterThan
	}
	return 0
}

func (x *GetBooksRequest) GetCopiesLessThan() int32 {
	if x != nil {
		return x.CopiesLessThan
	}
	return 0
}

func (x *GetBooksRequest) GetPriceLessThan() float64 {
	if x != nil {
		return x.PriceLessThan
	}
	return 0
}

func (x *GetBooksRequest) GetFooBar() *wrapperspb.StringValue {
	if x != nil {
		return x.FooBar
	}
	return nil
}

type GetBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books  []*Book                 `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	Foobar *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=foobar,proto3" json:"foobar,omitempty"`
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

func (x *GetBooksResponse) GetFoobar() *wrapperspb.StringValue {
	if x != nil {
		return x.Foobar
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

type BooksByAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Book `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *BooksByAuthor) Reset() {
	*x = BooksByAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksByAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksByAuthor) ProtoMessage() {}

func (x *BooksByAuthor) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BooksByAuthor.ProtoReflect.Descriptor instead.
func (*BooksByAuthor) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{4}
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
	Copies      int64                  `protobuf:"varint,7,opt,name=copies,proto3" json:"copies,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{5}
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
	if x != nil {
		return x.Copies
	}
	return 0
}

var File_books_proto protoreflect.FileDescriptor

var file_books_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x69, 0x74, 0x74, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x44, 0x6f,
	0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0xfd, 0x03, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x44, 0x01, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x68, 0x61, 0x72, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x4f,
	0x6e, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0e,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x68, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x12, 0x2e, 0x0a,
	0x13, 0x63, 0x6f, 0x70, 0x69, 0x65, 0x73, 0x5f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f,
	0x74, 0x68, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x63, 0x6f, 0x70, 0x69,
	0x65, 0x73, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x12, 0x28, 0x0a,
	0x10, 0x63, 0x6f, 0x70, 0x69, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x61,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x70, 0x69, 0x65, 0x73, 0x4c,
	0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x61, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x12,
	0x35, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x5f, 0x62, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x66, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x3a, 0x0f, 0xea, 0x43, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x34,
	0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x66, 0x6f,
	0x6f, 0x62, 0x61, 0x72, 0x22, 0xb4, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x36, 0x0a, 0x0d, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x22, 0xd8, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x70, 0x69, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x70, 0x69, 0x65, 0x73, 0x2a, 0x23,
	0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x79, 0x10, 0x01, 0x32, 0xdb, 0x01, 0x0a, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x3d, 0x0a,
	0x08, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09,
	0x64, 0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2e, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x10, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x2e, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x1a, 0x12, 0x82,
	0x44, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x35, 0x30, 0x30, 0x35,
	0x31, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_books_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_books_proto_goTypes = []interface{}{
	(Genre)(0),                       // 0: books.Genre
	(*DoNothing)(nil),                // 1: books.DoNothing
	(*GetBooksRequest)(nil),          // 2: books.GetBooksRequest
	(*GetBooksResponse)(nil),         // 3: books.GetBooksResponse
	(*GetBooksByAuthorResponse)(nil), // 4: books.GetBooksByAuthorResponse
	(*BooksByAuthor)(nil),            // 5: books.BooksByAuthor
	(*Book)(nil),                     // 6: books.Book
	nil,                              // 7: books.GetBooksByAuthorResponse.ResultsEntry
	(*wrapperspb.BoolValue)(nil),     // 8: google.protobuf.BoolValue
	(*wrapperspb.FloatValue)(nil),    // 9: google.protobuf.FloatValue
	(*timestamppb.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil),   // 11: google.protobuf.StringValue
	(*graphql.BatchRequest)(nil),     // 12: graphql.BatchRequest
}
var file_books_proto_depIdxs = []int32{
	8,  // 0: books.GetBooksRequest.hardback_only:type_name -> google.protobuf.BoolValue
	9,  // 1: books.GetBooksRequest.price:type_name -> google.protobuf.FloatValue
	0,  // 2: books.GetBooksRequest.genres:type_name -> books.Genre
	10, // 3: books.GetBooksRequest.released_after:type_name -> google.protobuf.Timestamp
	11, // 4: books.GetBooksRequest.foo_bar:type_name -> google.protobuf.StringValue
	6,  // 5: books.GetBooksResponse.books:type_name -> books.Book
	11, // 6: books.GetBooksResponse.foobar:type_name -> google.protobuf.StringValue
	7,  // 7: books.GetBooksByAuthorResponse.results:type_name -> books.GetBooksByAuthorResponse.ResultsEntry
	6,  // 8: books.BooksByAuthor.results:type_name -> books.Book
	0,  // 9: books.Book.genre:type_name -> books.Genre
	10, // 10: books.Book.release_date:type_name -> google.protobuf.Timestamp
	5,  // 11: books.GetBooksByAuthorResponse.ResultsEntry.value:type_name -> books.BooksByAuthor
	2,  // 12: books.Books.getBooks:input_type -> books.GetBooksRequest
	1,  // 13: books.Books.doNothing:input_type -> books.DoNothing
	12, // 14: books.Books.getBooksByAuthor:input_type -> graphql.BatchRequest
	3,  // 15: books.Books.getBooks:output_type -> books.GetBooksResponse
	1,  // 16: books.Books.doNothing:output_type -> books.DoNothing
	4,  // 17: books.Books.getBooksByAuthor:output_type -> books.GetBooksByAuthorResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
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
		file_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_books_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
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
