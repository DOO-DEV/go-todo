package userservice

import (
	"context"
	"go-todo/internal/domain"
)

type userRepository interface {
	RegisterUser(ctx context.Context, user domain.User) (domain.User, error)
	LoginUser(ctx context.Context, username string) (domain.User, error)
}

type tokenGenerator interface {
	CreateUserTokens(ctx context.Context, user domain.User) (domain.Auth, error)
}

type Service struct {
	userRepository userRepository
	tokenGenerator tokenGenerator
}

func New(userRepo userRepository, tokenGenerator tokenGenerator) Service {
	return Service{
		userRepository: userRepo,
		tokenGenerator: tokenGenerator,
	}
}
