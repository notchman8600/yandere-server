package domain

type Todo struct {
	Task   string
	Assign []Human
	IsDone bool
	TaskId string
	AgendaId string
}

type TodoSuccessMessage struct {
	Message string
	TaskId  string
}

type Todos []Todo
