package server

import (
	"context"
	"github.com/spf13/cobra"
	"go-todo/infra/mysql"
	"go-todo/internal/api/rest"
	healthH "go-todo/internal/api/rest/handler/health"
	"go-todo/internal/config"
	"go-todo/internal/service/health"
	"log/slog"
)

type Server struct{}

func (s Server) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	c := &cobra.Command{
		Use:   "server",
		Short: "http server",
		Run: func(_ *cobra.Command, _ []string) {
			s.main(ctx, cfg)
		},
	}

	return c
}

func (s Server) main(ctx context.Context, cfg *config.Config) {
	// setup infra
	db, err := mysql.NewClient(ctx, &mysql.Config{
		Username:     cfg.Database.MySql.Username,
		Password:     cfg.Database.MySql.Password,
		Host:         cfg.Database.MySql.Host,
		Port:         cfg.Database.MySql.Port,
		DatabaseName: cfg.Database.MySql.DbName,
		Timezone:     cfg.TZ,
	})
	if err != nil {
		slog.Error("failed to connect to db: ", err)
		return
	}

	// setup services
	healthSvc := health.New(db, cfg.HealthToken)

	// setup handlers
	healthHandler := healthH.New(healthSvc)

	server := rest.New(cfg.HttpApi)
	server.SetupMonitoringRoutes(healthHandler)
	server.SetupRoutes()

	if err := server.Serve(ctx); err != nil {
		slog.Error("server error: ", err)
		return
	}
}
