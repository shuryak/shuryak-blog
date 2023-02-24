package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/shuryak/shuryak-blog/internal/user/entity"
	"github.com/shuryak/shuryak-blog/internal/user/usecase"
	"github.com/shuryak/shuryak-blog/pkg/postgres"
)

type UsersRepo struct {
	*postgres.Postgres
}

func NewUsersRepo(pg *postgres.Postgres) *UsersRepo {
	return &UsersRepo{pg}
}

// Check for implementation
var _ usecase.UsersRepo = (*UsersRepo)(nil)

func (r UsersRepo) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("username", "hashed_password", "role", "created_at").
		Values(user.Username, user.HashedPassword, user.Role, user.CreatedAt).
		Suffix("RETURNING \"id\", \"username\", \"hashed_password\", \"role\", \"created_at\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresUserStore - Create - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	dbUser := entity.User{}
	if err = row.Scan(&dbUser.Id, &dbUser.Username, &dbUser.HashedPassword, &dbUser.Role, &dbUser.CreatedAt); err != nil {
		return nil, fmt.Errorf("PostgresUserStore - Create - row.Scan: %w", err)
	}

	return &dbUser, nil
}

func (r UsersRepo) GetById(ctx context.Context, id uint32) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UsersRepo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	sql, args, err := r.Builder.
		Select("id", "username", "hashed_password", "role", "created_at").
		From("users").
		Where(squirrel.Eq{"username": username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostgresUserStore - GetByUsername - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	u := entity.User{}
	err = row.Scan(&u.Id, &u.Username, &u.HashedPassword, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("PostgresUserStore - GetByUsername - row.Scan: %w", err)
	}

	return &u, nil
}

func (r UsersRepo) Update(ctx context.Context, u entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r UsersRepo) Delete(ctx context.Context, u entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
