package rest

import "go-todo/internal/api/rest/handler/health"

func (s Server) SetupMonitoringRoutes(healthHandler health.Handler) {
	r := s.engine

	r.GET("/health", healthHandler.Health)
}

func (s Server) SetupRoutes() {}
