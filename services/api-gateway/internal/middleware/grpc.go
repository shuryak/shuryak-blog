package middleware

import (
	"api-gateway/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GRPCAuthMiddleware struct {
	l logger.Interface
}

func NewGRPCAuthMiddleware(l logger.Interface) *GRPCAuthMiddleware {
	return &GRPCAuthMiddleware{l}
}

func (i *GRPCAuthMiddleware) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		i.l.Info("--> unary interceptor: ", method)

		token, err := contextGetToken(ctx)
		if err != nil {
			return fmt.Errorf("token not set in context: %w", err)
		}

		newCtx := metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
			"access_token": token,
		}))

		// TODO: I don't know if there is logic in creating a new context

		return invoker(newCtx, method, req, reply, cc, opts...)
	}
}

func (i *GRPCAuthMiddleware) StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		i.l.Info("--> stream interceptor: ", method)

		token, err := contextGetToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("token not set in context: %w", err)
		}

		newCtx := metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
			"access_token": token,
		}))

		return streamer(newCtx, desc, cc, method, opts...)
	}
}

func contextGetToken(ctx context.Context) (string, error) {
	token := ctx.Value("access_token")
	if token == nil {
		return "", fmt.Errorf("no token in context")
	}

	return token.(string), nil
}
