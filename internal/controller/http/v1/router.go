package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"shuryak-blog/internal/usecase"

	// Swagger docs
	_ "shuryak-blog/docs"
)

// NewRouter - .
// Swagger spec:
// @title Articles API
// @description Service for managing articles
// @version 1.0.0
// @host localhost:8080
// @BasePath /v1
func NewRouter(handler *gin.Engine, a usecase.Article) {
	// TODO: more options

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newArticlesRoutes(h, a)
	}
}
