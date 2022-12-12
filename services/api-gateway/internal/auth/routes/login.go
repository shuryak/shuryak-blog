package routes

import (
	"api-gateway/internal/auth/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequestBody struct {
	Username string `form:"username" binding:"min=2,max=20,required"`
	Password string `form:"password" binding:"min=8,max=64,required"`
}

func (r *Routes) Login(ctx *gin.Context) {
	var req LoginRequestBody
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "auth - routes - Login")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := r.c.Login(context.Background(), &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		r.l.Error(err, "auth - routes - Login")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
