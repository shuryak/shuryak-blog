package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/dto"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
)

type UpdateRequest struct {
	Id        uint32                 `json:"id" binding:"min=1,required" example:"1000"`
	CustomId  string                 `json:"custom_id" binding:"min=3,max=20,required" example:"article-url"`
	Title     string                 `json:"title" binding:"min=5,max=150,required" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" binding:"url,required" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content" binding:"required"`
}

// Update godoc
// @Summary     Updates article by ID
// @Description Updates article by ID
// @Accept      json
// @Produce     json
// @Param       request body     UpdateRequest true "article to update"
// @Success     200     {object} dto.SingleArticleResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Failure     502     {object} errors.Response
// @Router      /articles/update [put]
// @Security 	BearerAuth
func (r *Routes) Update(ctx *gin.Context) {
	var req UpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "articles - routes - Update")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	content, err := structpb.NewStruct(req.Content)
	if err != nil {
		r.l.Error(err, "articles - routes - Update")
		errors.ErrorResponse(ctx, http.StatusInternalServerError, "some problems")
	}

	authorId, _ := ctx.Get("user_id")

	article, err := r.c.Update(context.Background(), &pb.UpdateRequest{
		Id:        req.Id,
		CustomId:  req.CustomId,
		AuthorId:  authorId.(uint32),
		Title:     req.Title,
		Thumbnail: req.Thumbnail,
		Content:   content,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - Update")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	res := dto.SingleArticleResponse{
		Id:        article.Id,
		CustomId:  article.CustomId,
		AuthorId:  article.AuthorId,
		Title:     article.Title,
		Thumbnail: article.Thumbnail,
		Content:   article.Content.AsMap(),
		CreatedAt: article.CreatedAt.AsTime(),
	}

	ctx.JSON(http.StatusOK, &res)
}
