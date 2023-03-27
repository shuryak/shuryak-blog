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

// RefreshSession godoc
// @Summary     Refresh user session
// @Description Refresh user session
// @Tags        User
// @Accept      json
// @Produce     json
// @Success     200     {object} dto.AccessTokenResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Failure     502     {object} errors.Response
// @Router      /user/refresh_session [get]
// @Security 	BearerAuth
func (r *Routes) RefreshSession(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie(constants.RefreshTokenCookieName)
	if err != nil {
		errors.ErrorResponse(ctx, http.StatusBadRequest, "no refresh token")
		return
	}
	accessToken, ok := ctx.Get(constants.AuthMetadataName)
	if !ok {
		errors.ErrorResponse(ctx, http.StatusBadRequest, "no access token")
		return
	}

	tokenPair, err := r.c.RefreshSession(ctx.Request.Context(), &pb.RefreshSessionRequest{
		AccessToken:  accessToken.(string),
		RefreshToken: refreshToken,
	})
	if err != nil {
		errors.ErrorResponse(ctx, http.StatusBadGateway, "user service error")
		r.l.Error(err, "user - routes - RefreshSession")
		return
	}

	resp := dto.AccessTokenResponse{
		AccessToken: tokenPair.GetAccessToken(),
		ExpiresAt:   tokenPair.GetExpiresAt().AsTime(),
	}

	sessionExpIn := tokenPair.GetExpiresAt().AsTime().Sub(time.Now()).Seconds()

	ctx.SetCookie(constants.RefreshTokenCookieName, tokenPair.GetRefreshToken(), int(sessionExpIn), "/", "", true, true)
	ctx.JSON(http.StatusOK, &resp)
}
