-- Удаление индекса для поля status в таблице tasks
DROP INDEX IF EXISTS idx_tasks_status;

-- Удаление внешнего ключа, ссылающегося на таблицу users в таблице tasks
ALTER TABLE tasks DROP CONSTRAINT IF EXISTS tasks_user_id_fkey;

-- Удаление таблицы задач
DROP TABLE IF EXISTS tasks;

-- Удаление типа task_status
DROP TYPE IF EXISTS task_status;

-- Удаление индекса для поля username в таблице users
DROP INDEX IF EXISTS idx_users_username;

-- Удаление таблицы пользователей
DROP TABLE IF EXISTS users;
