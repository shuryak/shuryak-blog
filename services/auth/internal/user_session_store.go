package internal

import (
	"auth/internal/entity"
	"auth/pkg/postgres"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type UserSessionStore interface {
	Create(ctx context.Context, us entity.UserSession) (*entity.UserSession, error)
	Update(ctx context.Context, us entity.UserSession, refreshToken string) (*entity.UserSession, error)
	Delete(ctx context.Context, userId uint, refreshToken string) (*entity.UserSession, error)
}

type PostgresUserSessionStore struct {
	*postgres.Postgres
}

// Check for implementation
var _ UserSessionStore = (*PostgresUserSessionStore)(nil)

func NewPostgresUserSessionStore(pg *postgres.Postgres) *PostgresUserSessionStore {
	return &PostgresUserSessionStore{pg}
}

func (s PostgresUserSessionStore) Create(ctx context.Context, us entity.UserSession) (*entity.UserSession, error) {
	sql, args, err := s.Builder.
		Insert("user_sessions").
		Columns("user_id", "refresh_token", "expires_at", "created_at").
		Values(us.UserId, us.RefreshToken, us.ExpiresAt, us.CreatedAt).
		Suffix("RETURNING \"user_id\", \"refresh_token\", \"expires_at\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserSessionStore - Create - s.Builder: %w", err)
	}

	row := s.Pool.QueryRow(ctx, sql, args...)
	newUserSession := entity.UserSession{}
	if err = row.Scan(
		&newUserSession.UserId,
		&newUserSession.RefreshToken,
		&newUserSession.ExpiresAt,
		&newUserSession.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("UserSessionStore - Create - row.Scan: %w", err)
	}

	return &newUserSession, nil
}

func (s PostgresUserSessionStore) Update(
	ctx context.Context,
	us entity.UserSession,
	refreshToken string,
) (*entity.UserSession, error) {
	clauses := make(map[string]interface{})
	clauses["user_id"] = us.UserId
	clauses["refresh_token"] = us.RefreshToken
	clauses["expires_at"] = us.ExpiresAt
	clauses["created_at"] = us.CreatedAt

	sql, args, err := s.Builder.
		Update("user_sessions").
		SetMap(clauses).
		Where(squirrel.Eq{"user_id": us.UserId, "refresh_token": refreshToken}).
		Suffix("RETURNING \"user_id\", \"refresh_token\", \"expires_at\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserSessions - Update - s.Builder: %w", err)
	}

	row := s.Pool.QueryRow(ctx, sql, args...)
	newUserSession := entity.UserSession{}
	if err = row.Scan(
		&newUserSession.UserId,
		&newUserSession.RefreshToken,
		&newUserSession.ExpiresAt,
		&newUserSession.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("UserSessionStore - Update - row.Scan: %w", err)
	}

	return &newUserSession, nil
}

func (s PostgresUserSessionStore) Delete(ctx context.Context, userId uint, refreshToken string) (*entity.UserSession, error) {
	sql, args, err := s.Builder.
		Delete("user_sessions").
		Where(squirrel.Eq{"user_id": userId, "refresh_token": refreshToken}).
		Suffix("RETURNING \"user_id\", \"refresh_token\", \"expires_at\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserSessions - Delete - s.Builder: %w", err)
	}

	row := s.Pool.QueryRow(ctx, sql, args...)
	us := entity.UserSession{}
	if err = row.Scan(&us.UserId, &us.RefreshToken, &us.ExpiresAt, &us.CreatedAt); err != nil {
		return nil, fmt.Errorf("UserSessions - Delete - row.Scan: %w", err)
	}

	return &us, nil
}
