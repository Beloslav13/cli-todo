package main

import (
	"github.com/Beloslav13/cli-todo/internal/app/cli"
	"github.com/Beloslav13/cli-todo/internal/config"
	"github.com/Beloslav13/cli-todo/internal/db"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.MustLoad[config.CLIConfig]()
	// init logger
	log := setupLogger(cfg.Base.Env)

	log.Info("starting CLI", slog.String("env", cfg.Base.Env))

	// init db
	path := db.StoragePath()
	storage, err := db.New()
	if err != nil {
		log.Error("failed to connect to database", "error", err.Error())
		os.Exit(1)
	}

	log.Info("connected to database ok", slog.String("env", cfg.Base.Env))
	log.Debug("connected storage info", slog.String("info", path))

	// init app
	cliApp := cli.New(log, storage)
	log.Info("initial app success", slog.String("env", cfg.Base.Env))

	// start app
	err = cliApp.Run(os.Args)
	if err != nil {
		log.Error("app in main run error", err.Error())
		os.Exit(1)
	}
	log.Info("command success", slog.String("env", cfg.Base.Env))
	log.Info("stopping CLI", slog.String("env", cfg.Base.Env))

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		// Default to Debug for unknown env values
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
		log.Warn("Unknown environment value, defaulting to debug level")
	}
	return log
}
