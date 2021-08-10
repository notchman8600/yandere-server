package usecase

import (
	"github.com/pkg/errors"

	"log"

	"d-meeting.com/d-server/domain"
)

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Create(meeting domain.Todo) (message domain.TodoSuccessMessage, err error) {
	identifier, err := interactor.TaskRepository.Store(meeting)
	if err != nil {
		return
	}
	message.TaskId = identifier
	message.Message = "success"
	// _, err = interactor.TaskRepository.FindById(identifier)

	return
}
func (interactor *TaskInteractor) CalcProgress(id string) (progress float64, err error) {
	progress, err = interactor.TaskRepository.CalcProgress(id)
	if err != nil {
		log.Fatalln(errors.WithStack(err))
	}
	return
}
