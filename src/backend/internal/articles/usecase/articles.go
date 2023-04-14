package usecase

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/articles/entity"
)

type ArticlesUseCase struct {
	repo ArticlesRepo
}

func NewArticlesUseCase(repo ArticlesRepo) *ArticlesUseCase {
	return &ArticlesUseCase{repo}
}

// Check for implementation
var _ Articles = (*ArticlesUseCase)(nil)

func (uc ArticlesUseCase) Create(ctx context.Context, a entity.Article) (*entity.Article, error) {
	e, err := uc.repo.Create(ctx, a)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Create - uc.repo.Create: %w", err)
	}

	return e, nil
}

func (uc ArticlesUseCase) GetByCustomId(ctx context.Context, customId string) (*entity.Article, error) {
	e, err := uc.repo.GetByCustomId(ctx, customId)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - GetByCustomId - uc.repo.GetByCustomId: %w", err)
	}

	return e, nil
}

func (uc ArticlesUseCase) GetMany(ctx context.Context, offset uint32, count uint32, drafts bool) ([]*entity.ShortArticle, error) {
	e, err := uc.repo.GetMany(ctx, offset, count, drafts)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - GetMany - uc.repo.GetMany: %w", err)
	}

	res := make([]*entity.ShortArticle, len(e))
	for i, v := range e {
		res[i] = &entity.ShortArticle{
			Id:           v.Id,
			CustomId:     v.CustomId,
			AuthorId:     v.AuthorId,
			Title:        v.Title,
			Thumbnail:    v.Thumbnail,
			ShortContent: "Lorem ipsum.", // TODO: short content
			IsDraft:      v.IsDraft,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
	}

	return res, nil
}

func (uc ArticlesUseCase) Update(ctx context.Context, a entity.Article) (*entity.Article, error) {
	e, err := uc.repo.Update(ctx, a)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Update - uc.repo.Update: %w", err)
	}

	return e, nil
}

func (uc ArticlesUseCase) Delete(ctx context.Context, id uint32) (*entity.Article, error) {
	e, err := uc.repo.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("ArticlesUseCase - Delete - uc.repo.Delete: %w", err)
	}

	return e, nil
}
