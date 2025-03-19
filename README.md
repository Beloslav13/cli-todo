# CLI-Todo


migrate: ```go run ./cmd/migrator/main.go --migrations-path=./migrations --rollback=false```


Task status Enum:
* new
* in_progress
* completed


commands:
* Add task - add-task as (at): ```go run cmd/cli/main.go add-task --user_id=1 --name="Тестовый таск3" --status="in_progress"```
* Task list - list-task as (tl) ```go run cmd/cli/main.go task-list --user_id=1```
* Add user - add-user as (au): ```go run cmd/cli/main.go add-user --username=bobra```