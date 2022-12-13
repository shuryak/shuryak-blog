package routes

import (
	"api-gateway/internal/auth/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequestBody struct {
	Username string `form:"username" binding:"min=2,max=20,required"`
	Password string `form:"password" binding:"min=8,max=64,required"`
}

// Register
// @Summary     Method to register
// @Description User registration
// @Produce  	json
// @Param  	 	username query int true "Username"
// @Param  	 	password query int true "User password"
// @Success     200   	 {object} pb.RegisterResponse
// @Failure     400      {object} errors.Response
// @Failure     502      {object} errors.Response
// @Router      /auth/register [get]
func (r *Routes) Register(ctx *gin.Context) {
	var req RegisterRequestBody
	if err := ctx.ShouldBindQuery(&req); err != nil {
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
