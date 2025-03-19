package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Beloslav13/cli-todo/internal/db"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // Импорт PostgreSQL драйвера
)

func main() {
	var migrationsPath, migrationsTable string
	var rollback bool
	// Флаги для путей миграций
	flag.StringVar(
		&migrationsPath, "migrations-path", "", "path to migrations",
	)
	flag.StringVar(
		&migrationsTable, "migrations-table", "migrations", "path to migrations table",
	)
	flag.BoolVar(
		&rollback, "rollback", false, "flag to rollback migrations",
	)
	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path flag is required")
	}

	// Получаем строку подключения
	storagePath := db.StoragePath()

	// Создаем мигратор
	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("%s&x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	// Если указан флаг отката, то откатываем миграции
	if rollback {
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to rollback")
				return
			}
			panic(err)
		}
		fmt.Println("migration rolled back successfully")
		return
	}

	// Выполняем миграции
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(err)
	}
	fmt.Println("migrations applied successfully")
}
