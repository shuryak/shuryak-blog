package usecase

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/user/entity"
	"math/big"
	"time"
)

type UserSessionUseCase struct {
	repo UserSessionsRepo
	// TODO: config?
	refreshTokenLength uint32
	sessionDuration    time.Duration
}

func NewUserSessionsUseCase(repo UserSessionsRepo, refreshTokenLength uint32, sessionDuration time.Duration) *UserSessionUseCase {
	return &UserSessionUseCase{repo, refreshTokenLength, sessionDuration}
}

// Check for implementation
var _ UserSessions = (*UserSessionUseCase)(nil)

func (uc UserSessionUseCase) Add(ctx context.Context, userId uint32) (*entity.UserSession, error) {
	refreshToken, err := uc.generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("UserSessionUseCase - Add - uc.generateRefreshToken: %w", err)
	}

	now := time.Now()

	e, err := uc.repo.Create(ctx, entity.UserSession{
		UserId:       userId,
		RefreshToken: refreshToken,
		ExpiresAt:    now.Add(uc.sessionDuration),
		UpdatedAt:    now,
		CreatedAt:    now,
	})

	return e, nil
}

func (uc UserSessionUseCase) Refresh(ctx context.Context, userId uint32, refreshToken string) (*entity.UserSession, error) {
	newRefreshToken, err := uc.generateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("UserSessionUseCase - Refresh - uc.generateRefreshToken: %w", err)
	}

	now := time.Now()

	e, err := uc.repo.Update(ctx, entity.UserSession{
		UserId:       userId,
		RefreshToken: newRefreshToken,
		ExpiresAt:    now.Add(uc.sessionDuration),
		UpdatedAt:    now,
	}, refreshToken)

	return e, nil
}

func (uc UserSessionUseCase) generateRefreshToken() (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	result := make([]byte, uc.refreshTokenLength)
	for i := 0; i < int(uc.refreshTokenLength); i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}

	return string(result), nil
}
