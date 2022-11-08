package client

import (
	"context"
	"github.com/shuryak/shuryak-blog/internal/controller/grpc/auth_grpc"
	"google.golang.org/grpc"
	"time"
)

type AuthClient struct {
	service  auth_grpc.AuthClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := auth_grpc.NewAuthClient(cc)
	return &AuthClient{service, username, password}
}

type TokenPair struct {
	accessToken  string
	refreshToken string
}

func (client *AuthClient) Login() (*TokenPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &auth_grpc.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		accessToken:  res.GetAccessToken(),
		refreshToken: res.GetRefreshToken(),
	}, nil
}
