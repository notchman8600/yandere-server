package domain

type Todo struct {
	Task   string `json:"task"`
	Desc   string `json:"desc"`
	TaskId string `json:"task_id"`
	IsDone bool   `json:"is_done"`
}

type TodoSuccessMessage struct {
	Message string
	TaskId  string
}

type Todos []Todo