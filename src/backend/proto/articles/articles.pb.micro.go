// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/articles/articles.proto

package articles

import (
	fmt "fmt"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Articles service

func NewArticlesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Articles service

type ArticlesService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*SingleArticleResponse, error)
	GetByCustomId(ctx context.Context, in *ArticleCustomIdRequest, opts ...client.CallOption) (*SingleArticleResponse, error)
	GetMany(ctx context.Context, in *GetManyRequest, opts ...client.CallOption) (*ShortArticlesResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*SingleArticleResponse, error)
	Delete(ctx context.Context, in *ArticleCustomIdRequest, opts ...client.CallOption) (*SingleArticleResponse, error)
}

type articlesService struct {
	c    client.Client
	name string
}

func NewArticlesService(name string, c client.Client) ArticlesService {
	return &articlesService{
		c:    c,
		name: name,
	}
}

func (c *articlesService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*SingleArticleResponse, error) {
	req := c.c.NewRequest(c.name, "Articles.Create", in)
	out := new(SingleArticleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesService) GetByCustomId(ctx context.Context, in *ArticleCustomIdRequest, opts ...client.CallOption) (*SingleArticleResponse, error) {
	req := c.c.NewRequest(c.name, "Articles.GetByCustomId", in)
	out := new(SingleArticleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesService) GetMany(ctx context.Context, in *GetManyRequest, opts ...client.CallOption) (*ShortArticlesResponse, error) {
	req := c.c.NewRequest(c.name, "Articles.GetMany", in)
	out := new(ShortArticlesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*SingleArticleResponse, error) {
	req := c.c.NewRequest(c.name, "Articles.Update", in)
	out := new(SingleArticleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesService) Delete(ctx context.Context, in *ArticleCustomIdRequest, opts ...client.CallOption) (*SingleArticleResponse, error) {
	req := c.c.NewRequest(c.name, "Articles.Delete", in)
	out := new(SingleArticleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Articles service

type ArticlesHandler interface {
	Create(context.Context, *CreateRequest, *SingleArticleResponse) error
	GetByCustomId(context.Context, *ArticleCustomIdRequest, *SingleArticleResponse) error
	GetMany(context.Context, *GetManyRequest, *ShortArticlesResponse) error
	Update(context.Context, *UpdateRequest, *SingleArticleResponse) error
	Delete(context.Context, *ArticleCustomIdRequest, *SingleArticleResponse) error
}

func RegisterArticlesHandler(s server.Server, hdlr ArticlesHandler, opts ...server.HandlerOption) error {
	type articles interface {
		Create(ctx context.Context, in *CreateRequest, out *SingleArticleResponse) error
		GetByCustomId(ctx context.Context, in *ArticleCustomIdRequest, out *SingleArticleResponse) error
		GetMany(ctx context.Context, in *GetManyRequest, out *ShortArticlesResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *SingleArticleResponse) error
		Delete(ctx context.Context, in *ArticleCustomIdRequest, out *SingleArticleResponse) error
	}
	type Articles struct {
		articles
	}
	h := &articlesHandler{hdlr}
	return s.Handle(s.NewHandler(&Articles{h}, opts...))
}

type articlesHandler struct {
	ArticlesHandler
}

func (h *articlesHandler) Create(ctx context.Context, in *CreateRequest, out *SingleArticleResponse) error {
	return h.ArticlesHandler.Create(ctx, in, out)
}

func (h *articlesHandler) GetByCustomId(ctx context.Context, in *ArticleCustomIdRequest, out *SingleArticleResponse) error {
	return h.ArticlesHandler.GetByCustomId(ctx, in, out)
}

func (h *articlesHandler) GetMany(ctx context.Context, in *GetManyRequest, out *ShortArticlesResponse) error {
	return h.ArticlesHandler.GetMany(ctx, in, out)
}

func (h *articlesHandler) Update(ctx context.Context, in *UpdateRequest, out *SingleArticleResponse) error {
	return h.ArticlesHandler.Update(ctx, in, out)
}

func (h *articlesHandler) Delete(ctx context.Context, in *ArticleCustomIdRequest, out *SingleArticleResponse) error {
	return h.ArticlesHandler.Delete(ctx, in, out)
}
