package db

import (
	"fmt"
	"os"
)

type Storage interface {
	AddTask(name string) (int64, error)
	ListTask() ([]string, error)
	ChangeTask() error
	DeleteTask() error
	ConnectionInfo() string
}

func StoragePath() string {
	path := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	return path
}
