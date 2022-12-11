package routes

import (
	"api-gateway/internal/articles/pb"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
)

type CreateArticleRequest struct {
	CustomId  string                 `json:"custom_id" binding:"min=3,max=20,required" example:"article-url"`
	Title     string                 `json:"title" binding:"min=5,max=150,required" example:"How to ..."`
	Thumbnail string                 `json:"thumbnail" binding:"url,required" example:"https://smth.com/thumbnail.png"`
	Content   map[string]interface{} `json:"content" binding:"required"`
}

func CreateArticle(ctx *gin.Context, c pb.ArticlesClient) {
	req := CreateArticleRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	content, err := structpb.NewStruct(req.Content)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
	}

	authorId, _ := ctx.Get("user_id")

	res, err := c.Create(context.Background(), &pb.CreateRequest{
		CustomId:  req.CustomId,
		AuthorId:  authorId.(uint32),
		Title:     req.Title,
		Thumbnail: req.Thumbnail,
		Content:   content,
	})

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(http.StatusCreated, &res)
}
