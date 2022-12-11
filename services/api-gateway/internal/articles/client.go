package articles

import (
	"api-gateway/config"
	"api-gateway/internal/articles/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ArticlesClient
}

func NewServiceClient(cfg *config.Config) pb.ArticlesClient {
	cc, err := grpc.Dial(cfg.Services.ArticlesURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		_ = fmt.Errorf("internal - articles - new")
	}

	return pb.NewArticlesClient(cc)
}
