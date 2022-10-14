package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
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

func (r *ArticlesRepo) Create(ctx context.Context, a entity.Article) (int64, error) {
	sql, args, err := r.Builder.
		Insert("articles").
		Columns("custom_id, author_id, title, thumbnail, content").
		Values(a.CustomId, a.AuthorId, a.Title, a.Thumbnail, a.Content).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("ArticlesRepo - Create - r.Builder: %w", err)
	}

	var id int64
	row := r.Pool.QueryRow(ctx, sql, args...)
	if err = row.Scan(&id); err != nil {
		return 0, fmt.Errorf("ArticlesRepo - Create - row.Scan: %w", err)
	}

	return id, nil
}

func (r *ArticlesRepo) GetById(ctx context.Context, id int64) (*entity.Article, error) {
	sql, args, err := r.Builder.
		Select("id, custom_id, author_id, title, content, thumbnail, created_at").
		From("articles").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - GetById - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	a := entity.Article{}
	err = row.Scan(&a.Id, &a.CustomId, &a.AuthorId, &a.Title, &a.Content, &a.Thumbnail, &a.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - GetById - row.Scan: %w", err)
	}

	return &a, nil
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
