package db

import "github.com/Beloslav13/cli-todo/internal/models"

type Storage interface {
	AddTask(name, status string) (int64, error)
	ListTask() ([]models.Task, error)
	ChangeTask(id int64, name, status *string) error
	DeleteTask(id int64) error
	Close() error
}
