package entity

import "time"

type Article struct {
	Id        int                    `json:"id"`
	CustomId  string                 `json:"custom_id"`
	AuthorId  int                    `json:"author_id"`
	Title     string                 `json:"title"`
	Thumbnail string                 `json:"thumbnail"`
	Content   map[string]interface{} `json:"content"`
	CreatedAt time.Time              `json:"created_at"`
}
