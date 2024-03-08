package server

import (
	"context"
	"github.com/spf13/cobra"
	"go-todo/infra/mysql"
	"go-todo/internal/api/rest"
	"go-todo/internal/api/rest/handler/healthhandler"
	"go-todo/internal/api/rest/handler/userhandler"
	"go-todo/internal/api/rest/transformer/usertransformer"
	"go-todo/internal/config"
	tokenrepository "go-todo/internal/repository/token"
	"go-todo/internal/repository/user"
	"go-todo/internal/service/authservice"
	"go-todo/internal/service/healthservice"
	"go-todo/internal/service/userservice"
	"go-todo/internal/validation/uservalidation"
	"go-todo/pkg/gorm"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
}

func (s Server) Command(ctx context.Context, cfg *config.Config, logger *zap.Logger) *cobra.Command {
	s.logger = logger
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
		s.logger.Error("failed to connect to db", zap.Error(err))
		return
	}
	gormDb, err := gorm.NewMysqlGorm(db, cfg.AppDebug)
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	// setup repositories
	userRepo := user.New(gormDb)
	tokenRepo := tokenrepository.New(gormDb)

	// setup services
	healthSvc := healthservice.New(db, cfg.HealthToken)
	authSvc := authservice.New(tokenRepo, cfg.UserToken)
	userSvc := userservice.New(userRepo, authSvc)

	// setup validators
	userValidator := uservalidation.New(userRepo)

	// setup transformers
	userTransformer := usertransformer.New()

	// setup handlers
	healthHandler := healthhandler.New(healthSvc)
	userHandler := userhandler.New(userSvc, userValidator, userTransformer)

	l := s.logger.Named("server:")
	server := rest.New(cfg.HttpApi, l)
	server.SetupMonitoringRoutes(healthHandler)
	server.SetupRoutes(userHandler)

	if err := server.Serve(ctx); err != nil {
		s.logger.Error("server error", zap.Error(err))
		return
	}
}
