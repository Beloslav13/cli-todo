-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

-- Добавление индекса на поле username в таблице users для быстрого поиска
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Создание ENUM типа для статуса задач
CREATE TYPE task_status AS ENUM ('new', 'in_progress', 'completed');

-- Создание таблицы задач
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE, -- Внешний ключ для связи с пользователями
    name VARCHAR(255) NOT NULL,
    status task_status NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

-- Индекс для поля status в таблице tasks
CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
