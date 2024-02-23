package uservalidation

import (
	"context"
	"go-todo/internal/domain"
)

type userRepository interface {
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}

type Validator struct {
	userRepository userRepository
}

func New(userRepo userRepository) Validator {
	return Validator{
		userRepository: userRepo,
	}
}
