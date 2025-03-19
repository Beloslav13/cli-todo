package cli

import (
	"fmt"
	"github.com/Beloslav13/cli-todo/internal/models"
	"github.com/urfave/cli/v2"
	"log/slog"
	"strings"
)

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

func (app *App) ListTasksByUser(cCtx *cli.Context) error {
	userID := cCtx.Int64("user_id")

	tasks, err := app.storage.ListTasksByUser(userID)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Printf("Задачи пользователя %d не найдены.\n", userID)
		return nil
	}

	app.log.Info("listing tasks successfully", slog.Int64("user_id", userID))

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Задачи пользователя %d:\n", userID))

	for i, task := range tasks {
		createdAtFormatted := task.CreatedAt.Format("02 Jan 2006 15:04")

		sb.WriteString(fmt.Sprintf("\n--------------- Задача %d ---------------\n", i+1))
		sb.WriteString(fmt.Sprintf("Идентификатор: %d\n", task.ID))
		sb.WriteString(fmt.Sprintf("Наименование: %s\n", task.Name))
		sb.WriteString(fmt.Sprintf("Статус: %s\n", task.Status))
		sb.WriteString(fmt.Sprintf("Создана: %s\n", createdAtFormatted))
		sb.WriteString(fmt.Sprintf("--------------- Задача %d ---------------\n", i+1))

		//if i < len(tasks)-1 {
		//	sb.WriteString("\n=======> Ищу следующую\n")
		//}
	}

	sb.WriteString(fmt.Sprintf("\nКоличество задач: %d\n", len(tasks)))

	fmt.Println(sb.String())
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
