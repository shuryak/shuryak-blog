package articles

import (
	"api-gateway/config"
	"api-gateway/internal/articles/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(cfg *config.Config) (pb.ArticlesClient, error) {
	cc, err := grpc.Dial(cfg.Services.ArticlesURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("internal - articles - NewServiceClient: %w", err)
	}

	return pb.NewArticlesClient(cc), nil
}
