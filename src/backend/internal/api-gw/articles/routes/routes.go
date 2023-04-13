package routes

import (
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
)

type Routes struct {
	c pb.ArticlesService
	l logger.Interface
}

func NewRoutes(c pb.ArticlesService, l logger.Interface) *Routes {
	return &Routes{c, l}
}
