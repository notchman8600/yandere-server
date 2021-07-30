package usecase

import "d-meeting.com/d-server/domain"

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
