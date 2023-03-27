package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	"github.com/shuryak/shuryak-blog/pkg/constants"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"go-micro.dev/v4/metadata"
	"net/http"
	"strings"
)

type TokenMiddleware struct {
	l logger.Interface
}

func NewTokenMiddleware(l logger.Interface) *TokenMiddleware {
	return &TokenMiddleware{l}
}

func (m *TokenMiddleware) ContextBearerString(ctx *gin.Context) {
	bearerString := ctx.Request.Header.Get("authorization")
	if bearerString == "" {
		errors.ErrorResponse(ctx, http.StatusUnauthorized, "no token")
		return
	}

	// TODO: validate bearer string

	newCtx := metadata.Set(ctx.Request.Context(), constants.AuthMetadataName, bearerString)
	ctx.Request = ctx.Request.WithContext(newCtx)

	ctx.Next()
}

func (m *TokenMiddleware) ContextToken(ctx *gin.Context) {
	bearerString := ctx.Request.Header.Get("authorization")
	if bearerString == "" {
		errors.ErrorResponse(ctx, http.StatusUnauthorized, "no token")
		return
	}

	bearerSplit := strings.Split(bearerString, "Bearer ")

	if len(bearerSplit) != 2 {
		errors.ErrorResponse(ctx, http.StatusUnauthorized, "invalid bearer payload")
		return
	}

	token := strings.TrimSpace(bearerSplit[1])

	ctx.Set(constants.AuthMetadataName, token)

	ctx.Next()
}
