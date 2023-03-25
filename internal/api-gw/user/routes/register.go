package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	"github.com/shuryak/shuryak-blog/internal/api-gw/user/dto"
	pb "github.com/shuryak/shuryak-blog/proto/user"
	"net/http"
)

// Register godoc
// @Summary     Register a user
// @Description Register a user
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       request body     dto.RegisterUserRequest true "user to register"
// @Success     201     {object} dto.SingleUserResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Failure     502     {object} errors.Response
// @Router      /user/register [post]
func (r *Routes) Register(ctx *gin.Context) {
	var req dto.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "user - routes - Register")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := r.c.Register(ctx.Request.Context(), &pb.RegisterRequest{
		Username:      req.Username,
		Role:          "user",
		PlainPassword: req.Password,
	})
	if err != nil {
		errors.ErrorResponse(ctx, http.StatusBadGateway, "user service error")
		r.l.Error(err, "user - routes - Register")
		return
	}

	resp := dto.SingleUserResponse{
		Id:       int(user.Id),
		Username: user.Username,
		Role:     user.Role,
	}

	ctx.JSON(http.StatusCreated, &resp)
}
