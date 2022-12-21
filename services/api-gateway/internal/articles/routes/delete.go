package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/dto"
	"api-gateway/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Delete godoc
// @Summary     Deletes article by ID
// @Description Deletes article by ID
// @Produce  	json
// @Param  	 	id query int true "ID to delete"
// @Success     200   	 {object} dto.SingleArticleResponse
// @Failure     400      {object} errors.Response
// @Failure     502      {object} errors.Response
// @Router      /articles/delete [delete]
// @Security 	BearerAuth
func (r *Routes) Delete(ctx *gin.Context) {
	var req dto.ArticleIdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - Delete")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	article, err := r.c.Delete(ctx, &pb.ArticleIdRequest{
		Id: req.Id,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - Delete")
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
