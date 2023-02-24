package app

import (
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/articles/config"
	"github.com/shuryak/shuryak-blog/internal/articles/handlers"
	"github.com/shuryak/shuryak-blog/internal/articles/middleware"
	"github.com/shuryak/shuryak-blog/internal/articles/usecase"
	"github.com/shuryak/shuryak-blog/internal/articles/usecase/repo"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	"go-micro.dev/v4"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	uc := usecase.NewArticlesUseCase(repo.NewArticlesRepo(pg))
	h := handlers.NewArticlesHandler(*uc, l)
	auth := middleware.AuthWrapper

	srv := micro.NewService(
		micro.Name(cfg.Service.Name),
		micro.Version(cfg.Service.Version),
		micro.WrapHandler(auth),
	)

	srv.Init()

	// Register handlers
	if err := pb.RegisterArticlesHandler(srv.Server(), h); err != nil {
		l.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		l.Fatal(err)
	}
}
