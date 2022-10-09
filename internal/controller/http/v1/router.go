package v1

import (
	"github.com/gin-gonic/gin"
	"shuryak-blog/internal/usecase"
)

func NewRouter(handler *gin.Engine, a usecase.Article) {
	// TODO: more options

	// Routers
	h := handler.Group("/v1")
	{
		newArticlesRoutes(h, a)
	}
}
