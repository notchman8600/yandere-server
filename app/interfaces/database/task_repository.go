package database

import (
	"fmt"
	"log"
	"time"

	"d-meeting.com/d-server/domain"
)

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) Store(task domain.Todo) (id string, err error) {
	_, err = repo.Execute(
		"INSERT INTO tasks(task_id, name, agenda_id) values($1,$2,$3)", "example_id", "hogehoge", "example_agenda_id",
	)
	if err != nil {
		return "", err
	}
	id = "example_id"
	return
}
func (repo *TaskRepository) FindById(identifier string) (data domain.Todos, err error) {
	rows, err := repo.Query("select * from tasks where task_id = $1", identifier)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var task_id string
		var name string
		var agenda_id string
		var is_done bool
		var created_at time.Time
		var updated_at time.Time

		if err := rows.Scan(&task_id, &name, &is_done, &agenda_id, &created_at, &updated_at); err != nil {
			log.Fatalln(err)
			continue
		}
		task := domain.Todo{
			Task:     name,
			TaskId:   task_id,
			AgendaId: agenda_id,
			IsDone:   false,
		}
		fmt.Println(task)
		data = append(data, task)
	}
	return
}
