package database

import (
	"log"
	"time"

	"d-meeting.com/d-server/domain"
	"github.com/google/uuid"
)

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) Store(task domain.Todo) (id string, err error) {
	//TaskIDの設定が無い時は新規発行
	if len(task.TaskId) == 0 {
		uid, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		uid_str := uid.String()
		task.TaskId = uid_str
	}
	_, err = repo.Execute(
		"INSERT INTO tasks(task_id, name, description,user_id,task_status,deadline) values($1,$2,$3,$4,$5,$6)",
		task.TaskId, task.Task, task.Desc, "example_user_id", task.Status, task.Deadline,
	)
	if err != nil {
		return "", err
	}
	id = task.TaskId
	return
}

func (repo *TaskRepository) FindById(identifier string) (data domain.Todos, err error) {
	rows, err := repo.Query("select * from tasks where user_id = $1", identifier)

	if err != nil {
		log.Fatalln(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task_id string
		var name string
		var desc string
		var is_done bool
		var task_status int
		var deadline time.Time
		var created_at time.Time
		var updated_at time.Time
		var user_id string
		if err := rows.Scan(&task_id, &name, &desc, &is_done, &user_id, &task_status, &deadline, &created_at, &updated_at); err != nil {
			log.Fatalln(err)
			continue
		}
		task := domain.Todo{
			Task:   name,
			TaskId: task_id,
			Desc:   desc,
			IsDone: false,
		}

		data = append(data, task)
	}
	return data, nil
}
