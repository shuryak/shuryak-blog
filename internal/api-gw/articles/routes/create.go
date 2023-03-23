package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/articles/dto"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
)

// Create godoc
// @Summary     Creates an article
// @Description Creates an article
// @Accept      json
// @Produce     json
// @Param       request body     dto.ArticleCreateRequest true "article to create"
// @Success     200     {object} dto.SingleArticleResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Router      /articles/create [post]
// @Security 	BearerAuth
func (r *Routes) Create(ctx *gin.Context) {
	var req dto.ArticleCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "articles - routes - Create")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	content, err := structpb.NewStruct(req.Content)
	if err != nil {
		r.l.Error(err, "articles - routes - Create")
		errors.ErrorResponse(ctx, http.StatusInternalServerError, "some problems")
		return
	}

	a, err := r.c.Create(ctx.Request.Context(), &pb.CreateRequest{
		CustomId:  req.CustomId,
		Title:     req.Title,
		Thumbnail: req.Thumbnail,
		Content:   content,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - Create")
		return
	}

	res := dto.SingleArticleResponse{
		Id:        int(a.Id),
		CustomId:  a.CustomId,
		AuthorId:  int(a.AuthorId),
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   a.Content.AsMap(),
		CreatedAt: a.CreatedAt.AsTime(),
	}

	ctx.JSON(http.StatusCreated, &res)
}
