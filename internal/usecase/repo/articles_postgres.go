package repo

import (
	"context"
	"fmt"
	"shuryak-blog/internal/entity"
	"shuryak-blog/pkg/postgres"
)

const _defaultEntityCap = 64

type ArticlesRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *ArticlesRepo {
	return &ArticlesRepo{pg}
}

func (r *ArticlesRepo) Create(ctx context.Context, a entity.Article) error {
	sql, args, err := r.Builder.
		Insert("articles").
		Columns("custom_id, author_id, title, content").
		Values(a.CustomId, a.AuthorId, a.Title, a.Content).
		ToSql()
	if err != nil {
		return fmt.Errorf("ArticlesRepo - Create - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ArticlesRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *ArticlesRepo) GetMany(ctx context.Context) ([]entity.Article, error) {
	sql, _, err := r.Builder.
		Select("id, custom_id, author_id, title, thumbnail, content, created_at").
		From("articles").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - GetMany - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - GetMany - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Article, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Article{}

		err = rows.Scan(&e.Id, &e.CustomId, &e.AuthorId, &e.Title, &e.Thumbnail, &e.Content, &e.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("ArticlesRepo - GetMany - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
