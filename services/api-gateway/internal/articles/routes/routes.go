package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/pkg/logger"
)

type Routes struct {
	c pb.ArticlesClient
	l logger.Interface
}

func NewRoutes(c pb.ArticlesClient, l logger.Interface) *Routes {
	return &Routes{c, l}
}
