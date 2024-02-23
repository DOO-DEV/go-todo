package health

import "go-todo/internal/service/health"

type Handler struct {
	healthSvc health.Service
}

func New(healthSvc health.Service) Handler {
	return Handler{
		healthSvc: healthSvc,
	}
}
