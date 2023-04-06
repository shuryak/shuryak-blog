package app

import (
	"context"
	"fmt"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/shuryak/shuryak-blog/internal/articles/config"
	"github.com/shuryak/shuryak-blog/internal/articles/handlers"
	"github.com/shuryak/shuryak-blog/internal/articles/middleware"
	"github.com/shuryak/shuryak-blog/internal/articles/usecase"
	"github.com/shuryak/shuryak-blog/internal/articles/usecase/repo"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
	"github.com/shuryak/shuryak-blog/pkg/tracing"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"strings"
	"time"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	// Migrations
	m, err := migrate.New("file://migrations", cfg.PG.URL)
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

	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)

	uc := usecase.NewArticlesUseCase(repo.NewArticlesRepo(pg))
	h := handlers.NewArticlesHandler(*uc, l)
	auth := middleware.NewAuthWrapper(srv.Client(), l)

	opts := []micro.Option{
		micro.Name(cfg.Service.Name),
		micro.Version(cfg.Service.Version),
		micro.WrapHandler(auth.Use),
	}

	// Jaeger
	tp, err := tracing.NewTracerProvider(cfg.Service.Name, cfg.Service.Version, srv.Server().Options().Id, cfg.Jaeger.URL)
	if err != nil {
		l.Fatal(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	traceOpts := []opentelemetry.Option{
		opentelemetry.WithHandleFilter(func(ctx context.Context, r server.Request) bool {
			if e := r.Endpoint(); strings.HasPrefix(e, "Health.") {
				return true
			}
			return false
		}),
	}
	opts = append(opts, micro.WrapHandler(opentelemetry.NewHandlerWrapper(traceOpts...)))
	opts = append(opts, micro.WrapClient(opentelemetry.NewClientWrapper(traceOpts...)))

	srv.Init(opts...)

	// Register handlers
	if err := pb.RegisterArticlesHandler(srv.Server(), h); err != nil {
		l.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handlers.HealthHandler)); err != nil {
		l.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		l.Fatal(err)
	}
}
