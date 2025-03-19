package app

import (
	"github.com/Beloslav13/cli-todo/internal/db"
	"github.com/Beloslav13/cli-todo/internal/models"
	"github.com/urfave/cli/v2"
	"log/slog"
	"os"
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

func (app *App) Run() error {
	cliApp := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add-task",
				Aliases: []string{"at"},
				Usage:   "add a task to the list",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:     "user_id",
						Usage:    "User ID for the task",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Task name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "status",
						Usage:    "Task status (new, in_progress, completed)",
						Required: true,
					},
				},
				Action: app.AddTask,
			},
			{
				Name:    "add-user",
				Aliases: []string{"au"},
				Usage:   "add a user",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "username",
						Usage:    "Username",
						Required: true,
					},
				},
				Action: app.AddUser,
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		return err
	}
	return nil
}

func (app *App) AddTask(cCtx *cli.Context) error {
	userID := cCtx.Int64("user_id")
	status := cCtx.String("status")
	name := cCtx.String("name")

	task := models.Task{
		UserID: userID,
		Name:   name,
		Status: status,
	}

	taskId, err := app.storage.AddTask(task)
	if err != nil {
		return err
	}

	app.log.Info("task added successfully", slog.Int64("task_id", taskId))
	return nil
}

func (app *App) AddUser(cCtx *cli.Context) error {
	username := cCtx.String("username")

	user := models.User{
		Username: username,
	}
	userId, err := app.storage.AddUser(user)
	if err != nil {
		return err
	}
	app.log.Info("user added successfully", slog.Int64("user_id", userId))
	return nil
}
