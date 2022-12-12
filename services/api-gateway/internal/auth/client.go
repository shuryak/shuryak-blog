package auth

import (
	"api-gateway/config"
	"api-gateway/internal/auth/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(cfg *config.Config) (pb.AuthClient, error) {
	cc, err := grpc.Dial(cfg.Services.AuthURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("internal - auth - NewServiceClient: %w", err)
	}

	return pb.NewAuthClient(cc), nil
}
