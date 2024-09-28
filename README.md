# Go todos / CLI приложения "Список задач"
****
### RUN
1. `go run cmd/todos/todos.go` `command` `arguments (if need be)`
****
### BUILD and RUN
1. `go build cmd/todos/todos.go`
2. `todos.extenstion` `command` `arguments (if need be)`
****
### Commands
* `--ls` for get all todos
* `--add {title}` for add new todo
* `--rm {id}` for delete todo by ID
* `--cs {id}` for change todo status by ID
* `--ct {id} {title}` for change todo title by ID
****
### TESTS:
###### MODELS/TODOS
1. `go test -v -coverprofile` `models_todos_cover.out`  `models\todos.go` `models\todos_test.go`
2. `go tool cover --html=models_todos_cover.out`
###### COMMANDS
1. `go test -v -coverprofile` `commands_cover.out`  `commands.go` `commands.go`
2. `go tool cover --html=commands_cover.out`