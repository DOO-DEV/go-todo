package migrator

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"go-todo/infra/mysql"
	"go-todo/internal/config"
	"go.uber.org/zap"
)

type Migrator struct {
	up   bool
	down bool

	logger *zap.Logger
}

func (m Migrator) Command(ctx context.Context, cfg *config.Config, logger *zap.Logger) *cobra.Command {
	m.logger = logger
	cmd := &cobra.Command{
		Use:   "migrator",
		Short: "handle migrations",
		Run: func(_ *cobra.Command, _ []string) {
			m.main(ctx, cfg)
		},
	}

	cmd.Flags().BoolVar(&m.up, "up", false, "up migrations")
	cmd.Flags().BoolVar(&m.down, "down", false, "down migrations")

	return cmd
}

func (m Migrator) main(ctx context.Context, cfg *config.Config) {
	if m.up {
		m.Up(ctx, cfg)
	}
	if m.down {
		m.Down(ctx, cfg)
	}
}

func (m Migrator) Up(ctx context.Context, cfg *config.Config) {
	db, err := mysql.NewClient(ctx, &mysql.Config{
		Username:     cfg.Database.MySql.Username,
		Password:     cfg.Database.MySql.Password,
		Host:         cfg.Database.MySql.Host,
		Port:         cfg.Database.MySql.Port,
		DatabaseName: cfg.Database.MySql.DbName,
		Timezone:     cfg.TZ,
	})
	if err != nil {
		m.logger.Fatal("can't connect to mysql", zap.Error(err))
	}
	defer db.Close()

	driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
	if err != nil {
		m.logger.Fatal("migration driver failed", zap.Error(err))
	}

	dir := "file://migrations"
	mi, err := migrate.NewWithDatabaseInstance(dir, "public", driver)
	if err != nil {
		m.logger.Fatal("failed to init migration", zap.Error(err))
	}

	m.logger.Info("migration starting")

	if err := mi.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			m.logger.Fatal("migration failed", zap.Error(err))
		}
		m.logger.Info("migration no change")
	}

	m.logger.Info("migration successful")
}

func (m Migrator) Down(ctx context.Context, cfg *config.Config) {
	db, err := mysql.NewClient(ctx, &mysql.Config{
		Username:     cfg.Database.MySql.Username,
		Password:     cfg.Database.MySql.Password,
		Host:         cfg.Database.MySql.Host,
		Port:         cfg.Database.MySql.Port,
		DatabaseName: cfg.Database.MySql.DbName,
		Timezone:     cfg.TZ,
	})
	if err != nil {
		m.logger.Fatal("can't connect to mysql", zap.Error(err))
	}
	defer db.Close()

	driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
	if err != nil {
		m.logger.Fatal("migration driver failed", zap.Error(err))
	}

	dir := "file://migrations"
	mi, err := migrate.NewWithDatabaseInstance(dir, "mysql", driver)
	if err != nil {
		m.logger.Fatal("failed to init migration", zap.Error(err))
	}

	m.logger.Info("migration down starting")

	if err := mi.Down(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			m.logger.Fatal("migration downfailed", zap.Error(err))
		}
		m.logger.Info("migration no change")
	}

	m.logger.Info("migration down successful")
}
