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

// TODO: deal with the context

func (uc *ArticlesUseCase) Create(ctx context.Context, a entity.Article) (*entity.Article, error) {
	articleEntity, err := uc.repo.Create(context.Background(), a)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Create - s.repo.Create: %w", err)
	}

	return articleEntity, nil
}

func (uc *ArticlesUseCase) GetById(ctx context.Context, id uint) (*entity.Article, error) {
	a, err := uc.repo.GetById(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - GetById - s.repo.GetById")
	}

	return a, nil
}

func (uc *ArticlesUseCase) GetMany(ctx context.Context, offset uint, count uint) ([]entity.Article, error) {
	a, err := uc.repo.GetMany(context.Background(), offset, count)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - GetMany - s.repo.GetMany: %w", err)
	}

	return a, nil
}

func (uc *ArticlesUseCase) Update(ctx context.Context, a entity.Article) (*entity.Article, error) {
	updatedArticleEntity, err := uc.repo.Update(context.Background(), a)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Update - s.repo.Update: %w", err)
	}

	return updatedArticleEntity, nil
}

func (uc *ArticlesUseCase) Delete(ctx context.Context, id uint) (*entity.Article, error) {
	deletedArticleEntity, err := uc.repo.Delete(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Delete - s.repo.Delete: %w", err)
	}

	return deletedArticleEntity, nil
}
