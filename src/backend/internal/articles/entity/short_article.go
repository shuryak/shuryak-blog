package entity

import "time"

type ShortArticle struct {
	Id           uint32
	CustomId     string
	AuthorId     uint32
	Title        string
	Thumbnail    string
	ShortContent string
	IsDraft      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
