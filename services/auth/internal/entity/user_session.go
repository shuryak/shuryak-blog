package entity

import "time"

type UserSession struct {
	UserId       uint
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
