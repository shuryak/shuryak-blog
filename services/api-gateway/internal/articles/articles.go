package articles

import (
	"api-gateway/config"
	"api-gateway/internal/articles/routes"
	"api-gateway/internal/auth"
	"api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, cfg *config.Config, l logger.Interface) {
	authClient, err := auth.NewServiceClient(cfg)
	if err != nil {
		l.Error(err, "internal - articles - RegisterRoutes")
	}

	a := auth.NewMiddleware(authClient, l)

	articlesClient, err := NewServiceClient(cfg)
	if err != nil {
		l.Error(err, "internal - articles - RegisterRoutes")
	}

	r := routes.NewRoutes(articlesClient, l)

	h1 := engine.Group("/articles").Use(a.AuthRequired)
	{
		h1.POST("/create", r.Create)
		h1.PUT("/update", r.Update)
		h1.DELETE("/delete", r.Delete)
	}

	h2 := engine.Group("/articles")
	{
		h2.GET("/getById", r.GetById)
		h2.GET("/getMany", r.GetMany)
	}
}
