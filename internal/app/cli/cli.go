package cli

import (
	"github.com/Beloslav13/cli-todo/internal/db"
	"github.com/urfave/cli/v2"
	"log/slog"
)

type App struct {
	log     *slog.Logger
	storage db.Storage
}

func New(log *slog.Logger, storage db.Storage) *App {
	return &App{
		log:     log,
		storage: storage,
	}
}

func (app *App) Run(args []string) error {
	cliApp := &cli.App{
		Commands: app.commands(),
	}

	return cliApp.Run(args)
}
