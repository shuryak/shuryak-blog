package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"
	"github.com/go-playground/validator/v10"
	"github.com/shuryak/shuryak-blog/internal/api-gw/articles"
	"github.com/shuryak/shuryak-blog/internal/api-gw/config"
	"github.com/shuryak/shuryak-blog/internal/api-gw/swagger"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"github.com/shuryak/shuryak-blog/pkg/tracing"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"reflect"
	"strings"
	"time"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

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

	srv := micro.NewService(
		micro.Client(grpcc.NewClient()),
		micro.Server(grpcs.NewServer()),
	)
	opts := []micro.Option{
		micro.Name(cfg.Service.Name),
		micro.Version(cfg.Service.Version),
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

	engine := gin.New()

	swagger.RegisterSwagger(engine, cfg, l)
	articles.RegisterRoutes(engine, srv.Client(), cfg, l)

	_ = engine.Run("0.0.0.0" + cfg.HTTP.Port)
}
