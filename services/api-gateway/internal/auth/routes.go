package auth

import (
	"api-gateway/config"
	"api-gateway/internal/auth/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: NewServiceClient(cfg),
	}

	routesGroup := r.Group("/auth")
	routesGroup.POST("/register", svc.Register)
	routesGroup.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
