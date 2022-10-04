package entity

type Article struct {
	Id       int                    `json:"id"`
	CustomId string                 `json:"custom_id"`
	AuthorId int                    `json:"author_id"`
	Title    string                 `json:"title"`
	Content  map[string]interface{} `json:"content"`
}
