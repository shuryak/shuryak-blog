package grpc

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/controller/grpc/articles_grpc"
	"github.com/shuryak/shuryak-blog/internal/entity"
	"github.com/shuryak/shuryak-blog/internal/usecase"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

// Check for implementation
var _ articles_grpc.ArticlesServer = (*articlesGRPCServer)(nil)

type articlesGRPCServer struct {
	a usecase.Article
	l logger.Interface
	articles_grpc.UnimplementedArticlesServer
}

type Metadata struct {
	UserId   uint32
	Username string
	Role     string
}

func (g articlesGRPCServer) extractMetadata(ctx context.Context) (*Metadata, error) {
	var mtdt Metadata

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userId := md.Get("user_id"); len(userId) > 0 {
			u64, err := strconv.ParseUint(userId[0], 10, 32)
			if err != nil {
				return nil, fmt.Errorf("cannot read user id: %w", err)
			}
			mtdt.UserId = uint32(u64)
		} else {
			return nil, fmt.Errorf("no user id")
		}

		if username := md.Get("username"); len(username) > 0 {
			mtdt.Username = username[0]
		} else {
			return nil, fmt.Errorf("no username")
		}

		if role := md.Get("role"); len(role) > 0 {
			mtdt.Role = role[0]
		} else {
			return nil, fmt.Errorf("no role")
		}
	}

	return &mtdt, nil
}

func NewArticlesGrpcServer(server *grpc.Server, a usecase.Article, l logger.Interface) {
	ags := &articlesGRPCServer{
		a: a,
		l: l,
	}
	articles_grpc.RegisterArticlesServer(server, ags)
	reflection.Register(server)
}

func (g articlesGRPCServer) Create(ctx context.Context, request *articles_grpc.CreateRequest) (
	*articles_grpc.SingleArticleResponse,
	error,
) {
	mtdt, err := g.extractMetadata(ctx)
	if err != nil {
		g.l.Error(err, "grpc - extractMetadata: %w", err)
		return nil, status.Error(codes.Unauthenticated, "authorization error")
	}

	article, err := g.a.Create(ctx, entity.Article{
		CustomId:  request.CustomId,
		AuthorId:  mtdt.UserId,
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		Content:   request.Content.AsMap(),
	})
	if err != nil {
		g.l.Error(err, "grpc - Create")
		return nil, fmt.Errorf("grpc - Create - g.a.Create: %w", err) // TODO: do errors?
	}

	content, err := structpb.NewStruct(article.Content)
	if err != nil {
		g.l.Error(err, "grpc - Create")
		return nil, fmt.Errorf("grpc - Create - structpb.NewStruct: %w", err) // TODO: do errors?
	}

	sar := articles_grpc.SingleArticleResponse{
		Id:        article.Id,
		CustomId:  article.CustomId,
		AuthorId:  article.AuthorId,
		Title:     article.Title,
		Thumbnail: article.Thumbnail,
		Content:   content,
		CreatedAt: timestamppb.New(article.CreatedAt),
	}

	return &sar, nil
}

func (g articlesGRPCServer) GetById(ctx context.Context, request *articles_grpc.ArticleIdRequest) (
	*articles_grpc.SingleArticleResponse,
	error,
) {
	a, err := g.a.GetById(ctx, uint(request.Id))
	if err != nil {
		g.l.Error(err, "grpc - GetById")
		return nil, fmt.Errorf("grpc - GetById - g.a.GetById: %w", err) // TODO: do errors?
	}

	content, err := structpb.NewStruct(a.Content)
	if err != nil {
		g.l.Error(err, "grpc - GetById")
		return nil, fmt.Errorf("grpc - GetById - structpb.NewStruct: %w", err) // TODO: do errors?
	}

	sar := articles_grpc.SingleArticleResponse{
		Id:        a.Id,
		CustomId:  a.CustomId,
		AuthorId:  a.AuthorId,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   content,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}

	return &sar, nil
}

func (g articlesGRPCServer) GetMany(ctx context.Context, request *articles_grpc.GetManyRequest) (
	*articles_grpc.MultipleArticlesResponse,
	error,
) {
	articles, err := g.a.GetMany(ctx, uint(request.Offset), uint(request.Count))
	if err != nil {
		g.l.Error(err, "grpc - GetMany")
		return nil, fmt.Errorf("grpc - GetMany - g.a.GetMany: %w", err) // TODO: do errors?
	}

	sarSlice := make([]*articles_grpc.SingleArticleResponse, len(articles))
	for i, a := range articles {
		content, err := structpb.NewStruct(a.Content)
		if err != nil {
			g.l.Error(err, "grpc - GetMany")
			return nil, fmt.Errorf("grpc - GetMany - structpb.NewStruct: %w", err) // TODO: do errors?
		}

		sarSlice[i] = &articles_grpc.SingleArticleResponse{
			Id:        a.Id,
			CustomId:  a.CustomId,
			AuthorId:  a.AuthorId,
			Title:     a.Title,
			Thumbnail: a.Thumbnail,
			Content:   content,
			CreatedAt: timestamppb.New(a.CreatedAt),
		}
	}

	return &articles_grpc.MultipleArticlesResponse{Articles: sarSlice}, nil
}

func (g articlesGRPCServer) Update(ctx context.Context, request *articles_grpc.UpdateRequest) (
	*articles_grpc.SingleArticleResponse,
	error,
) {
	mtdt, err := g.extractMetadata(ctx)
	if err != nil {
		g.l.Error(err, "grpc - extractMetadata: %w", err)
		return nil, status.Error(codes.Unauthenticated, "authorization error")
	}

	a, err := g.a.Update(ctx, entity.Article{
		Id:        request.Id,
		CustomId:  request.CustomId,
		AuthorId:  mtdt.UserId,
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		Content:   request.Content.AsMap(),
	})
	if err != nil {
		g.l.Error(err, "grpc - Update")
		return nil, fmt.Errorf("grpc - Update - g.a.Update: %w", err) // TODO: do errors?
	}

	content, err := structpb.NewStruct(a.Content)
	if err != nil {
		g.l.Error(err, "grpc - Update")
		return nil, fmt.Errorf("grpc - Update - structpb.NewStruct: %w", err) // TODO: do errors?
	}

	sar := articles_grpc.SingleArticleResponse{
		Id:        a.Id,
		CustomId:  a.CustomId,
		AuthorId:  a.AuthorId,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   content,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}

	return &sar, nil
}

func (g articlesGRPCServer) Delete(ctx context.Context, request *articles_grpc.ArticleIdRequest) (
	*articles_grpc.SingleArticleResponse,
	error,
) {
	// TODO: authorize

	a, err := g.a.Delete(ctx, uint(request.Id))
	if err != nil {
		g.l.Error(err, "grpc - Delete")
		return nil, fmt.Errorf("grpc - Delete - g.a.Delete: %w", err) // TODO: do errors?
	}

	content, err := structpb.NewStruct(a.Content)
	if err != nil {
		g.l.Error(err, "grpc - Delete")
		return nil, fmt.Errorf("grpc - Delete - structpb.NewStruct: %w", err) // TODO: do errors?
	}

	sar := articles_grpc.SingleArticleResponse{
		Id:        a.Id,
		CustomId:  a.CustomId,
		AuthorId:  a.AuthorId,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   content,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}

	return &sar, nil
}
