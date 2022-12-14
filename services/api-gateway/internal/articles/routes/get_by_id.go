package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GetByIdRequest struct {
	Id uint32 `json:"id" binding:"min=1,required" example:"42"`
}

type GetByIdResponse struct {
	Id        uint32                 `json:"id" example:"1000"`
	CustomId  string                 `json:"custom_id" example:"article-url"`
	AuthorId  uint32                 `json:"author_id" example:"42"`
	Title     string                 `json:"title" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at" example:"2022-10-07T14:26:06.510465Z"`
}

// GetById godoc
// @Summary     Gets article by ID
// @Description Gets article by ID
// @Produce  	json
// @Param  	 	id query int true "ID to get"
// @Success     200   	 {object} GetByIdResponse
// @Failure     400      {object} errors.Response
// @Failure     502      {object} errors.Response
// @Router      /articles/getById [get]
func (r *Routes) GetById(ctx *gin.Context) {
	var req GetByIdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - GetById")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	article, err := r.c.GetById(context.Background(), &pb.ArticleIdRequest{
		Id: req.Id,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - GetById")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	res := CreateResponse{
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
