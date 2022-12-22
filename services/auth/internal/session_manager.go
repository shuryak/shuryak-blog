package internal

import (
	"auth/internal/entity"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type SessionManager struct {
	uss                UserSessionStore
	refreshTokenLength uint
	sessionDuration    time.Duration
}

func NewSessionManager(uss UserSessionStore, refreshTokenLength uint, sessionDuration time.Duration) *SessionManager {
	return &SessionManager{uss, refreshTokenLength, sessionDuration}
}

func (m *SessionManager) Start(user *entity.User) (*entity.UserSession, error) {
	refreshToken, err := m.generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("refresh token generation error")
	}

	createdAt := time.Now()
	us := entity.UserSession{
		UserId:       user.Id,
		RefreshToken: refreshToken,
		ExpiresAt:    createdAt.Add(m.sessionDuration),
		CreatedAt:    createdAt,
	}

	newUs, err := m.uss.Create(context.Background(), us)
	if err != nil {
		return nil, fmt.Errorf("session creating error")
	}

	return newUs, nil
}

func (m *SessionManager) generateRefreshToken() (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	result := make([]byte, m.refreshTokenLength)
	for i := 0; i < int(m.refreshTokenLength); i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}

	return string(result), nil
}
