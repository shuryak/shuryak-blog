package routes

import (
	"api-gateway/internal/auth/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequestBody struct {
	Username string `json:"username" binding:"min=2,max=20,required"`
	Password string `json:"password" binding:"min=8,max=64,required"`
}

func (r *Routes) Register(ctx *gin.Context) {
	var req RegisterRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "auth - routes - Register")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := r.c.Register(context.Background(), &pb.RegisterRequest{
		Username: req.Username,
		Role:     "user",
		Password: req.Password,
	})
	if err != nil {
		r.l.Error(err, "auth - routes - Register")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
