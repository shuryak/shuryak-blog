// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: internal/controller/grpc/articles.proto

package articles_grpc

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

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type articlesClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesClient(cc grpc.ClientConnInterface) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/articlesGrpc.Articles/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
// All implementations must embed UnimplementedArticlesServer
// for forward compatibility
type ArticlesServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedArticlesServer()
}

// UnimplementedArticlesServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesServer struct {
}

func (UnimplementedArticlesServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArticlesServer) mustEmbedUnimplementedArticlesServer() {}

// UnsafeArticlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesServer will
// result in compilation errors.
type UnsafeArticlesServer interface {
	mustEmbedUnimplementedArticlesServer()
}

func RegisterArticlesServer(s grpc.ServiceRegistrar, srv ArticlesServer) {
	s.RegisterService(&Articles_ServiceDesc, srv)
}

func _Articles_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlesGrpc.Articles/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Articles_ServiceDesc is the grpc.ServiceDesc for Articles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Articles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "articlesGrpc.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Articles_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/controller/grpc/articles.proto",
}
