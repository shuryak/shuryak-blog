package auth

import (
	"api-gateway/config"
	"api-gateway/internal/auth/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthClient
}

func NewServiceClient(cfg *config.Config) pb.AuthClient {
	cc, err := grpc.Dial(cfg.Services.AuthURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		_ = fmt.Errorf("internal - auth - new")
	}

	return pb.NewAuthClient(cc)
}
