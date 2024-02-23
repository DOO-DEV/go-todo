package user

import (
	"context"
	"go-todo/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(ctx context.Context, user domain.User) (domain.User, error)
	LoginUser(ctx context.Context, username string) (domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return repository{
		db: db,
	}
}
