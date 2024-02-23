package user

import (
	"context"
	"go-todo/internal/domain"
	"go-todo/internal/model"
)

func (r repository) RegisterUser(ctx context.Context, u domain.User) (domain.User, error) {
	user := model.User{}
	err := r.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user.ToDomain(), nil
}

func (r repository) LoginUser(ctx context.Context, username string) (domain.User, error) {
	return r.GetUserByUsername(ctx, username)
}

func (r repository) GetUserByID(ctx context.Context, userID uint) (domain.User, error) {
	user := model.User{ID: userID}

	err := r.db.WithContext(ctx).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user.ToDomain(), nil
}

func (r repository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user.ToDomain(), nil
}
