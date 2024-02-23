package userservice

import (
	"context"
	"go-todo/internal/domain"
)

type userRepository interface {
	RegisterUser(ctx context.Context, user domain.User) (domain.User, error)
	LoginUser(ctx context.Context, username string) (domain.User, error)
}

type Service struct {
	userRepository userRepository
}

func New(userRepo userRepository) Service {
	return Service{
		userRepository: userRepo,
	}
}
