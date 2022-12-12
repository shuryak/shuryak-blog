package auth

import (
	"api-gateway/internal/auth/pb"
	"api-gateway/internal/errors"
	"api-gateway/pkg/logger"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Middleware struct {
	c pb.AuthClient
	l logger.Interface
}

func NewMiddleware(c pb.AuthClient, l logger.Interface) Middleware {
	return Middleware{c, l}
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

	res, err := m.c.Validate(context.Background(), &pb.ValidateRequest{
		AccessToken: token[1],
	})
	if err != nil {
		m.l.Error(err, "auth - middleware - AuthRequired")
		errors.ErrorResponse(ctx, http.StatusUnauthorized, "invalid token")
		return
	}

	ctx.Set("username", res.Username)
	ctx.Set("user_id", res.UserId)

	ctx.Next()
}
