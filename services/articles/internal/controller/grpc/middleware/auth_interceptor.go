package middleware

import (
	"context"
	"github.com/shuryak/shuryak-blog/pkg/jwt"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
)

type AuthMiddleware struct {
	v jwt.Validator
	l logger.Interface
}

func NewAuthMiddleware(v jwt.Validator, l logger.Interface) *AuthMiddleware {
	return &AuthMiddleware{v, l}
}

func (m *AuthMiddleware) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		m.l.Info("--> unary interceptor: ", info.FullMethod)

		headers, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
		}

		tokens := headers.Get("access_token")
		for len(tokens) < 1 {
			return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
		}

		accessToken := tokens[0]

		uc, err := m.v.Decode(accessToken)
		if err != nil {
			return nil, status.New(codes.Unauthenticated, "bad access token").Err()
		}

		newCtx := metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
			"user_id":  strconv.Itoa(int(uc.UserId)),
			"username": uc.Username,
			"role":     uc.Role,
		}))

		return handler(newCtx, req)
	}
}

func (m *AuthMiddleware) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		m.l.Info("--> stream interceptor: ", info.FullMethod)

		wrapped := WrapServerStream(stream)

		headers, ok := metadata.FromIncomingContext(wrapped.Context())
		if !ok {
			return status.New(codes.Unauthenticated, "no auth provided").Err()
		}

		tokens := headers.Get("access_token")
		for len(tokens) < 1 {
			return status.New(codes.Unauthenticated, "no auth provided").Err()
		}

		accessToken := tokens[0]

		newCtx := metadata.AppendToOutgoingContext(wrapped.Context(), "access_token", accessToken)
		wrapped.WrappedContext = newCtx

		return handler(srv, wrapped)
	}
}
