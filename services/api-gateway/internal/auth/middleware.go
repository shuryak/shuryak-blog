package auth

import (
	"api-gateway/internal/auth/pb"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	svc *ServiceClient
}

func NewAuthMiddleware(svc *ServiceClient) AuthMiddleware {
	return AuthMiddleware{svc}
}

func (c *AuthMiddleware) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		AccessToken: token[1],
	})

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("username", res.Username)
	ctx.Set("user_id", res.UserId)

	ctx.Next()
}
