// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: internal/articles/pb/articles.proto

package pb

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
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	GetById(ctx context.Context, in *ArticleIdRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	GetMany(ctx context.Context, in *GetManyRequest, opts ...grpc.CallOption) (*MultipleArticlesResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
	Delete(ctx context.Context, in *ArticleIdRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error)
}

type articlesClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesClient(cc grpc.ClientConnInterface) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/articlesGrpc.Articles/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) GetById(ctx context.Context, in *ArticleIdRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/articlesGrpc.Articles/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) GetMany(ctx context.Context, in *GetManyRequest, opts ...grpc.CallOption) (*MultipleArticlesResponse, error) {
	out := new(MultipleArticlesResponse)
	err := c.cc.Invoke(ctx, "/articlesGrpc.Articles/GetMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/articlesGrpc.Articles/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Delete(ctx context.Context, in *ArticleIdRequest, opts ...grpc.CallOption) (*SingleArticleResponse, error) {
	out := new(SingleArticleResponse)
	err := c.cc.Invoke(ctx, "/articlesGrpc.Articles/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
// All implementations must embed UnimplementedArticlesServer
// for forward compatibility
type ArticlesServer interface {
	Create(context.Context, *CreateRequest) (*SingleArticleResponse, error)
	GetById(context.Context, *ArticleIdRequest) (*SingleArticleResponse, error)
	GetMany(context.Context, *GetManyRequest) (*MultipleArticlesResponse, error)
	Update(context.Context, *UpdateRequest) (*SingleArticleResponse, error)
	Delete(context.Context, *ArticleIdRequest) (*SingleArticleResponse, error)
	mustEmbedUnimplementedArticlesServer()
}

// UnimplementedArticlesServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesServer struct {
}

func (UnimplementedArticlesServer) Create(context.Context, *CreateRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArticlesServer) GetById(context.Context, *ArticleIdRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedArticlesServer) GetMany(context.Context, *GetManyRequest) (*MultipleArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedArticlesServer) Update(context.Context, *UpdateRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArticlesServer) Delete(context.Context, *ArticleIdRequest) (*SingleArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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

func _Articles_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlesGrpc.Articles/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).GetById(ctx, req.(*ArticleIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlesGrpc.Articles/GetMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).GetMany(ctx, req.(*GetManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlesGrpc.Articles/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlesGrpc.Articles/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Delete(ctx, req.(*ArticleIdRequest))
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
		{
			MethodName: "GetById",
			Handler:    _Articles_GetById_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _Articles_GetMany_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Articles_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Articles_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/articles/pb/articles.proto",
}