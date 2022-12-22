package articles

import (
	"api-gateway/config"
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/middleware"
	"api-gateway/pkg/logger"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(cfg *config.Config, l logger.Interface) (pb.ArticlesClient, error) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())

	m := middleware.NewGRPCAuthMiddleware(l)

	cc, err := grpc.Dial(
		cfg.Services.ArticlesURL,
		transportOption,
		grpc.WithUnaryInterceptor(m.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(m.StreamClientInterceptor()),
	)

	if err != nil {
		return nil, fmt.Errorf("internal - articles - NewServiceClient: %w", err)
	}

	return pb.NewArticlesClient(cc), nil
}
