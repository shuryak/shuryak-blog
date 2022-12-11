package routes

import (
	"api-gateway/internal/auth/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.AuthClient) {
	req := RegisterRequestBody{}

	if err := ctx.BindJSON(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Username: req.Username,
		Role:     "user",
		Password: req.Password,
	})

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
