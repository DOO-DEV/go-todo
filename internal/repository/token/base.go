package tokenrepository

import (
	"context"
	"go-todo/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, token domain.Token) (domain.Token, error)
}
type tokenRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return tokenRepository{
		db: db,
	}
}
