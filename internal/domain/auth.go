package domain

import "time"

type Auth struct {
	TokenType            string
	AccessToken          string
	RefreshToken         string
	AccessTokenExpiresAt time.Duration
}
