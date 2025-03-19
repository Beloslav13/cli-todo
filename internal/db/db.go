package db

import "github.com/Beloslav13/cli-todo/internal/models"

type Storage interface {
	TaskStorage
	UserStorage
	Close() error
}

type TaskStorage interface {
	AddTask(task models.Task) (int64, error)
	ListTasksByUser(userID int64, filters map[string]string) ([]models.Task, error)
	ListAllTasks() ([]models.Task, error) // Для админов
	ChangeTask(task models.Task) error
	DeleteTask(id int64) error
}

type UserStorage interface {
	AddUser(user models.User) (int64, error)
	ChangeUser(user models.User) error
	DeleteUser(id int64) error
}
