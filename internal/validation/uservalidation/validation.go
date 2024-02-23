package uservalidation

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go-todo/internal/api/rest/request"
	cErr "go-todo/internal/errors"
)

func (v Validator) ValidateRegisterUserRequest(ctx context.Context, req *request.RegisterUserRequest) error {
	return validation.ValidateStructWithContext(ctx, req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 50)),
		validation.Field(&req.Username, validation.Required, validation.Length(4, 15), validation.By(v.isUsernameUnique)),
	)
}

func (v Validator) ValidateLoginUserRequest(ctx context.Context, req *request.LoginUserRequest) error {
	return validation.ValidateStructWithContext(ctx, req,
		validation.Field(&req.Password, validation.Required),
		validation.Field(&req.Username, validation.Required, validation.By(v.isUsernameExists)),
	)
}

func (v Validator) isUsernameUnique(val interface{}) error {
	username, ok := val.(string)
	if !ok {
		return fmt.Errorf("validation failed")
	}
	_, err := v.userRepository.GetUserByUsername(context.Background(), username)
	if err != nil {
		if errors.Is(err, cErr.ErrNotFound) {
			return nil
		}
		return cErr.ErrSomethingWentWrong
	}

	return cErr.ErrUsernameIsNotUnique
}

func (v Validator) isUsernameExists(val interface{}) error {
	username, ok := val.(string)
	if !ok {
		return fmt.Errorf("validation failed")
	}
	_, err := v.userRepository.GetUserByUsername(context.Background(), username)
	if err != nil {
		if errors.Is(err, cErr.ErrNotFound) {
			return err
		}
		return cErr.ErrSomethingWentWrong
	}

	return nil
}
