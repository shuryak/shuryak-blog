package main

import (
	"api-gateway/config"
	"api-gateway/internal/articles"
	"api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, cfg)
	articles.RegisterRoutes(r, cfg, &authSvc)

	_ = r.Run("0.0.0.0:" + cfg.HTTP.Port)
}
