package userservice

import (
	"context"
	"go-todo/internal/api/rest/request"
	"go-todo/internal/domain"
	"go-todo/internal/errors"
	"go-todo/pkg/password"
)

func (s Service) Login(ctx context.Context, req request.LoginUserRequest) (domain.Auth, error) {
	user, err := s.userRepository.LoginUser(ctx, req.Username)
	if err != nil {
		return domain.Auth{}, err
	}

	if err := password.ComparePassword(user.Password, req.Password); err != nil {
		return domain.Auth{}, errors.ErrNotFound
	}

	authResult, err := s.tokenGenerator.CreateUserTokens(ctx, domain.User{ID: user.ID, Username: user.Username})
	if err != nil {
		return domain.Auth{}, err
	}

	return authResult, nil
}
