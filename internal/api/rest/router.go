package rest

import (
	"go-todo/internal/api/rest/handler/healthhandler"
	"go-todo/internal/api/rest/handler/userhandler"
)

func (s Server) SetupMonitoringRoutes(healthHandler healthhandler.Handler) {
	r := s.engine

	r.GET("/health", healthHandler.Health)
}

func (s Server) SetupRoutes(userHandler userhandler.Handler) {
	r := s.engine

	v1 := r.Group("/v1")

	{
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)
	}
}
