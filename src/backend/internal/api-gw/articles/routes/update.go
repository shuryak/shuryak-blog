package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/articles/dto"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	microErrors "go-micro.dev/v4/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
)

// Update godoc
// @Summary     Update an article
// @Description Update an article
// @Tags        Articles
// @Accept      json
// @Produce     json
// @Param       request body     dto.ArticleUpdateRequest true "article updated data"
// @Success     200     {object} dto.SingleArticleResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Failure     502     {object} errors.Response
// @Router      /articles/update [patch]
// @Security 	BearerAuth
func (r *Routes) Update(ctx *gin.Context) {
	var req dto.ArticleUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "articles - routes - Update: %v", err)
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	content, err := structpb.NewStruct(req.Content)
	if err != nil {
		r.l.Error(err, "articles - routes - Update: %v", err)
		errors.ErrorResponse(ctx, http.StatusInternalServerError, "some problems")
		return
	}

	a, err := r.c.Update(ctx.Request.Context(), &pb.UpdateRequest{
		CustomId:  req.CustomId,
		Title:     req.Title,
		Thumbnail: req.Thumbnail,
		Content:   content,
		IsDraft:   *req.IsDraft,
	})
	if err != nil {
		clientError := microErrors.FromError(err)

		switch clientError.Code {
		case http.StatusUnauthorized:
			errors.ErrorResponse(ctx, http.StatusUnauthorized, clientError.Detail)
		case http.StatusBadRequest:
			errors.ErrorResponse(ctx, http.StatusBadRequest, "bad request")
		case http.StatusBadGateway:
			errors.ErrorResponse(ctx, http.StatusBadGateway, "articles service error")
		default:
			return
		}
	}

	resp := dto.SingleArticleResponse{
		Id:        int(a.Id),
		CustomId:  a.CustomId,
		AuthorId:  int(a.AuthorId),
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   a.Content.AsMap(),
		IsDraft:   a.IsDraft,
		CreatedAt: a.CreatedAt.AsTime(),
		UpdatedAt: a.UpdatedAt.AsTime(),
	}

	ctx.JSON(http.StatusOK, &resp)
}
