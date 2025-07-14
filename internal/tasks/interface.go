package tasks

type Repository interface {
	GetAllTasks() ([]Task, error)
}
