package dto

type RefreshSessionRequest struct {
	Username     string `json:"username" binding:"required" example:"username"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
