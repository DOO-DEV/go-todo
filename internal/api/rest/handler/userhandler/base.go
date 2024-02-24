package userhandler

import (
	"go-todo/internal/api/rest/transformer/usertransformer"
	"go-todo/internal/service/userservice"
	"go-todo/internal/validation/uservalidation"
)

type Handler struct {
	userSvc     userservice.Service
	validator   uservalidation.Validator
	transformer usertransformer.Transformer
}

func New(userSvc userservice.Service, validator uservalidation.Validator, transformer usertransformer.Transformer) Handler {
	return Handler{
		userSvc:     userSvc,
		validator:   validator,
		transformer: transformer,
	}
}
