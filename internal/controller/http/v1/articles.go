package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shuryak-blog/internal/entity"
	"shuryak-blog/internal/usecase"
)

type articlesRoutes struct {
	a usecase.Article
	// TODO: logger
}

func newArticlesRoutes(handler *gin.RouterGroup, a usecase.Article) {
	r := &articlesRoutes{a}

	h := handler.Group("/articles")
	{
		h.GET("/getAll", r.getAll)
		h.POST("/create", r.create)
	}
}

type articlesResponse struct {
	Articles []entity.Article `json:"articles"`
}

func (r *articlesRoutes) getAll(c *gin.Context) {
	articles, err := r.a.GetMany(c.Request.Context())
	if err != nil {
		// TODO: do log
		errorResponse(c, http.StatusInternalServerError, "database problems: "+err.Error())

		return
	}

	c.JSON(http.StatusOK, articlesResponse{articles})
}

// TODO: make example attr
type createArticleRequest struct {
	CustomId  string                 `json:"custom_id" binding:"min=3,max=20,required"`
	AuthorId  int                    `json:"author_id" binding:"required"`
	Title     string                 `json:"title" binding:"min=5,max=150,required"`
	Thumbnail string                 `json:"thumbnail" binding:"url,required"`
	Content   map[string]interface{} `json:"content" binding:"required"`
}

func (r *articlesRoutes) create(c *gin.Context) {
	var request createArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO: do log
		errorResponse(c, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	article, err := r.a.Create(c.Request.Context(), entity.Article{
		CustomId: request.CustomId,
		AuthorId: request.AuthorId,
		Title:    request.Title,
		Content:  request.Content,
	})
	if err != nil {
		// TODO: do log and delete next line
		fmt.Println("problems: ", err.Error())
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, article)
}
