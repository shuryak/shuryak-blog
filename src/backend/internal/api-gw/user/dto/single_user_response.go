package dto

type SingleUserResponse struct {
	Id       int    `json:"id" example:"42"`
	Username string `json:"username" example:"shuryak"`
	Role     string `json:"role" example:"user"`
}
