# CLI-Todo


migrate: ```go run ./cmd/migrator/main.go --migrations-path=./migrations --rollback=false```


Status Enum:
* new
* in_progress
* completed


commands:
* Add task - add-task as (at): ```go run cmd/cli/main.go add-task --user_id=1 --name="Тестовый таск3" --status="in_progress"```
* Add user - add-user as (au): ```go run cmd/cli/main.go add-task --username=bobra```