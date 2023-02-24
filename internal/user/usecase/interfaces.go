package usecase

import (
	"context"
	"github.com/shuryak/shuryak-blog/internal/user/entity"
)

type (
	Users interface {
		Create(ctx context.Context, u entity.User) (*entity.User, error)
		GetByUsername(ctx context.Context, username string) (*entity.User, error)
	}

	UsersRepo interface {
		Create(ctx context.Context, u entity.User) (*entity.User, error)
		GetById(ctx context.Context, id uint32) (*entity.User, error)
		GetByUsername(ctx context.Context, username string) (*entity.User, error)
		Update(ctx context.Context, u entity.User) (*entity.User, error)
		Delete(ctx context.Context, u entity.User) (*entity.User, error)
	}

	UserSessions interface {
		Add(ctx context.Context, userId uint32) (*entity.UserSession, error)
		Refresh(ctx context.Context, userId uint32) (*entity.UserSession, error)
	}

	UserSessionsRepo interface {
		Create(ctx context.Context, us entity.UserSession) (*entity.UserSession, error)
		Get(ctx context.Context, id uint32) (*entity.UserSession, error)
		Update(ctx context.Context, us entity.UserSession) (*entity.UserSession, error)
		Delete(ctx context.Context, us entity.UserSession) (*entity.UserSession, error)
	}
)
