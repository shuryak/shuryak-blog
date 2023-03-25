package usecase

import (
	"context"
	"github.com/shuryak/shuryak-blog/internal/articles/entity"
)

type (
	Articles interface {
		Create(ctx context.Context, a entity.Article) (*entity.Article, error)
		GetByCustomId(ctx context.Context, customId string) (*entity.Article, error)
		GetMany(ctx context.Context, offset uint32, count uint32) ([]entity.Article, error)
		Update(ctx context.Context, a entity.Article) (*entity.Article, error)
		Delete(ctx context.Context, id uint32) (*entity.Article, error)
	}

	ArticlesRepo interface {
		Create(ctx context.Context, a entity.Article) (*entity.Article, error)
		GetByCustomId(ctx context.Context, customId string) (*entity.Article, error)
		GetMany(ctx context.Context, offset uint32, count uint32) ([]entity.Article, error)
		Update(ctx context.Context, a entity.Article) (*entity.Article, error)
		Delete(ctx context.Context, id uint32) (*entity.Article, error)
	}
)
