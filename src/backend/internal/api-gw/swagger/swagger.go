package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/config"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs
	_ "github.com/shuryak/shuryak-blog/internal/api-gw/docs"
)

// RegisterSwagger godoc
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func RegisterSwagger(engine *gin.Engine, cfg *config.Config, l logger.Interface) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	engine.GET("/swagger/*any", swaggerHandler)
}
