// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: book.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BookManagement_ListBooks_FullMethodName  = "/BookManagement/ListBooks"
	BookManagement_CreateBook_FullMethodName = "/BookManagement/CreateBook"
)

// BookManagementClient is the client API for BookManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookManagementClient interface {
	ListBooks(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListBookResponse, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
}

type bookManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewBookManagementClient(cc grpc.ClientConnInterface) BookManagementClient {
	return &bookManagementClient{cc}
}

func (c *bookManagementClient) ListBooks(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListBookResponse, error) {
	out := new(ListBookResponse)
	err := c.cc.Invoke(ctx, BookManagement_ListBooks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookManagementClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, BookManagement_CreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookManagementServer is the server API for BookManagement service.
// All implementations must embed UnimplementedBookManagementServer
// for forward compatibility
type BookManagementServer interface {
	ListBooks(context.Context, *EmptyRequest) (*ListBookResponse, error)
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	mustEmbedUnimplementedBookManagementServer()
}

// UnimplementedBookManagementServer must be embedded to have forward compatible implementations.
type UnimplementedBookManagementServer struct {
}

func (UnimplementedBookManagementServer) ListBooks(context.Context, *EmptyRequest) (*ListBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookManagementServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookManagementServer) mustEmbedUnimplementedBookManagementServer() {}

// UnsafeBookManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookManagementServer will
// result in compilation errors.
type UnsafeBookManagementServer interface {
	mustEmbedUnimplementedBookManagementServer()
}

func RegisterBookManagementServer(s grpc.ServiceRegistrar, srv BookManagementServer) {
	s.RegisterService(&BookManagement_ServiceDesc, srv)
}

func _BookManagement_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookManagementServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookManagement_ListBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookManagementServer).ListBooks(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookManagement_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookManagementServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookManagement_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookManagementServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookManagement_ServiceDesc is the grpc.ServiceDesc for BookManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookManagement",
	HandlerType: (*BookManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookManagement_ListBooks_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookManagement_CreateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
