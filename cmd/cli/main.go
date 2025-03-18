package main

import (
	"github.com/Beloslav13/cli-todo/internal/config"
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
	// logger
	log := setupLogger(cfg.Base.Env)

	log.Info("starting CLI", slog.String("env", cfg.Base.Env))
	log.Debug("debug message", slog.String("env", cfg.Base.Env))
	log.Warn("warning message", slog.String("env", cfg.Base.Env))
	log.Error("error message", slog.String("env", cfg.Base.Env))

	// init app
	// start app

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
