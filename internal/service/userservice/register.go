package userservice

import (
	"context"
	"go-todo/internal/api/rest/request"
	"go-todo/internal/domain"
	"go-todo/internal/errors"
	"go-todo/pkg/password"
)

func (s Service) Register(ctx context.Context, req request.RegisterUserRequest) (domain.Auth, error) {
	hashedPasswd, err := password.HashPassword(req.Password)
	if err != nil {
		return domain.Auth{}, errors.ErrSomethingWentWrong
	}

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPasswd,
	}

	if _, err := s.userRepository.RegisterUser(ctx, user); err != nil {
		return domain.Auth{}, err
	}

	authResult, err := s.tokenGenerator.CreateUserTokens(ctx, domain.User{ID: user.ID, Username: user.Username})
	if err != nil {
		return domain.Auth{}, err
	}

	return authResult, nil
}
