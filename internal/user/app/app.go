package app

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/shuryak/shuryak-blog/internal/user/config"
	"github.com/shuryak/shuryak-blog/internal/user/handlers"
	"github.com/shuryak/shuryak-blog/internal/user/usecase"
	"github.com/shuryak/shuryak-blog/internal/user/usecase/repo"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
	pb "github.com/shuryak/shuryak-blog/proto/user"
	"go-micro.dev/v4"
	"time"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	// Migrations
	m, err := migrate.New("file://../migrations", cfg.PG.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - migrate.New: %w", err))
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		l.Fatal(fmt.Errorf("app - Run - m.Up: %w", err))
	}

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	users := usecase.NewUsersUseCase(repo.NewUsersRepo(pg))
	userSessions := usecase.NewUserSessionsUseCase(repo.NewUserSessionsRepo(pg), 64, 7*24*time.Hour) // TODO: config?
	jwt := handlers.NewJWTManager("jwtsecret", 30*time.Minute)                                       // TODO: jwt is not handlers                                          // TODO: store secret

	h := handlers.NewUsersHandler(users, userSessions, *jwt, cfg.Service.Name, l)

	srv := micro.NewService(
		micro.Name(cfg.Service.Name),
		micro.Version(cfg.Service.Version),
	)

	srv.Init()

	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(), h); err != nil {
		l.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		l.Fatal(err)
	}
}
