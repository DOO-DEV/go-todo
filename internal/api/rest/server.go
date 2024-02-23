package rest

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo/internal/config"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	engine *echo.Echo
	cfg    config.HttpApi

	logger *zap.Logger
}

func New(cfg config.HttpApi, logger *zap.Logger) *Server {
	r := echo.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	return &Server{
		engine: r,
		cfg:    cfg,
		logger: logger,
	}
}

func (s Server) Serve(ctx context.Context) error {
	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler:           s.engine,
		ReadHeaderTimeout: 10 * time.Second,
	}

	s.logger.Info("server started at", zap.String("address", server.Addr))

	serverErr := make(chan error, 1)
	go func() {
		server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		s.logger.Info("server is shutting down")
		return server.Shutdown(ctx)
	case err := <-serverErr:
		return err
	}
}
