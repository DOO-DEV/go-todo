package usertransformer

import "go-todo/internal/domain"

type response struct {
	Username string `json:"username"`
}

func (t Transformer) TransformUserRegister(user domain.User) response {
	return response{Username: user.Username}
}

func (t Transformer) TransformUserLogin(user domain.User) response {
	return response{Username: user.Username}
}
