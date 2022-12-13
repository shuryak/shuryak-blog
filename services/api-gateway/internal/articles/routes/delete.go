package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DeleteRequest struct {
	Id uint32 `form:"id" binding:"min=1,required" example:"article-url"`
}

type DeleteResponse struct {
	Id        uint32                 `json:"id" example:"1000"`
	CustomId  string                 `json:"custom_id" example:"article-url"`
	AuthorId  uint32                 `json:"author_id" example:"42"`
	Title     string                 `json:"title" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at" example:"2022-10-07T14:26:06.510465Z"`
}

// Delete
// @Summary     Deletes article by ID
// @Description Deletes article by ID
// @Produce  	json
// @Param  	 	id query int true "ID to delete"
// @Success     200   	 {object} DeleteResponse
// @Failure     400      {object} errors.Response
// @Failure     502      {object} errors.Response
// @Router      /articles/delete [delete]
func (r *Routes) Delete(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - Delete")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	article, err := r.c.Delete(context.Background(), &pb.ArticleIdRequest{
		Id: req.Id,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - Delete")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	res := DeleteResponse{
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
