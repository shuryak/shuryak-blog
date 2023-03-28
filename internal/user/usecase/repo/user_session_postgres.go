package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/shuryak/shuryak-blog/internal/user/entity"
	"github.com/shuryak/shuryak-blog/internal/user/usecase"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
)

type UserSessionsRepo struct {
	*postgres.Postgres
}

func NewUserSessionsRepo(pg *postgres.Postgres) *UserSessionsRepo {
	return &UserSessionsRepo{pg}
}

// Check for implementation
var _ usecase.UserSessionsRepo = (*UserSessionsRepo)(nil)

func (r UserSessionsRepo) Create(ctx context.Context, us entity.UserSession) (*entity.UserSession, error) {
	sql, args, err := r.Builder.
		Insert("user_sessions").
		Columns("user_id", "refresh_token", "expires_at", "created_at").
		Values(us.UserId, us.RefreshToken, us.ExpiresAt, us.CreatedAt).
		Suffix("RETURNING \"user_id\", \"refresh_token\", \"expires_at\", \"updated_at\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresUserSessionStore - Create - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	storedUserSession := entity.UserSession{}
	if err = row.Scan(
		&storedUserSession.UserId,
		&storedUserSession.RefreshToken,
		&storedUserSession.ExpiresAt,
		&storedUserSession.UpdatedAt,
		&storedUserSession.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("PostgresUserSessionStore - Create - row.Scan: %w", err)
	}

	return &storedUserSession, nil
}

func (r UserSessionsRepo) Get(ctx context.Context, id uint32) (*entity.UserSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r UserSessionsRepo) Update(ctx context.Context, us entity.UserSession, refreshToken string) (*entity.UserSession, error) {
	clauses := make(map[string]interface{})
	clauses["user_id"] = us.UserId
	clauses["refresh_token"] = us.RefreshToken
	clauses["expires_at"] = us.ExpiresAt
	clauses["updated_at"] = us.UpdatedAt

	sql, args, err := r.Builder.
		Update("user_sessions").
		SetMap(clauses).
		Where(squirrel.Eq{"user_id": us.UserId, "refresh_token": refreshToken}).
		Suffix("RETURNING \"user_id\", \"refresh_token\", \"expires_at\", \"updated_at\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresUserSessionStore - Update - s.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	storedUserSession := entity.UserSession{}
	if err = row.Scan(
		&storedUserSession.UserId,
		&storedUserSession.RefreshToken,
		&storedUserSession.ExpiresAt,
		&storedUserSession.UpdatedAt,
		&storedUserSession.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("PostgresUserSessionStore - Update - row.Scan: %w", err)
	}

	return &storedUserSession, nil
}

func (r UserSessionsRepo) Delete(ctx context.Context, us entity.UserSession) (*entity.UserSession, error) {
	sql, args, err := r.Builder.
		Delete("user_sessions").
		Where(squirrel.Eq{"user_id": us.UserId, "refresh_token": us.RefreshToken}).
		Suffix("RETURNING \"user_id\", \"refresh_token\", \"expires_at\", \"updated_at\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresUserSessionStore - Delete - s.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	storedUserSession := entity.UserSession{}
	if err = row.Scan(&us.UserId, &us.RefreshToken, &us.ExpiresAt, &us.UpdatedAt, &us.CreatedAt); err != nil {
		return nil, fmt.Errorf("PostgresUserSessionStore - Delete - row.Scan: %w", err)
	}

	return &storedUserSession, nil
}
