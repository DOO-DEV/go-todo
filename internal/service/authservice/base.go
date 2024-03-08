package authservice

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-todo/internal/config"
	"go-todo/internal/domain"
	"log"
	"os"
)

type tokenRepository interface {
	Create(ctx context.Context, token domain.Token) (domain.Token, error)
}

type Service struct {
	tokenRepository tokenRepository
	tokenCfg        config.UserToken
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
}

func New(tokenRepo tokenRepository, tokenCfg config.UserToken) Service {
	svc := Service{
		tokenRepository: tokenRepo,
		tokenCfg:        tokenCfg,
	}
	privateKey, publicKey, err := svc.loadKeys()
	if err != nil {
		log.Fatalln(err)
	}
	svc.privateKey = privateKey
	svc.publicKey = publicKey

	return svc
}

func (s Service) loadKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	var bytesPrivateKey, bytesPublicKey []byte
	var err error
	if s.tokenCfg.PublicKey == "" {
		bytesPublicKey, err = os.ReadFile(s.tokenCfg.PublicKeyFilePath)
	} else {
		bytesPublicKey, err = base64.StdEncoding.DecodeString(s.tokenCfg.PublicKey)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("reading public key: %w", err)
	}

	if s.tokenCfg.PrivateKey == "" {
		bytesPrivateKey, err = os.ReadFile(s.tokenCfg.PrivateKeyFilePath)
	} else {
		bytesPrivateKey, err = base64.StdEncoding.DecodeString(s.tokenCfg.PrivateKey)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("reading private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bytesPrivateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing private key: %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(bytesPublicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing public key: %w", err)
	}

	return privateKey, publicKey, nil
}
