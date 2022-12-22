package dto

type ManyArticlesResponse struct {
	Articles []SingleArticleResponse `json:"articles"`
}
