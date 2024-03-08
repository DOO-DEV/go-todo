package domain

import "time"

type Token struct {
	ID                    uint
	UserID                uint
	JTI                   string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}
