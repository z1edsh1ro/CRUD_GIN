package service

import (
	"fmt"
	"main/internal/domain/model"
	"main/internal/domain/port"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	Port port.TodoRepository
}

func (service *TodoService) GetAllTodo() ([]model.Todo, error) {
	return service.Port.GetAll()
}

func (service *TodoService) GetTodo(id string) (model.Todo, error) {
	return service.Port.GetById(id)
}

func (service *TodoService) CreateTodo(entry model.Todo) error {
	return service.Port.Create(entry)
}

func (service *TodoService) UpdateTodo(id string, entry model.Todo) (*mongo.UpdateResult, error) {
	todo, err := service.Port.GetById(id)

	if err != nil {
		return nil, err
	}

	if (todo == model.Todo{}) {
		return nil, fmt.Errorf("ERROR TODO ID: %s NOT FOUND", id)
	}

	return service.Port.Update(id, entry)
}

func (service *TodoService) DeleteTodo(id string) error {
	return service.Port.Delete(id)
}
