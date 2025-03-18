package db

type Storage interface {
	AddTask(name string) (int64, error)
	ListTask() ([]string, error)
	ChangeTask() error
	DeleteTask() error
	ConnectionInfo() string
}
