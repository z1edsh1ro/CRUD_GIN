package service

import (
	"fmt"
	"main/internal/adapter/repository"
	"main/internal/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	Repo repository.DBTodoRepository
}

func NewTodoService(repo repository.DBTodoRepository) *TodoService {
	return &TodoService{Repo: repo}
}

func (service *TodoService) GetAllTodo() ([]model.Todo, error) {
	return service.Repo.GetAll()
}

func (service *TodoService) GetTodo(id string) (model.Todo, error) {
	return service.Repo.GetById(id)
}

func (service *TodoService) CreateTodo(entry model.Todo) error {
	return service.Repo.Create(entry)
}

func (service *TodoService) UpdateTodo(id string, entry model.Todo) (*mongo.UpdateResult, error) {
	todo, err := service.Repo.GetById(id)

	if err != nil {
		return nil, err
	}

	if (todo == model.Todo{}) {
		return nil, fmt.Errorf("ERROR TODO ID: %s NOT FOUND", id)
	}

	return service.Repo.Update(id, entry)
}

func (service *TodoService) DeleteTodo(id string) error {
	return service.Repo.Delete(id)
}
