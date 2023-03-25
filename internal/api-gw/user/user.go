package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/config"
	"github.com/shuryak/shuryak-blog/internal/api-gw/user/routes"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/user"
	"go-micro.dev/v4/client"
)

func RegisterRoutes(engine *gin.Engine, c client.Client, cfg *config.Config, l logger.Interface) {
	srv := pb.NewUserService("user", c)

	r := routes.NewRoutes(srv, l)

	h := engine.Group("/user")
	{
		h.POST("/register", r.Register)
		h.POST("/login", r.Login)
		h.POST("/refresh_session", r.RefreshSession)
	}
}
