package routes

import (
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/user"
)

type Routes struct {
	c pb.UserService
	l logger.Interface
}

func NewRoutes(c pb.UserService, l logger.Interface) *Routes {
	return &Routes{c, l}
}
