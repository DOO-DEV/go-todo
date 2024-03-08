package usertransformer

import (
	"go-todo/internal/domain"
)

type response struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t Transformer) TransformAuthCredentials(auth domain.Auth) response {
	return response{
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
	}
}
