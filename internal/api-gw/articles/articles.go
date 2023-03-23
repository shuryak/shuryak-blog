package articles

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/articles/routes"
	"github.com/shuryak/shuryak-blog/internal/api-gw/config"
	"github.com/shuryak/shuryak-blog/internal/api-gw/middleware"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	"go-micro.dev/v4/client"
)

func RegisterRoutes(engine *gin.Engine, c client.Client, cfg *config.Config, l logger.Interface) {
	a := middleware.NewTokenMiddleware(l)

	srv := pb.NewArticlesService("articles", c)

	r := routes.NewRoutes(srv, l)

	h := engine.Group("/articles").Use(a.TokenRequired)
	{
		h.POST("/create", r.Create)
	}
}
