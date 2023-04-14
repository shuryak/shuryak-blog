package handlers

import (
	"context"
	"github.com/shuryak/shuryak-blog/internal/articles/entity"
	"github.com/shuryak/shuryak-blog/internal/articles/usecase"
	"github.com/shuryak/shuryak-blog/pkg/constants"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	"go-micro.dev/v4/errors"
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

func (h *ArticlesHandler) Create(ctx context.Context, req *pb.CreateRequest, resp *pb.SingleArticleResponse) error {
	meta, _ := metadata.FromContext(ctx)
	userId, err := strconv.ParseUint(meta[constants.UserIdMetadataName], 10, 32)
	if err != nil {
		h.l.Error("articles - Create - metadata: %v", err)
		return errors.Unauthorized("metadata", "can't parse %s from metadata", constants.UserIdMetadataName)
	}

	a := entity.Article{
		CustomId:  req.GetCustomId(),
		AuthorId:  uint32(userId),
		Title:     req.GetTitle(),
		Thumbnail: req.GetThumbnail(),
		Content:   req.GetContent().AsMap(),
		IsDraft:   req.GetIsDraft(),
	}

	storedArticle, err := h.uc.Create(ctx, a)
	if err != nil {
		h.l.Error("articles - Create - h.uc.Create: %v", err)
		return errors.InternalServerError("articles", "article create error") // TODO: inner errors
	}

	content, err := structpb.NewStruct(storedArticle.Content)
	if err != nil {
		h.l.Error("articles - Create - structpb.NewStruct: %v", err)
		return errors.BadRequest("pbstruct", "invalid content")
	}

	resp.Id = storedArticle.Id
	resp.CustomId = storedArticle.CustomId
	resp.AuthorId = storedArticle.AuthorId
	resp.Title = storedArticle.Title
	resp.Thumbnail = storedArticle.Thumbnail
	resp.Content = content
	resp.IsDraft = storedArticle.IsDraft
	resp.CreatedAt = timestamppb.New(storedArticle.CreatedAt)
	resp.UpdatedAt = timestamppb.New(storedArticle.UpdatedAt)

	return nil
}

func (h *ArticlesHandler) GetByCustomId(ctx context.Context, req *pb.ArticleCustomIdRequest, resp *pb.SingleArticleResponse) error {
	a, err := h.uc.GetByCustomId(ctx, req.GetCustomId())
	if err != nil {
		h.l.Error("articles - GetByCustomId - h.uc.Create: %v", err)
		return errors.BadRequest("no", "there is no article on such id") // TODO: inner errors
	}

	content, err := structpb.NewStruct(a.Content)
	if err != nil {
		h.l.Error("articles - Create - structpb.NewStruct: %v", err)
		return errors.BadRequest("pbstruct", "invalid content")
	}

	resp.Id = a.Id
	resp.CustomId = a.CustomId
	resp.AuthorId = a.AuthorId
	resp.Title = a.Title
	resp.Thumbnail = a.Thumbnail
	resp.Content = content
	resp.IsDraft = a.IsDraft
	resp.CreatedAt = timestamppb.New(a.CreatedAt)
	resp.UpdatedAt = timestamppb.New(a.UpdatedAt)

	return nil
}

func (h *ArticlesHandler) GetMany(ctx context.Context, req *pb.GetManyRequest, resp *pb.ShortArticlesResponse) error {
	articles, err := h.uc.GetMany(ctx, req.GetOffset(), req.GetCount(), req.GetIsDrafts())
	if err != nil {
		h.l.Error("articles - GetByCustomId - h.uc.Create: %v", err)
		return errors.BadRequest("no", "there are no articles in the specified range") // TODO: inner errors
	}

	for _, a := range articles {
		resp.Articles = append(resp.Articles, &pb.ShortArticle{
			Id:           a.Id,
			CustomId:     a.CustomId,
			AuthorId:     a.AuthorId,
			Title:        a.Title,
			Thumbnail:    a.Thumbnail,
			ShortContent: a.ShortContent,
			IsDraft:      a.IsDraft,
			CreatedAt:    timestamppb.New(a.CreatedAt),
			UpdatedAt:    timestamppb.New(a.UpdatedAt),
		})
	}

	return nil
}

func (h *ArticlesHandler) Update(ctx context.Context, req *pb.UpdateRequest, resp *pb.SingleArticleResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *ArticlesHandler) Delete(ctx context.Context, req *pb.ArticleCustomIdRequest, resp *pb.SingleArticleResponse) error {
	//TODO implement me
	panic("implement me")
}
