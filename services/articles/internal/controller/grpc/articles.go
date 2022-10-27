package grpc

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/controller/grpc/articles_grpc"
	"github.com/shuryak/shuryak-blog/internal/entity"
	"github.com/shuryak/shuryak-blog/internal/usecase"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/structpb"
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
		l: l}
	articles_grpc.RegisterArticlesServer(server, ags)
	reflection.Register(server)
}

func (g articlesGRPCServer) Create(ctx context.Context, request *articles_grpc.CreateRequest) (*articles_grpc.CreateResponse, error) {
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

	cr := articles_grpc.CreateResponse{
		Id:        uint32(article.Id),
		CustomId:  article.CustomId,
		AuthorId:  uint32(article.AuthorId),
		Title:     article.Title,
		Thumbnail: article.Thumbnail,
		Content:   content,
	}

	return &cr, nil
}
