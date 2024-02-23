package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"go-todo/cmd/migrator"
	"go-todo/internal/config"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func Execute() {
	root := &cobra.Command{Short: "Todo application write in golang with fun :)"}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	location, err := time.LoadLocation(cfg.TZ)
	if err != nil {
		log.Fatalln(err)
	}
	time.Local = location

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	root.AddCommand(
		migrator.Migrator{}.Command(ctx, cfg),
	)

	err = root.Execute()
	if err != nil {
		log.Fatalf("failed to execute root command: %v\n", err)
	}
}
