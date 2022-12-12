package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetByIdRequest struct {
	Id uint32 `json:"id" binding:"min=1,required" example:"42"`
}

func (r *Routes) GetById(ctx *gin.Context) {
	var req GetByIdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - GetById")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := r.c.GetById(context.Background(), &pb.ArticleIdRequest{
		Id: req.Id,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - GetById")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
