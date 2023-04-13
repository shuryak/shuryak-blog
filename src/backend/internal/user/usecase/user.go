package usecase

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/user/entity"
)

type UsersUseCase struct {
	repo UsersRepo
}

func NewUsersUseCase(repo UsersRepo) *UsersUseCase {
	return &UsersUseCase{repo}
}

// Check for implementation
var _ Users = (*UsersUseCase)(nil)

func (uc UsersUseCase) Create(ctx context.Context, u entity.User) (*entity.User, error) {
	e, err := uc.repo.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - Create - uc.repo.Create: %w", err)
	}

	return e, nil
}

func (uc UsersUseCase) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	e, err := uc.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - GetByUsername - uc.repo.GetByUsername: %w", err)
	}

	return e, nil
}
