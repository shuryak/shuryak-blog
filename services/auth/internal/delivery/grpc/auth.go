package grpc

import (
	"auth/internal"
	"auth/internal/delivery/grpc/auth_grpc"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Check for implementation
var _ auth_grpc.AuthServer = (*AuthGRPCServer)(nil)

type AuthGRPCServer struct {
	userStore      internal.UserStore
	jwtManager     *internal.JWTManager
	sessionManager *internal.SessionManager
	auth_grpc.UnimplementedAuthServer
}

func NewAuthServer(s *grpc.Server, us internal.UserStore, jm *internal.JWTManager, sm *internal.SessionManager) {
	as := &AuthGRPCServer{
		userStore:      us,
		jwtManager:     jm,
		sessionManager: sm,
	}
	auth_grpc.RegisterAuthServer(s, as)
	reflection.Register(s)
}

func (s AuthGRPCServer) Login(ctx context.Context, req *auth_grpc.LoginRequest) (*auth_grpc.LoginResponse, error) {
	user, err := s.userStore.GetByUsername(ctx, req.GetUsername()) // TODO: ctx?
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %w", err)
	}

	if user == nil || !user.IsPasswordCorrect(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username or password")
	}

	session, err := s.sessionManager.Start(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate refresh token")
	}

	accessToken, err := s.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &auth_grpc.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: session.RefreshToken,
	}, nil
}
