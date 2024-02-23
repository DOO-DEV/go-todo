package userhandler

import "go-todo/internal/service/userservice"

type Handler struct {
	userSvc userservice.Service
}

func New(userSvc userservice.Service) Handler {
	return Handler{userSvc: userSvc}
}
