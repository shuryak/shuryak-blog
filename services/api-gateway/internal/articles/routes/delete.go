package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteRequest struct {
	Id uint32 `form:"id" binding:"min=1,required" example:"article-url"`
}

func (r *Routes) Delete(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - Delete")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := r.c.Delete(context.Background(), &pb.ArticleIdRequest{
		Id: req.Id,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - Delete")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
