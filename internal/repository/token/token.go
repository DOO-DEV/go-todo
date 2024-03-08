package tokenrepository

import (
	"context"
	"go-todo/internal/domain"
	"go-todo/internal/model"
)

func (t tokenRepository) Create(ctx context.Context, token domain.Token) (domain.Token, error) {
	mt := model.Token{
		UserID:                token.UserID,
		JTI:                   token.JTI,
		RefreshToken:          token.RefreshToken,
		AccessTokenExpiresAt:  token.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: token.RefreshTokenExpiresAt,
	}
	if err := t.db.WithContext(ctx).Create(&mt).Error; err != nil {
		return domain.Token{}, err
	}

	return mt.ToDomain(), nil
}
