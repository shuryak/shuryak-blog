package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shuryak-blog/internal/entity"
	"shuryak-blog/internal/usecase"
	"time"
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
	Articles []articleResponse `json:"articles"`
}

// @Summary Gets all articles
// @Description Gets all articles
// @Produce json
// @Success 200 {object} articlesResponse
// @Failure 500 {object} response
// @Router /articles/getAll [get]
func (r *articlesRoutes) getAll(c *gin.Context) {
	articles, err := r.a.GetMany(c.Request.Context())
	if err != nil {
		// TODO: do log
		errorResponse(c, http.StatusInternalServerError, "database problems: "+err.Error())

		return
	}

	ar := make([]articleResponse, len(articles))
	for i, a := range articles {
		ar[i] = articleResponse{
			Id:        a.Id,
			CustomId:  a.CustomId,
			AuthorId:  a.AuthorId,
			Title:     a.Title,
			Thumbnail: a.Thumbnail,
			Content:   a.Content,
			CreatedAt: a.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, articlesResponse{ar})
}

// TODO: make example attr
type createArticleRequest struct {
	CustomId  string                 `json:"custom_id" binding:"min=3,max=20,required" example:"article-url"`
	AuthorId  int                    `json:"author_id" binding:"required" example:"42"`
	Title     string                 `json:"title" binding:"min=5,max=150,required" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" binding:"url,required" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content" binding:"required"`
}

type articleResponse struct {
	Id        int                    `json:"id" example:"1000"`
	CustomId  string                 `json:"custom_id" example:"article-url"`
	AuthorId  int                    `json:"author_id" example:"42"`
	Title     string                 `json:"title" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at" example:"2022-10-07T14:26:06.510465Z"`
}

// @Summary Creates an article
// @Description Creates an article
// @Accept json
// @Produce json
// @Param request body createArticleRequest true "article to create"
// @Success 200 {object} articlesResponse
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Router /articles/create [post]
func (r *articlesRoutes) create(c *gin.Context) {
	var request createArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO: do log
		//errorResponse(c, http.StatusBadRequest, "invalid request body: "+err.Error())
		validationErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	article, err := r.a.Create(c.Request.Context(), entity.Article{
		CustomId:  request.CustomId,
		AuthorId:  request.AuthorId,
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		Content:   request.Content,
	})
	if err != nil {
		// TODO: do log and delete next line
		fmt.Println("problems: ", err.Error())
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	ar := articleResponse{
		Id:        article.Id,
		CustomId:  article.CustomId,
		AuthorId:  article.AuthorId,
		Title:     article.Title,
		Thumbnail: article.Thumbnail,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
	}

	c.JSON(http.StatusOK, ar)
}
