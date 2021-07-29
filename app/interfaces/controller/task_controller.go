package controller

import (
	"fmt"
	"log"

	"d-meeting.com/d-server/domain"
	"d-meeting.com/d-server/interfaces/database"
	"d-meeting.com/d-server/usecase"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(sqlHandler database.SqlHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TaskController) Create(c Context) {
	// テストオブジェクト

	task := domain.Todo{}
	err := c.Bind(&task)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(task.TaskId)
	u := domain.TodoSuccessMessage{}

	task_id, err := controller.Interactor.Create(task)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	u.TaskId = task_id.TaskId
	u.Message = "success"
	c.JSON(201, u)
}

func (controller *TaskController) TestCreate(c Context) {
	// テストオブジェクト
	task := domain.Todo{}
	task.Desc = "hogehoge"
	task.Task = "example task"

	u := domain.TodoSuccessMessage{}

	task_id, err := controller.Interactor.Create(task)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	u.TaskId = task_id.TaskId
	u.Message = "success"
	c.JSON(201, u)

}
func (controller *TaskController) Read(c Context) {
	tasks := domain.Todos{}
	c.Bind(&tasks)
	tasks, err := controller.Interactor.TaskRepository.FindById("example_user_id")
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, tasks)
}