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

// ================================Task implementation================================

func (db *Postgres) AddTask(task models.Task) (int64, error) {
	panic("implement me")
}

func (db *Postgres) ListTasksByUser(userID int64) ([]models.Task, error) {
	panic("implement me")
}

func (db *Postgres) ListAllTasks() ([]models.Task, error) {
	panic("implement me")
}

func (db *Postgres) ChangeTask(task models.Task) error {
	panic("implement me")
}

func (db *Postgres) DeleteTask(id int64) error {
	panic("implement me")
}

// ================================Task implementation================================

// ================================User implementation================================

func (db *Postgres) AddUser(user models.User) (int64, error) {
	panic("implement me")
}

func (db *Postgres) ChangeUser(user models.User) error {
	panic("implement me")
}

func (db *Postgres) DeleteUser(id int64) error {
	panic("implement me")
}

// ================================User implementation================================

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
