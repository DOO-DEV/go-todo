package healthhandler

import "go-todo/internal/service/healthservice"

type Handler struct {
	healthSvc healthservice.Service
}

func New(healthSvc healthservice.Service) Handler {
	return Handler{
		healthSvc: healthSvc,
	}
}
