package auth

import (
	"api-gateway/config"
	"api-gateway/internal/auth/routes"
	"api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, cfg *config.Config, l logger.Interface) {
	authClient, err := NewServiceClient(cfg)
	if err != nil {
		l.Error(err, "internal - auth - RegisterRoutes")
	}

	r := routes.NewRoutes(authClient, l)

	h := engine.Group("/auth")
	{
		h.GET("/register", r.Register)
		h.GET("/login", r.Login)
	}
}
