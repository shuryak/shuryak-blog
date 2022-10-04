package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"shuryak-blog/config"
	v1 "shuryak-blog/internal/controller/http/v1"
	"shuryak-blog/internal/usecase"
	"shuryak-blog/internal/usecase/repo"
	"shuryak-blog/pkg/httpserver"
	"shuryak-blog/pkg/postgres"
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
