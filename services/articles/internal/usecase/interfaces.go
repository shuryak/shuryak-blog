package usecase

import (
	"context"
	"github.com/shuryak/shuryak-blog/internal/entity"
)

type (
	// Article - .
	Article interface {
		Create(ctx context.Context, a entity.Article) (*entity.Article, error)
		GetById(ctx context.Context, id uint) (*entity.Article, error)
		GetMany(ctx context.Context, offset uint, count uint) ([]entity.Article, error)
		Update(ctx context.Context, a entity.Article) (*entity.Article, error)
		Delete(ctx context.Context, id uint) (*entity.Article, error)
	}

	ArticlesRepo interface {
		Create(ctx context.Context, a entity.Article) (*entity.Article, error)
		GetById(ctx context.Context, id uint) (*entity.Article, error)
		GetMany(ctx context.Context, offset uint, count uint) ([]entity.Article, error)
		Update(ctx context.Context, a entity.Article) (*entity.Article, error)
		Delete(ctx context.Context, id uint) (*entity.Article, error)
	}
)
