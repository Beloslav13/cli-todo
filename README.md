# CLI-Todo


migrate: ```go run ./cmd/migrator/main.go --migrations-path=./migrations --rollback=false```

Task status Enum:
* new
* in_progress
* completed

**команды**:
* task, t  Task-related commands
* user, u  User-related commands
* help, h  Shows a list of commands or help for one command

команды для **task**:
* add, a   Добавить таск
* list, l  Список тасок по user id

OPTIONS **add**:
* --user_id value  User ID for the task (default: 0)
* --name value     Task name
* --status value   Task status (new, in_progress, completed)

```go run cmd/cli/main.go task add --user_id=1 --name="Тестовый таск3" --status="in_progress"```

OPTIONS **list**:
* --user_id value  List tasks by user ID (default: 0)
* --status value   Filter tasks by status (new, in_progress, completed)
* --sort value     Sort tasks by id or created_at (default: "created_at")
* --sort value     Sort tasks by id or created_at (default: "created_at")

```go run cmd/cli/main.go task list --user_id=1 --status=new --sort=created_at --order=asc```

команды для **user**:
* add, a  Добавить пользователя

OPTIONS **add**:
* --username value  Username

```go run cmd/cli/main.go add-user --username=bobra```