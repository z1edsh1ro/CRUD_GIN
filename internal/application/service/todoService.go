package service

type TodoService struct {
	Id string
}

func NewTodoService() *TodoService {
	TodoService := TodoService{
		Id: "test",
	}

	return &TodoService
}
