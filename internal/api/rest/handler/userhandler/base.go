package userhandler

import (
	"go-todo/internal/service/userservice"
	"go-todo/internal/validation/uservalidation"
)

type Handler struct {
	userSvc   userservice.Service
	validator uservalidation.Validator
}

func New(userSvc userservice.Service, validator uservalidation.Validator) Handler {
	return Handler{
		userSvc:   userSvc,
		validator: validator,
	}
}
