package rest

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo/internal/config"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	engine *echo.Echo
	cfg    config.HttpApi
}

func New(cfg config.HttpApi) *Server {
	r := echo.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	return &Server{
		engine: r,
		cfg:    cfg,
	}
}

func (s Server) Serve(ctx context.Context) error {
	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler:           s.engine,
		ReadHeaderTimeout: 10 * time.Second,
	}

	slog.Info("server started at: ", "address", server.Addr)

	serverErr := make(chan error, 1)
	go func() {
		server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		slog.Info("server is shutting down")
		return server.Shutdown(ctx)
	case err := <-serverErr:
		return err
	}
}
