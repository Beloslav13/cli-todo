package db

import (
	"database/sql"
	"fmt"
	"github.com/Beloslav13/cli-todo/internal/models"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

type Postgres struct {
	conn *sql.DB
}

// ================================Task implementation================================

func (db *Postgres) AddTask(task models.Task) (int64, error) {
	const op = "db.postgres.AddTask"
	var id int64
	q := `INSERT INTO tasks (user_id, name, status) VALUES ($1, $2, $3) RETURNING id`

	err := db.conn.QueryRow(q, task.UserID, task.Name, task.Status).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (db *Postgres) ListTasksByUser(userID int64, filters map[string]string) ([]models.Task, error) {
	const op = "db.postgres.ListTasksByUser"
	q := `SELECT id, name, status, created_at FROM tasks WHERE user_id = $1`
	args := []interface{}{userID}
	argIdx := 2

	// Применяем фильтр по статусу, если передан
	if status, ok := filters["status"]; ok && status != "" {
		q += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, status)
		argIdx++
	}

	// Проверяем корректность сортировки
	sort := filters["sort"]
	sortColumns := map[string]bool{"id": true, "created_at": true}
	if !sortColumns[sort] {
		sort = "created_at"
	}

	// Проверяем корректность порядка сортировки
	order := strings.ToUpper(filters["order"])
	if order != "ASC" && order != "DESC" {
		order = "DESC"
	}

	q += fmt.Sprintf(" ORDER BY %s %s", sort, order)

	rows, err := db.conn.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.CreatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
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
	const op = "db.postgres.AddUser"
	var id int64
	q := `INSERT INTO users (username) VALUES ($1) RETURNING id`

	err := db.conn.QueryRow(q, user.Username).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
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
