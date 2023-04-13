package handlers

import (
	"context"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HealthHandler struct{}

// Check for implementation
var _ pb.HealthHandler = (*HealthHandler)(nil)

func (h HealthHandler) Check(ctx context.Context, req *pb.HealthCheckRequest, resp *pb.HealthCheckResponse) error {
	resp.Status = pb.HealthCheckResponse_SERVING
	return nil
}

func (h HealthHandler) Watch(ctx context.Context, request *pb.HealthCheckRequest, stream pb.Health_WatchStream) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
