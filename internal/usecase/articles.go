package usecase

import (
	"context"
	"fmt"
	"shuryak-blog/internal/entity"
)

type ArticlesUseCase struct {
	repo ArticlesRepo
}

// Check for implementation
var _ Article = (*ArticlesUseCase)(nil)

func New(r ArticlesRepo) *ArticlesUseCase {
	return &ArticlesUseCase{r}
}

func (uc *ArticlesUseCase) Create(ctx context.Context, a entity.Article) (*entity.Article, error) {
	id, err := uc.repo.Create(context.Background(), a)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Create - s.repo.Store: %w", err)
	}

	articleEntity, err := uc.repo.GetById(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Create - s.repo.GetById: %w", err)
	}

	return articleEntity, nil
}

func (uc *ArticlesUseCase) GetMany(ctx context.Context, offset uint, count uint) ([]entity.Article, error) {
	articles, err := uc.repo.GetMany(ctx, offset, count)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - GetMany - s.repo.GetMany: %w", err)
	}

	return articles, nil
}
