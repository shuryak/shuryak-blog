package dto

type ManyArticlesResponse struct {
	Articles []*ShortArticleResponse `json:"articles"`
}
