package entity

import (
	"github.com/google/uuid"
	"time"
)

type UserSession struct {
	// TODO: Fingerprint
	Id        uuid.UUID
	UserId    uint32
	ExpiresAt time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}
