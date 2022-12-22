package dto

type ArticleIdRequest struct {
	Id uint32 `json:"id" form:"id" binding:"min=1,required" example:"42"`
}
