package entity

import "time"

type Article struct {
	Id        uint                   `json:"id"`
	CustomId  string                 `json:"custom_id"`
	AuthorId  uint                   `json:"author_id"`
	Title     string                 `json:"title"`
	Thumbnail string                 `json:"thumbnail"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at"`
}
