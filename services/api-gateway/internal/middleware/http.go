package middleware

import (
	"api-gateway/internal/errors"
	"api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Middleware struct {
	l logger.Interface
}

func NewMiddleware(l logger.Interface) Middleware {
	return Middleware{l}
}

func (m *Middleware) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		errors.ErrorResponse(ctx, http.StatusUnauthorized, "no parameters for authorization")
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		errors.ErrorResponse(ctx, http.StatusUnauthorized, "invalid parameters for authorization")
		return
	}

	ctx.Set("access_token", token[1])

	ctx.Next()
}
