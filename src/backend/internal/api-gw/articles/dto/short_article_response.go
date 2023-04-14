package dto

import "time"

type ShortArticleResponse struct {
	Id           int       `json:"id" example:"1000"`
	CustomId     string    `json:"custom_id" example:"article-url"`
	AuthorId     int       `json:"author_id" example:"42"`
	Title        string    `json:"title" example:"How to ..."`
	Thumbnail    string    `json:"thumbnail" example:"https://smth.com/thumbnail.png"`
	ShortContent string    `json:"short_content" example:"Lorem ipsum."`
	IsDraft      bool      `json:"is_draft" example:"true"`
	CreatedAt    time.Time `json:"created_at" example:"2022-10-07T14:26:06.510465Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2022-10-07T14:26:06.510465Z"`
}
