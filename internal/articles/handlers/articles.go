package handlers

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/articles/entity"
	"github.com/shuryak/shuryak-blog/internal/articles/usecase"
	"github.com/shuryak/shuryak-blog/pkg/constants"
	"github.com/shuryak/shuryak-blog/pkg/errors"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type ArticlesHandler struct {
	uc usecase.Articles
	l  logger.Interface
}

func NewArticlesHandler(uc usecase.ArticlesUseCase, l logger.Interface) *ArticlesHandler {
	return &ArticlesHandler{uc, l}
}

// Check for implementation
var _ pb.ArticlesHandler = (*ArticlesHandler)(nil)

var GlobalErrors errors.ServerError

func (h *ArticlesHandler) Create(ctx context.Context, req *pb.CreateRequest, resp *pb.SingleArticleResponse) error {
	meta, _ := metadata.FromContext(ctx)
	userId, err := strconv.ParseUint(meta[constants.UserIdMetadataName], 10, 32)
	if err != nil {
		return fmt.Errorf("can't parse user_id from metadata: %w", err) // TODO: correct error
	}

	a := entity.Article{
		CustomId:  req.GetCustomId(),
		AuthorId:  uint32(userId),
		Title:     req.GetTitle(),
		Thumbnail: req.GetThumbnail(),
		Content:   req.GetContent().AsMap(),
	}

	storedArticle, err := h.uc.Create(ctx, a)
	if err != nil {
		return fmt.Errorf(GlobalErrors.AuthNoToken()) // TODO: correct error
	}

	content, err := structpb.NewStruct(storedArticle.Content)
	if err != nil {
		return fmt.Errorf(GlobalErrors.AuthNoToken()) // TODO: correct error
	}

	resp.Id = storedArticle.Id
	resp.CustomId = storedArticle.CustomId
	resp.AuthorId = storedArticle.AuthorId
	resp.Title = storedArticle.Title
	resp.Thumbnail = storedArticle.Thumbnail
	resp.Content = content
	resp.CreatedAt = timestamppb.New(storedArticle.CreatedAt)
	resp.UpdatedAt = timestamppb.New(storedArticle.UpdatedAt)

	return nil
}

func (h *ArticlesHandler) GetById(ctx context.Context, req *pb.ArticleCustomIdRequest, resp *pb.SingleArticleResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *ArticlesHandler) GetShortMany(ctx context.Context, req *pb.GetManyRequest, resp *pb.ShortArticlesResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *ArticlesHandler) Update(ctx context.Context, req *pb.UpdateRequest, resp *pb.SingleArticleResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *ArticlesHandler) Delete(ctx context.Context, req *pb.ArticleCustomIdRequest, resp *pb.SingleArticleResponse) error {
	//TODO implement me
	panic("implement me")
}
