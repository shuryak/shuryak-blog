package usecase

import (
	"context"
	"fmt"
	"shuryak-blog/internal/entity"
)

type ArticlesUseCase struct {
	repo ArticlesRepo
}

func New(r ArticlesRepo) *ArticlesUseCase {
	return &ArticlesUseCase{r}
}

func (uc *ArticlesUseCase) Create(ctx context.Context, a entity.Article) (entity.Article, error) {
	if err := uc.repo.Create(context.Background(), a); err != nil {
		return entity.Article{}, fmt.Errorf("ArticlesUseCase - Create - s.repo.Store: %w", err)
	}

	return a, nil
}

func (uc *ArticlesUseCase) GetMany(ctx context.Context) ([]entity.Article, error) {
	articles, err := uc.repo.GetMany(ctx)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - GetMany - s.repo.GetMany: %w", err)
	}

	return articles, nil
}
