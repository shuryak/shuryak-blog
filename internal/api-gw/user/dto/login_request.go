package dto

type LoginRequest struct {
	Username string `json:"username" binding:"min=3,max=20,required" example:"shuryak"`
	Password string `json:"password" binding:"min=6,max=128,required" example:"password@@10_25!!0"` // TODO: min length
}
