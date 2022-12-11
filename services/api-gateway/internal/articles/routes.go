package articles

import (
	"api-gateway/config"
	"api-gateway/internal/articles/routes"
	"api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.NewAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: NewServiceClient(cfg),
	}

	routesGroup := r.Group("/articles")
	routesGroup.Use(a.AuthRequired)
	routesGroup.POST("/create", svc.Create)

	return svc
}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.CreateArticle(ctx, svc.Client)
}
