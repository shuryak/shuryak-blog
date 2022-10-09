package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"os"
	"os/signal"
	"reflect"
	"shuryak-blog/config"
	v1 "shuryak-blog/internal/controller/http/v1"
	"shuryak-blog/internal/usecase"
	"shuryak-blog/internal/usecase/repo"
	"shuryak-blog/pkg/httpserver"
	"shuryak-blog/pkg/postgres"
	"strings"
	"syscall"
)

func Run(cfg *config.Config) {
	// TODO: logger init

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		// TODO: log postgres err
	}
	defer pg.Close()

	// Use case
	articlesUseCase := usecase.New(repo.New(pg))

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

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, articlesUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("App is running: " + s.String())
	// TODO: LOG: app is running...
	case err = <-httpServer.Notify():
		// TODO: LOG ERROR
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		// TODO: LOG ?
	}
}
