package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/dto"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetManyRequest struct {
	// https://github.com/go-playground/validator/issues/290#issuecomment-316062118
	Offset *uint32 `form:"offset" binding:"min=0,required" example:"42"`
	Count  uint32  `form:"count" binding:"min=1,required" example:"10"`
}

// GetMany godoc
// @Summary     Gets collection of articles
// @Description Gets collection of articles
// @Produce  	json
// @Param   	offset query int true "offset to get"
// @Param   	count  query int true "count to get"
// @Success     200   	 {object} dto.ManyArticlesResponse
// @Failure     400      {object} errors.Response
// @Failure     502      {object} errors.Response
// @Router      /articles/getMany [get]
func (r *Routes) GetMany(ctx *gin.Context) {
	var req GetManyRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - GetMany")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	articles, err := r.c.GetMany(context.Background(), &pb.GetManyRequest{
		Offset: *req.Offset,
		Count:  req.Count,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - GetMany")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	res := dto.ManyArticlesResponse{Articles: make([]dto.SingleArticleResponse, len(articles.Articles))}
	for i, a := range articles.Articles {
		res.Articles[i] = dto.SingleArticleResponse{
			Id:        a.Id,
			CustomId:  a.CustomId,
			AuthorId:  a.AuthorId,
			Title:     a.Title,
			Thumbnail: a.Thumbnail,
			Content:   a.Content.AsMap(),
			CreatedAt: a.CreatedAt.AsTime(),
		}
	}

	ctx.JSON(http.StatusOK, &res)
}
