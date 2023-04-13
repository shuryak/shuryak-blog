package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/shuryak/shuryak-blog/internal/articles/entity"
	"github.com/shuryak/shuryak-blog/internal/articles/usecase"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
)

type ArticlesRepo struct {
	*postgres.Postgres
}

func NewArticlesRepo(pg *postgres.Postgres) *ArticlesRepo {
	return &ArticlesRepo{pg}
}

// Check for implementation
var _ usecase.ArticlesRepo = (*ArticlesRepo)(nil)

func (r ArticlesRepo) Create(ctx context.Context, a entity.Article) (*entity.Article, error) {
	sql, args, err := r.Builder.
		Insert("articles").
		Columns("custom_id", "author_id", "title", "thumbnail", "content", "created_at", "updated_at").
		Values(a.CustomId, a.AuthorId, a.Title, a.Thumbnail, a.Content, a.CreatedAt, a.UpdatedAt).
		Suffix("RETURNING \"id\", \"custom_id\", \"author_id\", \"title\", \"thumbnail\", \"content\", " +
			"\"created_at\", \"updated_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Create - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	storedArticle := entity.Article{}
	if err = row.Scan(&storedArticle.Id, &storedArticle.CustomId, &storedArticle.AuthorId, &storedArticle.Title,
		&storedArticle.Thumbnail, &storedArticle.Content, &storedArticle.CreatedAt, &storedArticle.UpdatedAt); err != nil {
		return nil, fmt.Errorf("ArticlesRepo - Create - row.Scan: %w", err)
	}

	return &storedArticle, nil
}

func (r ArticlesRepo) GetByCustomId(ctx context.Context, customId string) (*entity.Article, error) {
	sql, args, err := r.Builder.
		Select("id", "custom_id", "author_id", "title", "content", "thumbnail", "created_at", "updated_at").
		From("articles").
		Where(squirrel.Eq{"custom_id": customId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ArticlesRepo - GetByCustomId - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	a := entity.Article{}
	err = row.Scan(&a.Id, &a.CustomId, &a.AuthorId, &a.Title, &a.Content, &a.Thumbnail, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("PostgresArticlesStore - GetByCustomId - row.Scan: %w", err)
	}

	return &a, nil
}

func (r ArticlesRepo) GetMany(ctx context.Context, offset uint32, count uint32) ([]entity.Article, error) {
	sql, args, err := r.Builder.
		Select("id", "custom_id", "author_id", "title", "thumbnail", "content", "created_at", "updated_at").
		From("articles").
		Offset(uint64(offset)).
		Limit(uint64(count)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresArticlesStore - GetMany - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PostgresArticlesStore - GetMany - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Article, 0, count)

	for rows.Next() {
		e := entity.Article{}

		err = rows.Scan(&e.Id, &e.CustomId, &e.AuthorId, &e.Title, &e.Thumbnail, &e.Content, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("PostgresArticlesStore - GetMany - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (r ArticlesRepo) Update(ctx context.Context, a entity.Article) (*entity.Article, error) {
	clauses := make(map[string]interface{})
	clauses["custom_id"] = a.CustomId
	clauses["title"] = a.Title
	clauses["thumbnail"] = a.Thumbnail
	clauses["content"] = a.Content

	sql, args, err := r.Builder.
		Update("articles").
		SetMap(clauses).
		Where(squirrel.Eq{"id": a.Id}).
		Suffix("RETURNING \"id\", \"custom_id\", \"author_id\", \"title\", \"thumbnail\", \"content\", " +
			"\"created_at\", \"updated_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresArticlesStore - Update - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	storedArticle := entity.Article{}
	if err = row.Scan(&storedArticle.Id, &storedArticle.CustomId, &storedArticle.AuthorId, &storedArticle.Title,
		&storedArticle.Thumbnail, &storedArticle.Content, &storedArticle.CreatedAt); err != nil {
		return nil, fmt.Errorf("PostgresArticlesStore - Update - row.Scan: %w", err)
	}

	return &storedArticle, nil
}

func (r ArticlesRepo) Delete(ctx context.Context, id uint32) (*entity.Article, error) {
	sql, args, err := r.Builder.
		Delete("articles").
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING \"id\", \"custom_id\", \"author_id\", \"title\", \"thumbnail\", \"content\", " +
			"\"created_at\", \"updated_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresArticlesStore - Delete - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	storedArticle := entity.Article{}
	if err = row.Scan(&storedArticle.Id, &storedArticle.CustomId, &storedArticle.AuthorId, &storedArticle.Title,
		&storedArticle.Thumbnail, &storedArticle.Content, &storedArticle.CreatedAt); err != nil {
		return nil, fmt.Errorf("PostgresArticlesStore - Delete - row.Scan: %w", err)
	}

	return &storedArticle, nil
}
