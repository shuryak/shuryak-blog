package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/shuryak/shuryak-blog/internal/entity"
	"github.com/shuryak/shuryak-blog/internal/usecase"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
)

const _defaultEntityCap = 64

type ArticlesRepo struct {
	*postgres.Postgres
}

// Check for implementation
var _ usecase.ArticlesRepo = (*ArticlesRepo)(nil)

func New(pg *postgres.Postgres) *ArticlesRepo {
	return &ArticlesRepo{pg}
}

func (r *ArticlesRepo) Create(ctx context.Context, a entity.Article) (*entity.Article, error) {
	sql, args, err := r.Builder.
		Insert("articles").
		Columns("custom_id, author_id, title, thumbnail, content").
		Values(a.CustomId, a.AuthorId, a.Title, a.Thumbnail, a.Content).
		Suffix("RETURNING \"id\", \"custom_id\", \"author_id\", \"title\", \"thumbnail\", \"content\", " +
			"\"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Create - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	newArticle := entity.Article{}
	if err = row.Scan(&newArticle.Id, &newArticle.CustomId, &newArticle.AuthorId, &newArticle.Title,
		&newArticle.Thumbnail, &newArticle.Content, &newArticle.CreatedAt); err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Create - row.Scan: %w", err)
	}

	return &newArticle, nil
}

func (r *ArticlesRepo) GetById(ctx context.Context, id uint) (*entity.Article, error) {
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

func (r *ArticlesRepo) GetMany(ctx context.Context, offset uint, count uint) ([]entity.Article, error) {
	sql, args, err := r.Builder.
		Select("id, custom_id, author_id, title, thumbnail, content, created_at").
		From("articles").
		Where(squirrel.GtOrEq{"id": offset}).
		Limit(uint64(count)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - GetMany - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - GetMany - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Article, 0, count)

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

func (r *ArticlesRepo) Update(ctx context.Context, a entity.Article) (*entity.Article, error) {
	clauses := make(map[string]interface{})
	clauses["custom_id"] = a.CustomId
	clauses["title"] = a.Title
	clauses["thumbnail"] = a.Thumbnail
	clauses["content"] = a.Content
	// TODO: updated_at

	sql, args, err := r.Builder.
		Update("articles").
		SetMap(clauses).
		Where(squirrel.Eq{"id": a.Id}).
		Suffix("RETURNING \"id\", \"custom_id\", \"author_id\", \"title\", \"thumbnail\", \"content\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Update - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	newArticle := entity.Article{}
	if err = row.Scan(&newArticle.Id, &newArticle.CustomId, &newArticle.AuthorId, &newArticle.Title,
		&newArticle.Thumbnail, &newArticle.Content, &newArticle.CreatedAt); err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Update - row.Scan: %w", err)
	}

	return &newArticle, nil
}

func (r *ArticlesRepo) Delete(ctx context.Context, id uint) (*entity.Article, error) {
	sql, args, err := r.Builder.
		Delete("articles").
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING \"id\", \"custom_id\", \"author_id\", \"title\", \"thumbnail\", \"content\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Delete - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	a := entity.Article{}
	if err = row.Scan(&a.Id, &a.CustomId, &a.AuthorId, &a.Title, &a.Thumbnail, &a.Content, &a.CreatedAt); err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Delete - row.Scan: %w", err)
	}

	return &a, nil
}
