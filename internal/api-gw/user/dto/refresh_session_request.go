package dto

type RefreshSessionRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
