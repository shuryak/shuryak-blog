package dto

type ArticleCustomIdRequest struct {
	CustomId string `form:"custom_id" json:"custom_id" example:"article-url"`
}
