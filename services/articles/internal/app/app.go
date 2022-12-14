package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shuryak/shuryak-blog/config"
	grpccontroller "github.com/shuryak/shuryak-blog/internal/controller/grpc"
	"github.com/shuryak/shuryak-blog/internal/controller/grpc/middleware"
	v1 "github.com/shuryak/shuryak-blog/internal/controller/http/v1"
	"github.com/shuryak/shuryak-blog/internal/usecase"
	"github.com/shuryak/shuryak-blog/internal/usecase/repo"
	"github.com/shuryak/shuryak-blog/pkg/httpserver"
	"github.com/shuryak/shuryak-blog/pkg/jwt"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
)

func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
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

	// gRPC Server
	go func() {
		list, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
		if err != nil {
			l.Fatal(fmt.Errorf("app - Run - net.Listen: %w", err))
			return
		}

		v, err := jwt.NewValidator(cfg.PublicKeyPEMPath)
		if err != nil {
			l.Fatal(fmt.Errorf("app - Run - jwt.NewValidator: %w", err))
			return
		}

		authMiddleware := middleware.NewAuthMiddleware(*v, l)

		s := grpc.NewServer(
			grpc.UnaryInterceptor(authMiddleware.UnaryServerInterceptor()),
			grpc.StreamInterceptor(authMiddleware.StreamServerInterceptor()),
		) // TODO: to pkg
		grpccontroller.NewArticlesGrpcServer(s, articlesUseCase, l)
		err = s.Serve(list)
		if err != nil {
			l.Fatal(fmt.Errorf("app - Run - s.Server: %w", err))
		}
	}()

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, articlesUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
