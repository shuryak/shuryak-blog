package dto

type RefreshSessionRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
}
