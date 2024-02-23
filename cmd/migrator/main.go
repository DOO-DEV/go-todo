package migrator

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"go-todo/infra/mysql"
	"go-todo/internal/config"
	"log"
)

type Migrator struct {
	up   bool
	down bool
}

func (m Migrator) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
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
		log.Fatalf("can't connect to mysql: %v", err)
	}
	defer db.Close()

	driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
	if err != nil {
		log.Fatalf("migration driver failed: %v\n", err)
	}

	dir := "file://migrations"
	mi, err := migrate.NewWithDatabaseInstance(dir, "public", driver)
	if err != nil {
		fmt.Println("error: ", err)
		log.Fatalf("failed to init migration: %v\n", err)
	}

	log.Println("migration starting")

	if err := mi.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migration failed: %v\n", err)
		}
		log.Println("migration no change")
	}

	log.Println("migration successful")
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
		log.Fatalf("can't connect to mysql: %v", err)
	}
	defer db.Close()

	driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
	if err != nil {
		log.Fatalf("migration driver failed: %v\n", err)
	}

	dir := "file://migrations"
	mi, err := migrate.NewWithDatabaseInstance(dir, "mysql", driver)
	if err != nil {
		log.Fatalf("failed to init migration: %v\n", err)
	}

	log.Println("migration down starting")

	if err := mi.Down(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migration downfailed: %v\n", err)
		}
		log.Println("migration no change")
	}

	log.Println("migration down successful")
}
