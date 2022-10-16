package usecase

import (
	"context"
	"shuryak-blog/internal/entity"
)

type (
	// Article - .
	Article interface {
		Create(ctx context.Context, a entity.Article) (*entity.Article, error)
		GetMany(ctx context.Context) ([]entity.Article, error)
	}

	ArticlesRepo interface {
		Create(ctx context.Context, a entity.Article) (int, error)
		GetById(ctx context.Context, id int) (*entity.Article, error)
		GetMany(ctx context.Context) ([]entity.Article, error)
		Delete(ctx context.Context, id int) (*entity.Article, error)
	}
)
