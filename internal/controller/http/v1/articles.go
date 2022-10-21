package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shuryak-blog/internal/entity"
	"shuryak-blog/internal/usecase"
	"shuryak-blog/pkg/logger"
	"time"
)

type articlesRoutes struct {
	a usecase.Article
	l logger.Interface
}

func newArticlesRoutes(handler *gin.RouterGroup, a usecase.Article, l logger.Interface) {
	r := &articlesRoutes{a, l}

	h := handler.Group("/articles")
	{
		h.GET("/getById", r.getById)
		h.GET("/getMany", r.getMany)
		h.POST("/create", r.create)
		h.PUT("/update", r.update)
		h.DELETE("/delete", r.delete)
	}
}

type createRequest struct {
	CustomId  string                 `json:"custom_id" binding:"min=3,max=20,required" example:"article-url"`
	AuthorId  uint                   `json:"author_id" binding:"required" example:"42"`
	Title     string                 `json:"title" binding:"min=5,max=150,required" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" binding:"url,required" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content" binding:"required"`
}

type articleResponse struct {
	Id        uint                   `json:"id" example:"1000"`
	CustomId  string                 `json:"custom_id" example:"article-url"`
	AuthorId  uint                   `json:"author_id" example:"42"`
	Title     string                 `json:"title" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at" example:"2022-10-07T14:26:06.510465Z"`
}

// @Summary     Creates an article
// @Description Creates an article
// @Accept      json
// @Produce     json
// @Param       request body     createRequest true "article to create"
// @Success     200     {object} getManyResponse
// @Failure     400     {object} response
// @Failure     500     {object} response
// @Router      /articles/create [post]
func (r *articlesRoutes) create(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - create")
		validationErrorResponse(c, http.StatusBadRequest, err) // TODO: make good error structs
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
		r.l.Error(err, "http - v1 - create")
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

type articleIdRequest struct {
	Id int `form:"id" binding:"min=0,required" example:"42"`
}

// @Summary Gets article by ID
// @Description Gets article by ID
// @Produce json
// @Param   id query    int true "ID to get"
// @Success 200    {object} articleResponse
// @Failure 400    {object} response
// @Failure 500    {object} response
// @Router  /articles/getById [get]
func (r *articlesRoutes) getById(c *gin.Context) {
	var request articleIdRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		r.l.Error(err, "http - v1 - getById")
		validationErrorResponse(c, http.StatusBadRequest, err) // TODO: make good error structs
		return
	}

	a, err := r.a.GetById(c.Request.Context(), uint(request.Id))
	if err != nil {
		r.l.Error(err, "http - v1 - getById")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	ar := articleResponse{
		Id:        a.Id,
		CustomId:  a.CustomId,
		AuthorId:  a.AuthorId,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   a.Content,
		CreatedAt: a.CreatedAt,
	}

	c.JSON(http.StatusOK, ar)
}

type getManyRequest struct {
	Offset int `form:"offset" binding:"min=0,required" example:"42"`
	Count  int `form:"count" binding:"min=1,required" example:"10"`
}

type getManyResponse struct {
	Articles []articleResponse `json:"articles"`
}

// @Summary Gets collection of articles
// @Description Gets collection of articles
// @Produce json
// @Param   offset query    int true "offset to get"
// @Param   count  query    int true "count to get"
// @Success 200    {object} getManyResponse
// @Failure 400    {object} response
// @Failure 500    {object} response
// @Router  /articles/getMany [get]
func (r *articlesRoutes) getMany(c *gin.Context) {
	var request getManyRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		r.l.Error(err, "http - v1 - getMany")
		validationErrorResponse(c, http.StatusBadRequest, err) // TODO: make good error structs
		return
	}

	articles, err := r.a.GetMany(c.Request.Context(), uint(request.Offset), uint(request.Count))
	if err != nil {
		r.l.Error(err, "http - v1 - getMany")
		errorResponse(c, http.StatusInternalServerError, "database problems")

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

	c.JSON(http.StatusOK, getManyResponse{ar})
}

type updateRequest struct {
	Id        uint                   `json:"id" binding:"min=0,required" example:"1000"`
	CustomId  string                 `json:"custom_id" binding:"min=3,max=20,required" example:"article-url"`
	AuthorId  uint                   `json:"author_id" binding:"required" example:"42"`
	Title     string                 `json:"title" binding:"min=5,max=150,required" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" binding:"url,required" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content" binding:"required"`
}

// @Summary Updates article by ID
// @Description Updates article by ID
// @Accept  json
// @Produce json
// @Param       request body     updateRequest true "article to update"
// @Success     200     {object} articleResponse
// @Failure     400     {object} response
// @Failure     500     {object} response
// @Router      /articles/update [put]
func (r *articlesRoutes) update(c *gin.Context) {
	var request updateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - update")
		validationErrorResponse(c, http.StatusBadRequest, err) // TODO: make good error structs
		return
	}

	a, err := r.a.Update(c.Request.Context(), entity.Article{
		Id:        request.Id,
		CustomId:  request.CustomId,
		AuthorId:  request.AuthorId,
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		Content:   request.Content,
	})
	if err != nil {
		r.l.Error(err, "http - v1 - update")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	ar := articleResponse{
		Id:        a.Id,
		CustomId:  a.CustomId,
		AuthorId:  a.AuthorId,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   a.Content,
	}

	c.JSON(http.StatusOK, ar)
}

// @Summary Deletes article by ID
// @Description Deletes article by ID
// @Produce json
// @Param   id query    int true "ID to delete"
// @Success     200     {object} articleResponse
// @Failure     400     {object} response
// @Failure     500     {object} response
// @Router      /articles/delete [delete]
func (r *articlesRoutes) delete(c *gin.Context) {
	var request articleIdRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		r.l.Error(err, "http - v1 - delete")
		validationErrorResponse(c, http.StatusBadRequest, err) // TODO: make good error structs
		return
	}

	a, err := r.a.Delete(c.Request.Context(), uint(request.Id))
	if err != nil {
		r.l.Error(err, "http - v1 - delete")
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	ar := articleResponse{
		Id:        a.Id,
		CustomId:  a.CustomId,
		AuthorId:  a.AuthorId,
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   a.Content,
		CreatedAt: a.CreatedAt,
	}

	c.JSON(http.StatusOK, ar)
}
