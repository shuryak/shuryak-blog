package grpc

import (
	"auth/internal"
	"auth/internal/delivery/grpc/auth_grpc"
	"auth/internal/entity"
	"auth/pkg/jwt"
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
	jwtManager     *jwt.Manager
	sessionManager *internal.SessionManager
	auth_grpc.UnimplementedAuthServer
}

func NewAuthGRPCServer(s *grpc.Server, us internal.UserStore, jm *jwt.Manager, sm *internal.SessionManager) {
	as := &AuthGRPCServer{
		userStore:      us,
		jwtManager:     jm,
		sessionManager: sm,
	}
	auth_grpc.RegisterAuthServer(s, as)
	reflection.Register(s)
}

func (s AuthGRPCServer) Register(ctx context.Context, req *auth_grpc.RegisterRequest) (
	*auth_grpc.RegisterResponse,
	error,
) {
	user, err := entity.NewUser(req.GetUsername(), req.GetPassword(), req.GetRole())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "password error: %w", err)
	}

	// TODO: handle already exists
	createdUser, err := s.userStore.Create(ctx, *user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %w", err)
	}

	return &auth_grpc.RegisterResponse{
		Id:       uint32(createdUser.Id),
		Username: createdUser.Username,
		Role:     createdUser.Role,
	}, nil
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

func (s AuthGRPCServer) Validate(ctx context.Context, req *auth_grpc.ValidateRequest) (
	*auth_grpc.ValidateResponse,
	error,
) {
	claims, err := s.jwtManager.Decode(req.GetAccessToken())

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "access token is bad")
	}

	user, err := s.userStore.GetByUsername(ctx, claims.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &auth_grpc.ValidateResponse{
		UserId:   uint32(user.Id),
		Username: user.Username,
	}, nil
}
