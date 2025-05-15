package service

import (
	"fmt"
	"log"
	"main/internal/domain/model"
	"main/internal/domain/port"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	Port port.TodoRepositoryInterface
}

func (service *TodoService) GetAllTodo() ([]model.Todo, error) {
	todos, err := service.Port.GetAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return todos, nil
}

func (service *TodoService) GetTodo(id string) (model.Todo, error) {
	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return model.Todo{}, err
	}

	todo, err := service.Port.GetById(mongoId)

	if err != nil {
		log.Println(err)
		return model.Todo{}, err
	}

	return todo, nil
}

func (service *TodoService) CreateTodo(entry model.Todo) error {
	err := service.Port.Create(entry)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (service *TodoService) UpdateTodo(id string, entry model.Todo) error {
	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return err
	}

	todo, err := service.Port.GetById(mongoId)

	if err != nil {
		return err
	}

	if (todo == model.Todo{}) {
		return fmt.Errorf("ERROR TODO ID: %s NOT FOUND", id)
	}

	err = service.Port.Update(mongoId, entry)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (service *TodoService) DeleteTodo(id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println(err)
		return err
	}

	err = service.Port.Delete(mongoId)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
