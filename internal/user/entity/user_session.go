package entity

import "time"

type UserSession struct {
	// TODO: IP and UserAgent
	UserId       uint32
	RefreshToken string
	ExpiresAt    time.Time
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
