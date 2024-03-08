package model

import (
	"go-todo/internal/domain"
	"time"
)

type Token struct {
	ID                    uint      `gorm:"id"`
	UserID                uint      `gorm:"user_id"`
	JTI                   string    `gorm:"jti"`
	RefreshToken          string    `gorm:"refresh_token"`
	User                  User      `gorm:"foreignKey:user_id"`
	AccessTokenExpiresAt  time.Time `gorm:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `gorm:"refresh_token_expires_at"`
	CreatedAt             time.Time `gorm:"created_at"`
	UpdatedAt             time.Time `gorm:"updated_at"`
}

func (t Token) TableName() string {
	return "token"
}

func (t Token) ToDomain() domain.Token {
	return domain.Token{
		ID:                    t.ID,
		UserID:                t.UserID,
		JTI:                   t.JTI,
		AccessTokenExpiresAt:  t.AccessTokenExpiresAt,
		RefreshToken:          t.RefreshToken,
		RefreshTokenExpiresAt: t.RefreshTokenExpiresAt,
	}
}
