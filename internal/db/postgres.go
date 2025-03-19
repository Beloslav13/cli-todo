package db

import (
	"database/sql"
	"fmt"
	"github.com/Beloslav13/cli-todo/internal/models"
	_ "github.com/lib/pq"
	"os"
)

type Postgres struct {
	conn *sql.DB
}

func (db *Postgres) AddTask(name, status string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Postgres) ListTask() ([]models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (db *Postgres) ChangeTask(id int64, name, status *string) error {
	//TODO implement me
	panic("implement me")
}

func (db *Postgres) DeleteTask(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (db *Postgres) Close() error {
	return db.conn.Close()
}

// New создаёт подключение к базе данных и возвращает интерфейс Storage
func New() (Storage, error) {
	path := StoragePath()
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	// Проверка соединения
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{conn: db}, nil
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
