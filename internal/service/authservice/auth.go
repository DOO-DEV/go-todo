package authservice

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go-todo/internal/domain"
	"strings"
	"time"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	Username string
}

func (s Service) CreateUserTokens(ctx context.Context, user domain.User) (domain.Auth, error) {
	accessToken, claims, err := s.createAccessToken(user.Username)
	if err != nil {
		return domain.Auth{}, err
	}

	refreshToken := s.createRefreshToken(accessToken)

	_, err = s.tokenRepository.Create(ctx, domain.Token{
		UserID:                user.ID,
		JTI:                   claims.ID,
		AccessTokenExpiresAt:  claims.ExpiresAt.Time,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: time.Now().Add(s.tokenCfg.RefreshTokenTTL),
	})

	return domain.Auth{
		TokenType:            "Bearer",
		AccessToken:          accessToken,
		AccessTokenExpiresAt: s.tokenCfg.AccessTokenTTL,
		RefreshToken:         refreshToken,
	}, nil
}

func (s Service) createAccessToken(username string) (string, AuthClaims, error) {
	var c AuthClaims
	now := time.Now()
	c.Username = username
	c.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(s.tokenCfg.AccessTokenTTL)),
		ID:        uuid.NewString(),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Subject:   "", // hash userId and save it in this field
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	tokenStr, err := token.SignedString(s.privateKey)

	return tokenStr, c, err
}

func (s Service) createRefreshToken(accessToken string) string {
	t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(accessToken)).String()
	refresh := base64.URLEncoding.EncodeToString([]byte(t))
	return strings.ToUpper(strings.TrimRight(refresh, "="))
}

func (s Service) hashRefreshToken(token string) string {
	hashed := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hashed[:])
}
