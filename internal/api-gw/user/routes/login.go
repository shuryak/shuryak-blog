package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	"github.com/shuryak/shuryak-blog/internal/api-gw/user/dto"
	"github.com/shuryak/shuryak-blog/pkg/constants"
	pb "github.com/shuryak/shuryak-blog/proto/user"
	"net/http"
	"time"
)

// Login godoc
// @Summary     Login to user account
// @Description Login to user account
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       request body     dto.LoginRequest true "user to register"
// @Success     200     {object} dto.AccessTokenResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Failure     502     {object} errors.Response
// @Router      /user/login [post]
func (r *Routes) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "user - routes - Login")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	tokenPair, err := r.c.Login(ctx.Request.Context(), &pb.LoginRequest{
		Username:      req.Username,
		PlainPassword: req.Password,
	})
	if err != nil {
		errors.ErrorResponse(ctx, http.StatusBadGateway, "user service error")
		r.l.Error(err, "user - routes - Login")
		return
	}

	resp := dto.AccessTokenResponse{
		AccessToken: tokenPair.GetAccessToken(),
		ExpiresAt:   tokenPair.GetExpiresAt().AsTime(),
	}

	sessionExpIn := tokenPair.GetExpiresAt().AsTime().Sub(time.Now()).Seconds()

	ctx.SetCookie(constants.RefreshTokenCookieName, tokenPair.GetRefreshToken(), int(sessionExpIn), "/", "", false, true)
	ctx.JSON(http.StatusOK, &resp)
}
