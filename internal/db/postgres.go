package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Postgres struct {
	conn *sql.DB
}

func (db Postgres) AddTask(name string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (db Postgres) ListTask() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (db Postgres) ChangeTask() error {
	//TODO implement me
	panic("implement me")
}

func (db Postgres) DeleteTask() error {
	//TODO implement me
	panic("implement me")
}

func (db Postgres) ConnectionInfo() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
}

// New создаёт подключение к базе данных и возвращает интерфейс Storage
func New(path string) (Storage, error) {
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

func (db Postgres) Close() error {
	return db.conn.Close()
}
