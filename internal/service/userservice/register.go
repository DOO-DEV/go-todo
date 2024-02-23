package userservice

import (
	"context"
	"go-todo/internal/api/rest/request"
	"go-todo/internal/domain"
	"go-todo/internal/errors"
	"go-todo/pkg/password"
)

func (s Service) Register(ctx context.Context, req request.RegisterUserRequest) (domain.User, error) {
	hashedPasswd, err := password.HashPassword(req.Password)
	if err != nil {
		return domain.User{}, errors.ErrSomethingWentWrong
	}

	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPasswd,
	}

	if _, err := s.userRepository.RegisterUser(ctx, user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}
