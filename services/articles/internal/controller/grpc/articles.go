package grpc

import (
	"context"
	"fmt"
	articles_grpc "github.com/shuryak/shuryak-blog/internal/controller/grpc/articles_grpc"
	"github.com/shuryak/shuryak-blog/internal/entity"
	"github.com/shuryak/shuryak-blog/internal/usecase"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Check for implementation
var _ articles_grpc.ArticlesServer = (*articlesGRPCServer)(nil)

type articlesGRPCServer struct {
	a usecase.Article
	l logger.Interface
	articles_grpc.UnimplementedArticlesServer
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
	article, err := g.a.Create(ctx, entity.Article{ // TODO: context things
		CustomId:  request.CustomId,
		AuthorId:  uint(request.AuthorId),
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
		Id:        uint32(article.Id),
		CustomId:  article.CustomId,
		AuthorId:  uint32(article.AuthorId),
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
	a, err := g.a.GetById(ctx, uint(request.Id)) // TODO: context things
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
		Id:        uint32(a.Id),
		CustomId:  a.CustomId,
		AuthorId:  uint32(a.AuthorId),
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
	articles, err := g.a.GetMany(ctx, uint(request.Offset), uint(request.Count)) // TODO: context things
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
			Id:        uint32(a.Id),
			CustomId:  a.CustomId,
			AuthorId:  uint32(a.AuthorId),
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
	a, err := g.a.Update(ctx, entity.Article{ // TODO: context things
		Id:        uint(request.Id),
		CustomId:  request.CustomId,
		AuthorId:  uint(request.AuthorId),
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
		Id:        uint32(a.Id),
		CustomId:  a.CustomId,
		AuthorId:  uint32(a.AuthorId),
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
	a, err := g.a.Delete(ctx, uint(request.Id)) // TODO: context things
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
		Id:        uint32(a.Id),
		CustomId:  a.CustomId,
		AuthorId:  uint32(a.AuthorId),
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   content,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}

	return &sar, nil
}
