package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetManyRequest struct {
	Offset uint32 `form:"offset" binding:"min=0,required" example:"42"`
	Count  uint32 `form:"count" binding:"min=1,required" example:"10"`
}

func (r *Routes) GetMany(ctx *gin.Context) {
	var req GetManyRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - GetMany")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := r.c.GetMany(context.Background(), &pb.GetManyRequest{
		Offset: req.Offset,
		Count:  req.Count,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - GetMany")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
