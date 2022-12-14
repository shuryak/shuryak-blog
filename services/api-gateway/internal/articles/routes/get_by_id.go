package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/dto"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetById godoc
// @Summary     Gets article by ID
// @Description Gets article by ID
// @Produce  	json
// @Param  	 	id query int true "ID to get"
// @Success     200   	 {object} dto.SingleArticleResponse
// @Failure     400      {object} errors.Response
// @Failure     502      {object} errors.Response
// @Router      /articles/getById [get]
func (r *Routes) GetById(ctx *gin.Context) {
	var req dto.ArticleIdRequest
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
