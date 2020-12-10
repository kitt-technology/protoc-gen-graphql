// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authors

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuthorsClient is the client API for Authors service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorsClient interface {
	GetAuthors(ctx context.Context, in *GetAuthorsRequest, opts ...grpc.CallOption) (*GetAuthorsResponse, error)
}

type authorsClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorsClient(cc grpc.ClientConnInterface) AuthorsClient {
	return &authorsClient{cc}
}

func (c *authorsClient) GetAuthors(ctx context.Context, in *GetAuthorsRequest, opts ...grpc.CallOption) (*GetAuthorsResponse, error) {
	out := new(GetAuthorsResponse)
	err := c.cc.Invoke(ctx, "/authors.Authors/getAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorsServer is the server API for Authors service.
// All implementations must embed UnimplementedAuthorsServer
// for forward compatibility
type AuthorsServer interface {
	GetAuthors(context.Context, *GetAuthorsRequest) (*GetAuthorsResponse, error)
	mustEmbedUnimplementedAuthorsServer()
}

// UnimplementedAuthorsServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorsServer struct {
}

func (UnimplementedAuthorsServer) GetAuthors(context.Context, *GetAuthorsRequest) (*GetAuthorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthors not implemented")
}
func (UnimplementedAuthorsServer) mustEmbedUnimplementedAuthorsServer() {}

// UnsafeAuthorsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorsServer will
// result in compilation errors.
type UnsafeAuthorsServer interface {
	mustEmbedUnimplementedAuthorsServer()
}

func RegisterAuthorsServer(s grpc.ServiceRegistrar, srv AuthorsServer) {
	s.RegisterService(&_Authors_serviceDesc, srv)
}

func _Authors_GetAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorsServer).GetAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authors.Authors/getAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorsServer).GetAuthors(ctx, req.(*GetAuthorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authors_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authors.Authors",
	HandlerType: (*AuthorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAuthors",
			Handler:    _Authors_GetAuthors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authors.proto",
}