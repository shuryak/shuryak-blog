package entity

import "time"

type Article struct {
	Id        uint32
	CustomId  string
	AuthorId  uint32
	Title     string
	Thumbnail string
	Content   map[string]interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
}
