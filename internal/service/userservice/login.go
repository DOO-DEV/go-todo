package userservice

import (
	"context"
	"go-todo/internal/api/rest/request"
	"go-todo/internal/domain"
	"go-todo/internal/errors"
	"go-todo/pkg/password"
)

func (s Service) Login(ctx context.Context, req request.LoginUserRequest) (domain.User, error) {
	user, err := s.userRepository.LoginUser(ctx, req.Username)
	if err != nil {
		return domain.User{}, err
	}

	if err := password.ComparePassword(user.Password, req.Password); err != nil {
		return domain.User{}, errors.ErrNotFound
	}

	return user, nil
}
