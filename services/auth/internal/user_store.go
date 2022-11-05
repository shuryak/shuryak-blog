package internal

import (
	"auth/internal/entity"
	// TODO: fix code duplication
	"auth/pkg/postgres"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type UserStore interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}

type PostgresUserStore struct {
	*postgres.Postgres
}

// Check for implementation
var _ UserStore = (*PostgresUserStore)(nil)

func NewPostgresUserStore(pg *postgres.Postgres) *PostgresUserStore {
	return &PostgresUserStore{pg}
}

func (us *PostgresUserStore) Create(ctx context.Context, u entity.User) (*entity.User, error) {
	sql, args, err := us.Builder.
		Insert("users").
		Columns("username", "hashed_password", "role", "created_at").
		Values(u.Username, u.HashedPassword, u.Role, u.CreatedAt).
		Suffix("RETURNING \"id\", \"username\", \"hashed_password\", \"role\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserStore - Create - us.Builder: %w", err)
	}

	row := us.Pool.QueryRow(ctx, sql, args...)
	newUser := entity.User{}
	if err = row.Scan(&newUser.Id, &newUser.Username, &newUser.HashedPassword, &newUser.CreatedAt); err != nil {
		return nil, fmt.Errorf("UserStore - Create - row.Scan: %w", err)
	}

	return &newUser, nil
}

func (us *PostgresUserStore) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	sql, args, err := us.Builder.
		Select("id", "username", "hashed_password", "role", "created_at").
		From("users").
		Where(squirrel.Eq{"username": username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserStore - GetByUsername - us.Builder: %w", err)
	}

	row := us.Pool.QueryRow(ctx, sql, args...)
	u := entity.User{}
	err = row.Scan(&u.Id, &u.Username, &u.HashedPassword, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("UserStore - GetByUsername - row.Scan: %w", err)
	}

	return &u, nil
}
