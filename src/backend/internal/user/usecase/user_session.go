package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shuryak/shuryak-blog/internal/user/entity"
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
	now := time.Now()

	e, err := uc.repo.Create(ctx, entity.UserSession{
		UserId:    userId,
		ExpiresAt: now.Add(uc.sessionDuration),
		UpdatedAt: now,
		CreatedAt: now,
	})
	if err != nil {
		return nil, fmt.Errorf("UserSessionUseCase - Add - uc.repo.Create: %w", err)
	}

	return e, nil
}

func (uc UserSessionUseCase) Refresh(ctx context.Context, sessionId uuid.UUID, userId uint32) (*entity.UserSession, error) {
	now := time.Now()

	e, err := uc.repo.Update(ctx, sessionId, entity.UserSession{
		UserId:    userId,
		ExpiresAt: now.Add(uc.sessionDuration),
		UpdatedAt: now,
	})
	if err != nil {
		return nil, fmt.Errorf("UserSessionUseCase - Refresh - uc.repo.Update: %w", err)
	}

	return e, nil
}
