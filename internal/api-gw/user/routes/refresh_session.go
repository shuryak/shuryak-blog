package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	"github.com/shuryak/shuryak-blog/internal/api-gw/user/dto"
	pb "github.com/shuryak/shuryak-blog/proto/user"
	"net/http"
)

// RefreshSession godoc
// @Summary     Refresh user session
// @Description Refresh user session
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       request body     dto.RefreshSessionRequest true "user to register"
// @Success     200     {object} dto.TokenPairResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Failure     502     {object} errors.Response
// @Router      /user/refresh_session [post]
func (r *Routes) RefreshSession(ctx *gin.Context) {
	var req dto.RefreshSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "user - routes - RefreshSession")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	tokenPair, err := r.c.RefreshSession(ctx.Request.Context(), &pb.RefreshSessionRequest{
		Username:     req.Username,
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		errors.ErrorResponse(ctx, http.StatusBadGateway, "user service error")
		r.l.Error(err, "user - routes - RefreshSession")
		return
	}

	resp := dto.TokenPairResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}

	ctx.JSON(http.StatusOK, &resp)
}
