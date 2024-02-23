package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"go-todo/cmd/migrator"
	"go-todo/cmd/server"
	"go-todo/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
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

	var lvl zapcore.Level
	if err := lvl.Set(cfg.LogLevel); err != nil {
		log.Printf("cannot parse log level %s: %s", cfg.LogLevel, err)

		lvl = zapcore.WarnLevel
	}
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCfg := &zap.Config{
		Level:             zap.NewAtomicLevelAt(lvl),
		Development:       cfg.AppDebug,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}
	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatalf("logger creation failed %s", err)
	}

	root.AddCommand(
		migrator.Migrator{}.Command(ctx, cfg, logger),
		server.Server{}.Command(ctx, cfg, logger),
	)

	err = root.Execute()
	if err != nil {
		log.Fatalf("failed to execute root command: %v\n", err)
	}
}
