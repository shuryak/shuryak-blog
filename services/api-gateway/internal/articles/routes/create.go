package routes

import (
	"api-gateway/internal/articles/pb"
	"api-gateway/internal/errors"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
	"time"
)

type CreateRequest struct {
	CustomId  string                 `json:"custom_id" binding:"min=3,max=20,required" example:"article-url"`
	Title     string                 `json:"title" binding:"min=5,max=150,required" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" binding:"url,required" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content" binding:"required"`
}

type CreateResponse struct {
	Id        uint32                 `json:"id" example:"1000"`
	CustomId  string                 `json:"custom_id" example:"article-url"`
	AuthorId  uint32                 `json:"author_id" example:"42"`
	Title     string                 `json:"title" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at" example:"2022-10-07T14:26:06.510465Z"`
}

// Create
// @Summary     Creates an article
// @Description Creates an article
// @Accept      json
// @Produce     json
// @Param       request body     CreateRequest true "article to create"
// @Success     200     {object} CreateResponse
// @Failure     400     {object} errors.Response
// @Failure     500     {object} errors.Response
// @Failure     502     {object} errors.Response
// @Router      /articles/create [post]
func (r *Routes) Create(ctx *gin.Context) {
	var req CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		r.l.Error(err, "articles - routes - Create")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	content, err := structpb.NewStruct(req.Content)
	if err != nil {
		r.l.Error(err, "articles - routes - Create")
		errors.ErrorResponse(ctx, http.StatusInternalServerError, "some problems")
	}

	authorId, _ := ctx.Get("user_id")

	article, err := r.c.Create(context.Background(), &pb.CreateRequest{
		CustomId:  req.CustomId,
		AuthorId:  authorId.(uint32),
		Title:     req.Title,
		Thumbnail: req.Thumbnail,
		Content:   content,
	})
	if err != nil {
		r.l.Error(err, "articles - routes - Create")
		errors.ErrorResponse(ctx, http.StatusBadGateway, "service error")
		return
	}

	res := CreateResponse{
		Id:        article.Id,
		CustomId:  article.CustomId,
		AuthorId:  article.AuthorId,
		Title:     article.Title,
		Thumbnail: article.Thumbnail,
		Content:   article.Content.AsMap(),
		CreatedAt: article.CreatedAt.AsTime(),
	}

	ctx.JSON(http.StatusCreated, &res)
}