package domain

type Todo struct {
	Task     string `json:"task"`
	Desc     string `json:"desc"`
	Deadline string `json:"deadline"`
	TaskId   string `json:"task_id"`
	IsDone   bool   `json:"is_done"`
	Status   int    `json:"status"`
}

type TodoSuccessMessage struct {
	Message string
	TaskId  string
}

type Todos []Todo
