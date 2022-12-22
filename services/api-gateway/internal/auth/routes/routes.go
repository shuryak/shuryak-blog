package routes

import (
	"api-gateway/internal/auth/pb"
	"api-gateway/pkg/logger"
)

type Routes struct {
	c pb.AuthClient
	l logger.Interface
}

func NewRoutes(c pb.AuthClient, l logger.Interface) *Routes {
	return &Routes{c, l}
}
