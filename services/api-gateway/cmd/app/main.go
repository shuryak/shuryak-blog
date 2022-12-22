package main

import (
	"api-gateway/config"
	"api-gateway/internal/articles"
	"api-gateway/internal/auth"
	"api-gateway/internal/swagger"
	"api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	l := logger.New(cfg.Log.Level)

	// https://blog.depa.do/post/gin-validation-errors-handling
	// https://github.com/go-playground/validator/blob/21c910fc6d9c3556c28252b04beb17de0c2d40ec/validator_instance.go#L137
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	engine := gin.New()

	swagger.RegisterSwagger(engine, cfg, l)
	articles.RegisterRoutes(engine, cfg, l)
	auth.RegisterRoutes(engine, cfg, l)

	_ = engine.Run("0.0.0.0:" + cfg.HTTP.Port)
}
