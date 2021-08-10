package usecase

import "d-meeting.com/d-server/domain"

type TaskRepository interface {
	Store(domain.Todo) (string, error)
	FindById(string) (domain.Todos, error)
	CalcProgress(string) (float64, error)
}
